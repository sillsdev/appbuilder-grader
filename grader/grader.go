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
	grader := &Grader{TargetDir: targetDir}

	// Try to find and parse .appDef file
	appDefPath := filepath.Join(targetDir, "*.appDef")
	appDefFiles, err := filepath.Glob(appDefPath)
	if err != nil || len(appDefFiles) == 0 {
		return grader
	}

	appDef, err := parseAppDef(appDefFiles[0])
	if err == nil {
		grader.AppDef = appDef
	}
	return grader
}

func parseAppDef(appDefFile string) (models.AppDef, error) {
	var appDef models.AppDef

	file, err := os.Open(appDefFile)
	if err != nil {
		return appDef, fmt.Errorf("failed to open appdef file: %v", err)
	}
	defer file.Close()

	if err := xml.NewDecoder(file).Decode(&appDef); err != nil {
		return appDef, fmt.Errorf("failed to parse appdef XML: %v", err)
	}

	return appDef, nil
}

func createIgnoredItem(name, desc string, maxScore float64) models.LineItem {
	return models.LineItem{
		Name:        name,
		Description: desc,
		MaxScore:    maxScore,
		Score:       0.0,
		Status:      models.StatusIgnored,
		Details:     "details.not_implemented_yet",
	}
}

func normalizeWeight(weight float64) float64 {
	if weight <= 0 {
		return 1.0
	}
	return weight
}

func determineLineItemStatus(lineItems []models.LineItem) models.Status {
	status := models.StatusPass
	for _, li := range lineItems {
		if li.Status == models.StatusIgnored {
			continue
		}
		if li.Status == models.StatusError {
			status = models.StatusError
		} else if li.Status == models.StatusWarning && status != models.StatusError {
			status = models.StatusWarning
		}
	}
	return status
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

	var totalWeight, weightedScore, unweightedScore, unweightedMax float64

	// First pass: compute scores and determine status
	for i := range categories {
		c := &categories[i]

		var computedScore, computedMax float64
		activeItems := 0

		for _, li := range c.LineItems {
			if li.Status == models.StatusIgnored {
				continue
			}
			activeItems++
			computedScore += li.Score
			computedMax += li.MaxScore
		}

		c.Score = computedScore
		c.MaxScore = computedMax

		// Determine status if not already set
		if c.Status == "" {
			if activeItems == 0 && len(c.LineItems) > 0 {
				c.Status = models.StatusIgnored
			} else {
				c.Status = determineLineItemStatus(c.LineItems)
				if c.Status == models.StatusPass && c.Score < c.MaxScore {
					c.Status = models.StatusWarning
				}
			}
		}

		if c.Status != models.StatusIgnored {
			totalWeight += normalizeWeight(c.Weight)
		}
	}

	// Second pass: calculate percentages and aggregates
	for i := range categories {
		c := &categories[i]
		if c.Status == models.StatusIgnored {
			c.WeightPercentage = 0
			continue
		}

		weight := normalizeWeight(c.Weight)
		if totalWeight > 0 {
			c.WeightPercentage = (weight / totalWeight) * 100
		}
		if c.MaxScore > 0 {
			weightedScore += (c.Score / c.MaxScore) * weight
		}
		unweightedScore += c.Score
		unweightedMax += c.MaxScore
	}

	percentage := 0.0
	if totalWeight > 0 {
		percentage = (weightedScore / totalWeight) * 100
	}

	report := &models.Report{
		TargetDirectory:    g.TargetDir,
		Categories:         categories,
		UnweightedScore:    unweightedScore,
		UnweightedMaxScore: unweightedMax,
		TotalWeight:        totalWeight,
		Percentage:         percentage,
		TotalScore:         percentage, // Total score matches percentage out of 100
		MaxTotalScore:      100.0,
	}

	return report, nil
}
