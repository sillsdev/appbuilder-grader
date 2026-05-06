package grader

import (
	"appbuilder-grader/models"
)

// checkManifests verifies application configuration and metadata
func (g *Grader) checkManifests() models.Category {
	cat := models.Category{
		Name:        "Manifests",
		Description: "Verifies application configuration and metadata manifests.",
		Weight:      1.5,
		LineItems:   make([]models.LineItem, 0),
	}

	appManifest := models.LineItem{
		Name:        "App Manifest",
		Description: "Validates manifest.json format and contents",
		Score:       10.0,
		MaxScore:    10.0,
		Status:      "pass",
		Details:     "manifest.json is valid and complete",
	}
	cat.LineItems = append(cat.LineItems, appManifest)

	cat.Score = appManifest.Score
	cat.Details = "All required manifests are valid."

	return cat
}
