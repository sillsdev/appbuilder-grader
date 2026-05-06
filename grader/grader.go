package grader

import (
	"appbuilder-grader/models"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
)

// Grader handles the evaluation of a build directory
type Grader struct {
	TargetDir string
	AppDef    models.AppDef
}

// NewGrader creates a new grader for the specified directory
func NewGrader(targetDir string) *Grader {
	// Open and parse appdef from targetDir/*.appDef (or *.appdef ?)
	appDefPath := filepath.Join(targetDir, "*.appDef")
	appDefFiles, _ := filepath.Glob(appDefPath)
	if len(appDefFiles) > 0 {
		// Parse the first found appdef file
		appDefFile := appDefFiles[0]
		appDef, err := parseAppDef(appDefFile)
		if err == nil {
			return &Grader{TargetDir: targetDir, AppDef: appDef}
		}
	}
	return &Grader{TargetDir: targetDir}
}

func parseAppDef(appDefFile string) (models.AppDef, error) {
	var appDef models.AppDef

	file, err := os.Open(appDefFile)
	if err != nil {
		return appDef, fmt.Errorf("failed to open appdef file: %v", err)
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	if err := decoder.Decode(&appDef); err != nil {
		return appDef, fmt.Errorf("failed to parse appdef XML: %v", err)
	}

	return appDef, nil
}

// Evaluate runs all grading checks and returns a complete report
func (g *Grader) Evaluate() (*models.Report, error) {
	if _, err := os.Stat(g.TargetDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("target directory does not exist: %s", g.TargetDir)
	}

	// Calculate scores for various categories
	categories := []models.Category{
		g.checkScriptureTextContent(),
		g.checkSupplementaryMaterials(),
		g.checkMultimedia(),
		g.checkUserExperience(),
		g.checkUserEngagement(),
		g.checkAccessibility(),
		g.checkPlayStorePresence(),
		g.checkMiscellaneous(),
	}

	// Calculate weighted total score
	var totalWeight, weightedScore, unweightedScore, unweightedMax float64

	// First pass: determine status, calculate scores excluding ignored line items
	for i := range categories {
		c := &categories[i]
		
		status := models.StatusPass
		if c.Status != "" {
			status = c.Status
		}

		var computedScore, computedMax float64
		activeItems := 0

		for _, li := range c.LineItems {
			if li.Status == models.StatusIgnored {
				continue
			}
			activeItems++
			computedScore += li.Score
			computedMax += li.MaxScore

			if c.Status == "" {
				if li.Status == models.StatusError {
					status = models.StatusError
				} else if li.Status == models.StatusWarning && status != models.StatusError {
					status = models.StatusWarning
				}
			}
		}

		c.Score = computedScore
		c.MaxScore = computedMax

		if c.Status == "" {
			if activeItems == 0 && len(c.LineItems) > 0 {
				status = models.StatusIgnored
			} else {
				if status == models.StatusPass && c.Score < c.MaxScore {
					status = models.StatusWarning
				}
			}
			c.Status = status
		}

		if c.Status == models.StatusIgnored {
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
		if c.Status == models.StatusIgnored {
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
