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
	return createIgnoredItem("line_items.menu_navigation_name", "line_items.menu_navigation_desc", 2.0)
}

func (g *Grader) checkKeyboard() models.LineItem {
	return createIgnoredItem("line_items.keyboard_name", "line_items.keyboard_desc", 3.0)
}

func (g *Grader) checkContentsMenu() models.LineItem {
	return createIgnoredItem("line_items.contents_menu_name", "line_items.contents_menu_desc", 1.0)
}

func (g *Grader) checkCustomizedGraphics() models.LineItem {
	return createIgnoredItem("line_items.customized_graphics_name", "line_items.customized_graphics_desc", 1.0)
}

func (g *Grader) checkAppropriateMenuOptions() models.LineItem {
	return createIgnoredItem("line_items.appropriate_menu_options_name", "line_items.appropriate_menu_options_desc", 1.0)
}

func (g *Grader) checkStyleAdjustments() models.LineItem {
	return createIgnoredItem("line_items.style_adjustments_name", "line_items.style_adjustments_desc", 1.0)
}

func (g *Grader) checkTextChanges() models.LineItem {
	return createIgnoredItem("line_items.text_changes_name", "line_items.text_changes_desc", 1.0)
}

func (g *Grader) checkAboutBoxInfo() models.LineItem {
	return createIgnoredItem("line_items.about_box_information_name", "line_items.about_box_information_desc", 1.0)
}

func (g *Grader) checkAboutBoxVernacular() models.LineItem {
	return createIgnoredItem("line_items.about_box_in_vernacular_name", "line_items.about_box_in_vernacular_desc", 1.0)
}
