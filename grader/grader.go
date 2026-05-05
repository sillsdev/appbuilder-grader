package grader

import (
	"appbuilder-grader/models"
	"fmt"
	"os"
)

// Grader handles the evaluation of a build directory
type Grader struct {
	TargetDir string
}

// NewGrader creates a new grader for the specified directory
func NewGrader(targetDir string) *Grader {
	return &Grader{TargetDir: targetDir}
}

// Evaluate runs all grading checks and returns a complete report
func (g *Grader) Evaluate() (*models.Report, error) {
	if _, err := os.Stat(g.TargetDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("target directory does not exist: %s", g.TargetDir)
	}

	// Calculate scores for various categories
	categories := []models.Category{
		g.checkFileStructure(),
		g.checkAssets(),
		g.checkManifests(),
	}

	// Calculate weighted total score
	var totalWeight, weightedScore, unweightedScore, unweightedMax float64
	for _, c := range categories {
		weight := c.Weight
		if weight <= 0 {
			weight = 1.0
		}
		totalWeight += weight
	}

	for i := range categories {
		c := &categories[i]
		weight := c.Weight
		if weight <= 0 {
			weight = 1.0
		}

		c.WeightPercentage = (weight / totalWeight) * 100
		weightedScore += (c.Score / c.MaxScore) * weight
		unweightedScore += c.Score
		unweightedMax += c.MaxScore

		if c.Status == "" {
			status := "pass"
			for _, li := range c.LineItems {
				if li.Status == "error" {
					status = "error"
					break
				} else if li.Status == "warning" && status != "error" {
					status = "warning"
				}
			}
			if status == "pass" && c.Score < c.MaxScore {
				status = "warning"
			}
			c.Status = status
		}
	}

	report := &models.Report{
		TargetDirectory:    g.TargetDir,
		Categories:         categories,
		UnweightedScore:    unweightedScore,
		UnweightedMaxScore: unweightedMax,
		TotalWeight:        totalWeight,
	}

	report.Percentage = (weightedScore / totalWeight) * 100
	report.TotalScore = report.Percentage // Assuming total score matches percentage out of 100 for simplicity
	report.MaxTotalScore = 100.0

	return report, nil
}
