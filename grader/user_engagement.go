package grader

import (
	"appbuilder-grader/models"
)

func (g *Grader) checkUserEngagement() models.Category {
	cat := models.Category{
		Name:        "categories.user_engagement_name",
		Description: "categories.user_engagement_desc",
		Weight:      1.0,
		LineItems:   make([]models.LineItem, 0),
	}

	cat.LineItems = append(cat.LineItems, g.checkAnalytics())
	cat.LineItems = append(cat.LineItems, g.checkTextFeedback())
	cat.LineItems = append(cat.LineItems, g.checkFurtherInfoLink())
	cat.LineItems = append(cat.LineItems, g.checkContactDetails())
	cat.LineItems = append(cat.LineItems, g.checkDeepLinking())

	return cat
}

func (g *Grader) checkAnalytics() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.analytics_name",
		Description: "line_items.analytics_desc",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "details.not_implemented_yet"
	return item
}

func (g *Grader) checkTextFeedback() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.text_feedback_name",
		Description: "line_items.text_feedback_desc",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "details.not_implemented_yet"
	return item
}

func (g *Grader) checkFurtherInfoLink() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.further_info_link_name",
		Description: "line_items.further_info_link_desc",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "details.not_implemented_yet"
	return item
}

func (g *Grader) checkContactDetails() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.contact_details_name",
		Description: "line_items.contact_details_desc",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "details.not_implemented_yet"
	return item
}

func (g *Grader) checkDeepLinking() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.deep_linking_name",
		Description: "line_items.deep_linking_desc",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "details.not_implemented_yet"
	return item
}
