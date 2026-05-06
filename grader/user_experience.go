package grader

import (
	"appbuilder-grader/models"
)

func (g *Grader) checkUserExperience() models.Category {
	cat := models.Category{
		Name:        "User Experience",
		Description: "Checks for enhanced app features like menus, keyboard, and graphics",
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
		Name:        "Menu Navigation",
		Description: "Checks navigation localization (0=English, 1=Regional language, 2=Local vernacular)",
		MaxScore:    2.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}

func (g *Grader) checkKeyboard() models.LineItem {
	item := models.LineItem{
		Name:        "Keyboard",
		Description: "Keyboard for search/notes (0=No vernacular, 1=Vernacular buttons, 2=Embedded Keyman, 3=Predictive)",
		MaxScore:    3.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}

func (g *Grader) checkContentsMenu() models.LineItem {
	item := models.LineItem{
		Name:        "Contents Menu",
		Description: "Contents menu for books/sections/languages (0=No, 1=Yes)",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}

func (g *Grader) checkCustomizedGraphics() models.LineItem {
	item := models.LineItem{
		Name:        "Customized Graphics",
		Description: "Customized splash screen, drawer, interface icon (0=No, 1=Yes)",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}

func (g *Grader) checkAppropriateMenuOptions() models.LineItem {
	item := models.LineItem{
		Name:        "Appropriate Menu Options",
		Description: "Appropriate options available on menus (0=No, 1=Yes)",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}

func (g *Grader) checkStyleAdjustments() models.LineItem {
	item := models.LineItem{
		Name:        "Style Adjustments",
		Description: "Appropriate adjustments to Styles to improve app (0=No, 1=Yes)",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}

func (g *Grader) checkTextChanges() models.LineItem {
	item := models.LineItem{
		Name:        "Text Changes",
		Description: "Appropriate changes made to underlying text (0=No, 1=Enhancements defined)",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}

func (g *Grader) checkAboutBoxInfo() models.LineItem {
	item := models.LineItem{
		Name:        "About Box Information",
		Description: "Appropriate information in About... box copyright, etc. (0=No, 1=Yes)",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}

func (g *Grader) checkAboutBoxVernacular() models.LineItem {
	item := models.LineItem{
		Name:        "About Box in Vernacular",
		Description: "About box information in vernacular language (not just English) (0=No, 1=Yes)",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}
