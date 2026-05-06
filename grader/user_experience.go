package grader

import (
	"appbuilder-grader/models"
)

func (g *Grader) checkUserExperience() models.Category {
	cat := models.Category{
		Name:        "categories.user_experience_name",
		Description: "categories.user_experience_desc",
		Weight:      1.0,
		LineItems:   make([]models.LineItem, 0),
	}

	cat.LineItems = append(cat.LineItems, g.checkMenuNavigation())
	cat.LineItems = append(cat.LineItems, g.checkKeyboard())
	cat.LineItems = append(cat.LineItems, g.checkContentsMenu())
	cat.LineItems = append(cat.LineItems, g.checkCustomizedGraphics())
	cat.LineItems = append(cat.LineItems, g.checkAppropriateMenuOptions())
	cat.LineItems = append(cat.LineItems, g.checkStyleAdjustments())
	cat.LineItems = append(cat.LineItems, g.checkTextChanges())
	cat.LineItems = append(cat.LineItems, g.checkAboutBoxInfo())
	cat.LineItems = append(cat.LineItems, g.checkAboutBoxVernacular())

	return cat
}

func (g *Grader) checkMenuNavigation() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.menu_navigation_name",
		Description: "line_items.menu_navigation_desc",
		MaxScore:    2.0,
	}
	item.Score = 0.0
	item.Status = models.StatusIgnored
	item.Details = "details.not_implemented_yet"
	return item
}

func (g *Grader) checkKeyboard() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.keyboard_name",
		Description: "line_items.keyboard_desc",
		MaxScore:    3.0,
	}
	item.Score = 0.0
	item.Status = models.StatusIgnored
	item.Details = "details.not_implemented_yet"
	return item
}

func (g *Grader) checkContentsMenu() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.contents_menu_name",
		Description: "line_items.contents_menu_desc",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = models.StatusIgnored
	item.Details = "details.not_implemented_yet"
	return item
}

func (g *Grader) checkCustomizedGraphics() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.customized_graphics_name",
		Description: "line_items.customized_graphics_desc",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = models.StatusIgnored
	item.Details = "details.not_implemented_yet"
	return item
}

func (g *Grader) checkAppropriateMenuOptions() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.appropriate_menu_options_name",
		Description: "line_items.appropriate_menu_options_desc",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = models.StatusIgnored
	item.Details = "details.not_implemented_yet"
	return item
}

func (g *Grader) checkStyleAdjustments() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.style_adjustments_name",
		Description: "line_items.style_adjustments_desc",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = models.StatusIgnored
	item.Details = "details.not_implemented_yet"
	return item
}

func (g *Grader) checkTextChanges() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.text_changes_name",
		Description: "line_items.text_changes_desc",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = models.StatusIgnored
	item.Details = "details.not_implemented_yet"
	return item
}

func (g *Grader) checkAboutBoxInfo() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.about_box_information_name",
		Description: "line_items.about_box_information_desc",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = models.StatusIgnored
	item.Details = "details.not_implemented_yet"
	return item
}

func (g *Grader) checkAboutBoxVernacular() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.about_box_in_vernacular_name",
		Description: "line_items.about_box_in_vernacular_desc",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = models.StatusIgnored
	item.Details = "details.not_implemented_yet"
	return item
}
