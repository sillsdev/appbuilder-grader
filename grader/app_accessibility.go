package grader

import (
	"appbuilder-grader/models"
)

func (g *Grader) checkAccessibility() models.Category {
	cat := models.Category{
		Name:        "categories.app_accessibility_name",
		Description: "categories.app_accessibility_desc",
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
	return createIgnoredItem("line_items.shareable_one_to_one_name", "line_items.shareable_one_to_one_desc", 1.0)
}

func (g *Grader) checkSDCardAvailability() models.LineItem {
	return createIgnoredItem("line_items.sd_card_availability_name", "line_items.sd_card_availability_desc", 1.0)
}

func (g *Grader) checkDirectLinks() models.LineItem {
	return createIgnoredItem("line_items.direct_links_name", "line_items.direct_links_desc", 1.0)
}

func (g *Grader) checkPlayStorePublishing() models.LineItem {
	return createIgnoredItem("line_items.play_store_publishing_name", "line_items.play_store_publishing_desc", 2.0)
}
