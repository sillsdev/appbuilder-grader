package grader

import (
	"appbuilder-grader/models"
)

func (g *Grader) checkMiscellaneous() models.Category {
	cat := models.Category{
		Name:        "Thinking Beyond",
		Description: "Checks for availability outside just the standard app",
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
		Name:        "Web Site Text",
		Description: "Text available through web site (0=None, 1=Online, 2=Online with synchronized audio)",
		MaxScore:    2.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}

func (g *Grader) checkDBLText() models.LineItem {
	item := models.LineItem{
		Name:        "DBL Text",
		Description: "Text in DBL available for other Paratext users (0=No, 1=Yes)",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}

func (g *Grader) checkYouVersionText() models.LineItem {
	item := models.LineItem{
		Name:        "YouVersion Text",
		Description: "Text available in YouVersion app (0=No, 1=Yes)",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}
