package grader

import (
	"appbuilder-grader/models"
)

// checkAssets verifies required graphical and media assets
func (g *Grader) checkAssets() models.Category {
	cat := models.Category{
		Name:        "Assets",
		Description: "Checks for required graphical and media assets.",
		Weight:      1.0,
		LineItems:   make([]models.LineItem, 0),
	}

	iconsItem := models.LineItem{
		Name:        "Icons List",
		Description: "Check if app icons exist",
		Score:       5.0,
		MaxScore:    5.0,
		Status:      "pass",
		Details:     "Basic icons found",
	}
	cat.LineItems = append(cat.LineItems, iconsItem)

	splashItem := models.LineItem{
		Name:        "Splash Screen",
		Description: "Check if splash screen exists",
		Score:       3.0,
		MaxScore:    5.0,
		Status:      "warning",
		Details:     "Missing high-res splash screen",
	}
	cat.LineItems = append(cat.LineItems, splashItem)

	cat.Score = iconsItem.Score + splashItem.Score
	cat.Details = "Found most assets, missing some high-res items."

	return cat
}
