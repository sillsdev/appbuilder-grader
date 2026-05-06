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
	item := models.LineItem{
		Name:        "line_items.shareable_one_to_one_name",
		Description: "line_items.shareable_one_to_one_desc",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = models.StatusIgnored
	item.Details = "details.not_implemented_yet"
	return item
}

func (g *Grader) checkSDCardAvailability() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.sd_card_availability_name",
		Description: "line_items.sd_card_availability_desc",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = models.StatusIgnored
	item.Details = "details.not_implemented_yet"
	return item
}

func (g *Grader) checkDirectLinks() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.direct_links_name",
		Description: "line_items.direct_links_desc",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = models.StatusIgnored
	item.Details = "details.not_implemented_yet"
	return item
}

func (g *Grader) checkPlayStorePublishing() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.play_store_publishing_name",
		Description: "line_items.play_store_publishing_desc",
		MaxScore:    2.0,
	}
	item.Score = 0.0
	item.Status = models.StatusIgnored
	item.Details = "details.not_implemented_yet"
	return item
}
