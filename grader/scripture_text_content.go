package grader

import (
	"appbuilder-grader/models"
	"strings"
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

// NOTE: Do we want all of these?
// Miscellaneous books:
var AllMiscBooks = []string{
	"TOB", "JDT", "ESG", "WIS", "SIR", "BAR", "LJE", "S3Y", "SUS", "BEL",
	"1MA", "2MA", "3MA", "4MA", "1ES", "2ES", "MAN", "PS2", "ODA", "PSS",
	"EZA", "5EZ", "6EZ", "DAG", "PS3", "2BA", "LBA", "JUB", "ENO", "1MQ",
	"2MQ", "3MQ", "REP", "4BA", "LAO", "FRT", "BAK", "OTH", "INT", "CNC",
	"GLO", "TDX", "NDX", "XXA", "XXB", "XXC", "XXD", "XXE", "XXF", "XXG",
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
	contentItem.SetDetails("details.no_books_found")

	// Check for book completeness
	includedNTBooks := make(map[string]bool)
	for _, book := range AllNTBooks {
		includedNTBooks[book] = false
	}
	includedOTBooks := make(map[string]bool)
	for _, book := range AllOTBooks {
		includedOTBooks[book] = false
	}
	includedMiscBooks := make(map[string]bool)
	for _, book := range AllMiscBooks {
		includedMiscBooks[book] = false
	}

	ntCount := 0
	otCount := 0
	miscCount := 0
	unknownCount := 0
	unknownBooks := make(map[string]bool)

	for _, bookCollection := range g.AppDef.Books {
		for _, bookOriginal := range bookCollection.Book {
			book := strings.ToUpper(bookOriginal.Id)
			if _, ok := includedNTBooks[book]; ok {
				if !includedNTBooks[book] {
					includedNTBooks[book] = true
					ntCount++
				}
			} else if _, ok := includedOTBooks[book]; ok {
				if !includedOTBooks[book] {
					includedOTBooks[book] = true
					otCount++
				}
			} else if _, ok := includedMiscBooks[book]; ok {
				if !includedMiscBooks[book] {
					includedMiscBooks[book] = true
					miscCount++
				}
			} else {
				if !unknownBooks[book] {
					unknownBooks[book] = true
					unknownCount++
				}
			}
		}
	}

	otherBooksList := make([]string, 0, len(unknownBooks))
	for book := range unknownBooks {
		otherBooksList = append(otherBooksList, book)
	}
	otherBooks := strings.Join(otherBooksList, ", ")

	// Check if any books are present
	if len(g.AppDef.Books) > 0 {
		contentItem.Score = 1.0
		contentItem.Status = models.StatusWarning
		contentItem.SetDetails("details.found_books_not_nt_ot", len(g.AppDef.Books), otherBooks)
	} else {
		return contentItem
	}

	// Check for portions of NT or OT
	if ntCount > 0 || otCount > 0 {
		contentItem.Score = 1.5
		contentItem.Status = models.StatusPass
		contentItem.SetDetails("details.found_nt_ot_books", ntCount, otCount, otherBooks)
	} else {
		return contentItem
	}

	// Check for full NT (27 books)
	if ntCount == len(AllNTBooks) {
		contentItem.Score = 2.0
		contentItem.SetDetails("details.all_nt_no_ot", otherBooks)
	} else {
		return contentItem
	}

	// Check for OT portions
	if otCount > 0 {
		contentItem.Score = 3.0
		contentItem.SetDetails("details.all_nt_plus_other", otCount, miscCount, unknownCount, otherBooks)

	} else {
		return contentItem
	}

	// Check for full OT (39 books)
	if otCount == len(AllOTBooks) {
		contentItem.Score = 4.0
		contentItem.SetDetails("details.all_ot_nt", miscCount)
	} else {
		return contentItem
	}

	if unknownCount > 0 {
		contentItem.Score = 3.5
		contentItem.SetDetails("details.all_ot_nt_plus_other", miscCount, unknownCount, otherBooks)
	}

	return contentItem
}

func (g *Grader) checkClickableReferences() models.LineItem {
	// Placeholder implementation
	// 0=Unlinked
	// 1=Cross-references and parallel passages are linked to text
	return createIgnoredItem("line_items.clickable_references_name", "line_items.clickable_references_desc", 1.0)
}

func (g *Grader) checkMultilingualScripts() models.LineItem {
	// Placeholder implementation
	// 0=No other texts or scripts
	// 1=Text in additional script(s)
	// 2=Text in other regional language(s)
	// 3=Parallel Back Translation
	return createIgnoredItem("line_items.multilingual_scripts_name", "line_items.multilingual_scripts_desc", 3.0)
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
