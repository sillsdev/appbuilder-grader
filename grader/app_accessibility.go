package grader

import (
	"appbuilder-grader/models"
)

func (g *Grader) checkAccessibility() models.Category {
	cat := models.Category{
		Name:        "App Accessibility",
		Description: "Checks how the app is shared and distributed",
		Weight:      1.0,
		LineItems:   make([]models.LineItem, 0),
	}

	cat.LineItems = append(cat.LineItems, g.checkShareable())
	cat.LineItems = append(cat.LineItems, g.checkSDCardAvailability())
	cat.LineItems = append(cat.LineItems, g.checkDirectLinks())
	cat.LineItems = append(cat.LineItems, g.checkPlayStorePublishing())

	return cat
}

func (g *Grader) checkShareable() models.LineItem {
	item := models.LineItem{
		Name:        "Shareable One-to-One",
		Description: "App is shareable one-to-one (0=No, 1=Yes)",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}

func (g *Grader) checkSDCardAvailability() models.LineItem {
	item := models.LineItem{
		Name:        "SD Card Availability",
		Description: "Available on SD cards locally (0=No, 1=Yes)",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}

func (g *Grader) checkDirectLinks() models.LineItem {
	item := models.LineItem{
		Name:        "Direct Links",
		Description: "Available directly through other websites/links (0=No, 1=Yes)",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}

func (g *Grader) checkPlayStorePublishing() models.LineItem {
	item := models.LineItem{
		Name:        "Play Store Publishing",
		Description: "Published on Google PlayStore (0=No, 1=Local account, 2=SIL/Kalaam media)",
		MaxScore:    2.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}
