package grader

import (
	"appbuilder-grader/models"
)

func (g *Grader) checkPlayStorePresence() models.Category {
	cat := models.Category{
		Name:        "Play Store Presence",
		Description: "Checks the find-ability and attractiveness of the Play Store listing",
		Weight:      1.0,
		LineItems:   make([]models.LineItem, 0),
	}

	cat.LineItems = append(cat.LineItems, g.checkAppDescription())
	cat.LineItems = append(cat.LineItems, g.checkLocalizedDescriptionLWC())
	cat.LineItems = append(cat.LineItems, g.checkLocalizedDescriptionVernacular())
	cat.LineItems = append(cat.LineItems, g.checkPhoneScreenshots())
	cat.LineItems = append(cat.LineItems, g.checkTabletScreenshots())
	cat.LineItems = append(cat.LineItems, g.checkEnhancedScreenshots())
	cat.LineItems = append(cat.LineItems, g.checkDemoVideo())

	return cat
}

func (g *Grader) checkAppDescription() models.LineItem {
	item := models.LineItem{
		Name:        "App Description",
		Description: "Helpful description of App and good range of searchable key words (0=No, 1=Yes)",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}

func (g *Grader) checkLocalizedDescriptionLWC() models.LineItem {
	item := models.LineItem{
		Name:        "Localized Description LWC",
		Description: "Localized description in LWC (0=No, 1=Yes)",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}

func (g *Grader) checkLocalizedDescriptionVernacular() models.LineItem {
	item := models.LineItem{
		Name:        "Localized Description Vernacular",
		Description: "Localised description in vernacular (0=No, 1=Yes)",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}

func (g *Grader) checkPhoneScreenshots() models.LineItem {
	item := models.LineItem{
		Name:        "Phone Screenshots",
		Description: "Phone screenshots available in listing (0=No, 1=Yes)",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}

func (g *Grader) checkTabletScreenshots() models.LineItem {
	item := models.LineItem{
		Name:        "Tablet Screenshots",
		Description: "Tablet screenshots available in listing (0=No, 1=Yes)",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}

func (g *Grader) checkEnhancedScreenshots() models.LineItem {
	item := models.LineItem{
		Name:        "Enhanced Screenshots",
		Description: "Enhanced/edited screenshots highlighting usage/features (0=No, 1=Yes)",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}

func (g *Grader) checkDemoVideo() models.LineItem {
	item := models.LineItem{
		Name:        "Demo Video",
		Description: "Demo video in PlayStore on how to use features (0=No, 1=Yes)",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}
