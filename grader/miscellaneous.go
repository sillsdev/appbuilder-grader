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
	// Check feature export-html-pwa
	item := models.LineItem{
		Name:        "line_items.web_site_text_name",
		Description: "line_items.web_site_text_desc",
		MaxScore:    1.0,
		Status:      models.StatusSuggested,
		Details:     "details.web_site_text_missing",
	}

	if g.hasFeatureValue("export-html-pwa", "true") {
		item.Score = 1.0
		item.Status = models.StatusPass
		item.SetDetails("details.web_site_text_present")
	}

	return item
}

func (g *Grader) checkDBLText() models.LineItem {
	return createIgnoredItem("line_items.dbl_text_name", "line_items.dbl_text_desc", 1.0)
}

func (g *Grader) checkYouVersionText() models.LineItem {
	return createIgnoredItem("line_items.youversion_text_name", "line_items.youversion_text_desc", 1.0)
}
