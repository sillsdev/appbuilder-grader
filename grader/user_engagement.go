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
	return createIgnoredItem("line_items.analytics_name", "line_items.analytics_desc", 1.0)
}

func (g *Grader) checkTextFeedback() models.LineItem {
	return createIgnoredItem("line_items.text_feedback_name", "line_items.text_feedback_desc", 1.0)
}

func (g *Grader) checkFurtherInfoLink() models.LineItem {
	return createIgnoredItem("line_items.further_info_link_name", "line_items.further_info_link_desc", 1.0)
}

func (g *Grader) checkContactDetails() models.LineItem {
	return createIgnoredItem("line_items.contact_details_name", "line_items.contact_details_desc", 1.0)
}

func (g *Grader) checkDeepLinking() models.LineItem {
	return createIgnoredItem("line_items.deep_linking_name", "line_items.deep_linking_desc", 1.0)
}
