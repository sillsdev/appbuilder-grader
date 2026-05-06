package grader

import (
	"appbuilder-grader/models"
)

// checkPerformance provides an example of a fully ignored category
func (g *Grader) checkPerformance() models.Category {
	cat := models.Category{
		Name:        "Performance Metrics",
		Description: "Evaluates the build artifacts for performance optimization (ignored in this run).",
		Weight:      1.0,
		Status:      "ignored",
		Details:     "Performance tests are currently disabled due to missing Lighthouse configuration.",
		LineItems:   make([]models.LineItem, 0),
	}

	bundleSizeItem := models.LineItem{
		Name:        "Bundle Size",
		Description: "Check if the final JavaScript bundle is under 500KB",
		Score:       2.0,
		MaxScore:    10.0,
		Status:      "warning",
		Details:     "Ignored size threshold calculation.",
	}
	cat.LineItems = append(cat.LineItems, bundleSizeItem)

	speedItem := models.LineItem{
		Name:        "Initial Load Speed",
		Description: "Simulate initial load speed",
		Score:       7.0,
		MaxScore:    10.0,
		Status:      "warning",
		Details:     "Ignored network simulation.",
	}
	cat.LineItems = append(cat.LineItems, speedItem)

	return cat
}
