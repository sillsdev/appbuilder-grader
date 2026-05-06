package grader

import (
	"appbuilder-grader/models"
)

func (g *Grader) checkMiscellaneous() models.Category {
	cat := models.Category{
		Name:        "categories.thinking_beyond_name",
		Description: "categories.thinking_beyond_desc",
		Weight:      1.0,
		LineItems:   make([]models.LineItem, 0),
	}

	cat.LineItems = append(cat.LineItems, g.checkWebSiteText())
	cat.LineItems = append(cat.LineItems, g.checkDBLText())
	cat.LineItems = append(cat.LineItems, g.checkYouVersionText())

	return cat
}

func (g *Grader) checkWebSiteText() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.web_site_text_name",
		Description: "line_items.web_site_text_desc",
		MaxScore:    2.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "details.not_implemented_yet"
	return item
}

func (g *Grader) checkDBLText() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.dbl_text_name",
		Description: "line_items.dbl_text_desc",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "details.not_implemented_yet"
	return item
}

func (g *Grader) checkYouVersionText() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.youversion_text_name",
		Description: "line_items.youversion_text_desc",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "details.not_implemented_yet"
	return item
}
