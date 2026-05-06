package grader

import (
	"appbuilder-grader/models"
)

func (g *Grader) checkPlayStorePresence() models.Category {
	cat := models.Category{
		Name:        "categories.play_store_presence_name",
		Description: "categories.play_store_presence_desc",
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
		Name:        "line_items.app_description_name",
		Description: "line_items.app_description_desc",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "details.not_implemented_yet"
	return item
}

func (g *Grader) checkLocalizedDescriptionLWC() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.localized_description_lwc_name",
		Description: "line_items.localized_description_lwc_desc",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "details.not_implemented_yet"
	return item
}

func (g *Grader) checkLocalizedDescriptionVernacular() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.localized_description_vernacular_name",
		Description: "line_items.localized_description_vernacular_desc",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "details.not_implemented_yet"
	return item
}

func (g *Grader) checkPhoneScreenshots() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.phone_screenshots_name",
		Description: "line_items.phone_screenshots_desc",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "details.not_implemented_yet"
	return item
}

func (g *Grader) checkTabletScreenshots() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.tablet_screenshots_name",
		Description: "line_items.tablet_screenshots_desc",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "details.not_implemented_yet"
	return item
}

func (g *Grader) checkEnhancedScreenshots() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.enhanced_screenshots_name",
		Description: "line_items.enhanced_screenshots_desc",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "details.not_implemented_yet"
	return item
}

func (g *Grader) checkDemoVideo() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.demo_video_name",
		Description: "line_items.demo_video_desc",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "details.not_implemented_yet"
	return item
}
