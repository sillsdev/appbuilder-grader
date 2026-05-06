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
	return createIgnoredItem("line_items.app_description_name", "line_items.app_description_desc", 1.0)
}

func (g *Grader) checkLocalizedDescriptionLWC() models.LineItem {
	return createIgnoredItem("line_items.localized_description_lwc_name", "line_items.localized_description_lwc_desc", 1.0)
}

func (g *Grader) checkLocalizedDescriptionVernacular() models.LineItem {
	return createIgnoredItem("line_items.localized_description_vernacular_name", "line_items.localized_description_vernacular_desc", 1.0)
}

func (g *Grader) checkPhoneScreenshots() models.LineItem {
	return createIgnoredItem("line_items.phone_screenshots_name", "line_items.phone_screenshots_desc", 1.0)
}

func (g *Grader) checkTabletScreenshots() models.LineItem {
	return createIgnoredItem("line_items.tablet_screenshots_name", "line_items.tablet_screenshots_desc", 1.0)
}

func (g *Grader) checkEnhancedScreenshots() models.LineItem {
	return createIgnoredItem("line_items.enhanced_screenshots_name", "line_items.enhanced_screenshots_desc", 1.0)
}

func (g *Grader) checkDemoVideo() models.LineItem {
	return createIgnoredItem("line_items.demo_video_name", "line_items.demo_video_desc", 1.0)
}
