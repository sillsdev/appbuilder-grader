package grader

import (
	"appbuilder-grader/models"
)

// Old Testament Books:
var AllOTBooks = []string{
	"GEN", "EXO", "LEV", "NUM", "DEU", "JOS", "JDG", "RUT", "1SA", "2SA",
	"1KI", "2KI", "1CH", "2CH", "EZR", "NEH", "EST", "JOB", "PSA", "PRO",
	"ECC", "SNG", "ISA", "JER", "LAM", "EZK", "DAN", "HOS", "JOL", "AMO",
	"OBA", "JON", "MIC", "NAM", "HAB", "ZEP", "HAG", "ZEC", "MAL",
}

// New Testament Books:
var AllNTBooks = []string{
	"MAT", "MRK", "LUK", "JHN", "ACT", "ROM", "1CO", "2CO", "GAL", "EPH",
	"PHP", "COL", "1TH", "2TH", "1TI", "2TI", "TIT", "PHM", "HEB", "JAS",
	"1PE", "2PE", "1JN", "2JN", "3JN", "JUD", "REV",
}

func (g *Grader) checkScriptureTextContent() models.Category {
	cat := models.Category{
		Name:        "categories.scripture_text_content_name",
		Description: "categories.scripture_text_content_desc",
		Weight:      2.0,
		LineItems:   make([]models.LineItem, 0),
	}

	cat.LineItems = append(cat.LineItems, g.checkContentPresence())
	cat.LineItems = append(cat.LineItems, g.checkClickableReferences())
	cat.LineItems = append(cat.LineItems, g.checkMultilingualScripts())
	cat.LineItems = append(cat.LineItems, g.checkRedLetterText())

	return cat
}

func (g *Grader) checkContentPresence() models.LineItem {
	contentItem := models.LineItem{
		Name:        "line_items.content_presence_name",
		Description: "line_items.content_presence_desc",
		MaxScore:    4.0,
	}

	contentItem.Score = 0.0
	contentItem.Status = models.StatusError
	contentItem.Details = "details.no_books_found"

	// Check for book completeness
	includedNTBooks := make(map[string]bool)
	for _, book := range AllNTBooks {
		includedNTBooks[book] = false
	}
	includedOTBooks := make(map[string]bool)
	for _, book := range AllOTBooks {
		includedOTBooks[book] = false
	}

	ntCount := 0
	otCount := 0

	for _, bookCollection := range g.AppDef.Books {
		for _, book := range bookCollection.Book {
			if _, ok := includedNTBooks[book.Id]; ok {
				if !includedNTBooks[book.Id] {
					includedNTBooks[book.Id] = true
					ntCount++
				}
			} else if _, ok := includedOTBooks[book.Id]; ok {
				if !includedOTBooks[book.Id] {
					includedOTBooks[book.Id] = true
					otCount++
				}
			}
		}
	}
	// Check if any books are present
	if len(g.AppDef.Books) > 0 {
		contentItem.Score = 1.0
		contentItem.Status = models.StatusWarning
		contentItem.Details = "details.found_books_not_nt_ot"
		contentItem.DetailsArgs = []any{len(g.AppDef.Books)}
	} else {
		return contentItem
	}

	// Check for portions of NT or OT
	if ntCount > 0 || otCount > 0 {
		contentItem.Score = 1.5
		contentItem.Status = models.StatusPass
		contentItem.Details = "details.found_nt_ot_books"
		contentItem.DetailsArgs = []any{ntCount, otCount}
	} else {
		return contentItem
	}

	// Check for full NT (27 books)
	if ntCount == len(AllNTBooks) {
		contentItem.Score = 2.0
		contentItem.Details = "details.all_nt_no_ot"
	} else {
		return contentItem
	}

	// Check for OT portions
	if otCount > 0 {
		contentItem.Score = 3.0
		contentItem.Details = "details.all_nt_plus_other"
		contentItem.DetailsArgs = []any{len(g.AppDef.Books) - len(AllNTBooks)}
	} else {
		return contentItem
	}

	// Check for full OT (39 books)
	if otCount == len(AllOTBooks) {
		contentItem.Score = 4.0
		contentItem.Details = "details.all_ot_nt"
	} else {
		return contentItem
	}

	return contentItem
}

func (g *Grader) checkClickableReferences() models.LineItem {
	// Placeholder implementation
	// 0=Unlinked
	// 1=Cross-references and parallel passages are linked to text

	return models.LineItem{
		Name:        "line_items.clickable_references_name",
		Description: "line_items.clickable_references_desc",
		Score:       0.0,
		MaxScore:    1.0,
		Status:      models.StatusIgnored,
		Details:     "details.clickable_references_details",
	}
}

func (g *Grader) checkMultilingualScripts() models.LineItem {
	// Placeholder implementation
	// 0=No other texts or scripts
	// 1=Text in additional script(s)
	// 2=Text in other regional language(s)
	// 3=Parallel Back Translation

	return models.LineItem{
		Name:        "line_items.multilingual_scripts_name",
		Description: "line_items.multilingual_scripts_desc",
		Score:       0.0,
		MaxScore:    3.0,
		Status:      models.StatusIgnored,
		Details:     "details.multilingual_scripts_details",
	}
}

func (g *Grader) checkRedLetterText() models.LineItem {
	redLetterItem := models.LineItem{
		Name:        "line_items.red_letter_text_name",
		Description: "line_items.red_letter_text_desc",
		Score:       0.0,
		MaxScore:    1.0,
		Status:      models.StatusWarning,
		Details:     "details.red_letter_text_details",
	}
	// if g.AppDef.Features includes feature with name "show-red-letter"
	for _, feature := range g.AppDef.Features.Feature {
		if feature.Name == "show-red-letters" {
			redLetterItem.Score = 1.0
			redLetterItem.Status = models.StatusPass
			redLetterItem.Details = "details.red_letter_available"
			break
		}
	}
	return redLetterItem
}
