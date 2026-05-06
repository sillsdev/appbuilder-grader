package grader

import (
	"appbuilder-grader/models"
)

func (g *Grader) checkUserEngagement() models.Category {
	cat := models.Category{
		Name:        "User Engagement",
		Description: "Checks for user engagement and feedback features",
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
		Name:        "Analytics",
		Description: "Analytics being monitored (0=No, 1=Yes)",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}

func (g *Grader) checkTextFeedback() models.LineItem {
	item := models.LineItem{
		Name:        "Text Feedback",
		Description: "Text editable for feedback by email (0=No, 1=Yes)",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}

func (g *Grader) checkFurtherInfoLink() models.LineItem {
	item := models.LineItem{
		Name:        "Further Info Link",
		Description: "Link for further information (0=No, 1=Yes)",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}

func (g *Grader) checkContactDetails() models.LineItem {
	item := models.LineItem{
		Name:        "Contact Details",
		Description: "Contact details for follow-up like phone, email or website (0=No, 1=Yes)",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}

func (g *Grader) checkDeepLinking() models.LineItem {
	item := models.LineItem{
		Name:        "Deep Linking",
		Description: "Deep Linking implemented (0=No, 1=Yes)",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}
