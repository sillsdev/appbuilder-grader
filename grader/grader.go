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

	// First pass: determine status, calculate scores excluding ignored line items
	for i := range categories {
		c := &categories[i]
		
		status := "pass"
		if c.Status != "" {
			status = c.Status
		}

		var computedScore, computedMax float64
		activeItems := 0

		for _, li := range c.LineItems {
			if li.Status == "ignored" {
				continue
			}
			activeItems++
			computedScore += li.Score
			computedMax += li.MaxScore

			if c.Status == "" {
				if li.Status == "error" {
					status = "error"
				} else if li.Status == "warning" && status != "error" {
					status = "warning"
				}
			}
		}

		c.Score = computedScore
		c.MaxScore = computedMax

		if c.Status == "" {
			if activeItems == 0 && len(c.LineItems) > 0 {
				status = "ignored"
			} else {
				if status == "pass" && c.Score < c.MaxScore {
					status = "warning"
				}
			}
			c.Status = status
		}

		if c.Status == "ignored" {
			continue
		}

		weight := c.Weight
		if weight <= 0 {
			weight = 1.0
		}
		totalWeight += weight
	}

	// Second pass: calculate percentages and aggregates
	for i := range categories {
		c := &categories[i]
		if c.Status == "ignored" {
			c.WeightPercentage = 0
			continue
		}

		weight := c.Weight
		if weight <= 0 {
			weight = 1.0
		}

		if totalWeight > 0 {
			c.WeightPercentage = (weight / totalWeight) * 100
		}
		if c.MaxScore > 0 {
			weightedScore += (c.Score / c.MaxScore) * weight
		}
		unweightedScore += c.Score
		unweightedMax += c.MaxScore
	}

	report := &models.Report{
		TargetDirectory:    g.TargetDir,
		Categories:         categories,
		UnweightedScore:    unweightedScore,
		UnweightedMaxScore: unweightedMax,
		TotalWeight:        totalWeight,
	}

	if totalWeight > 0 {
		report.Percentage = (weightedScore / totalWeight) * 100
	} else {
		report.Percentage = 0
	}
	report.TotalScore = report.Percentage // Assuming total score matches percentage out of 100 for simplicity
	report.MaxTotalScore = 100.0

	return report, nil
}
