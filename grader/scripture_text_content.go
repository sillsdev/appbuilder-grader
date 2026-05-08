package grader

import (
	"appbuilder-grader/models"
	"fmt"
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
	item := models.LineItem{
		Name:        "line_items.clickable_references_name",
		Description: "line_items.clickable_references_desc",
		MaxScore:    1.0,
		Status:      models.StatusWarning,
		Details:     "details.clickable_references_missing",
	}

	bookFiles := g.bookFiles()
	if len(bookFiles) == 0 {
		item.Status = models.StatusError
		item.SetDetails("details.no_book_files_found")
		return item
	}

	linkedFiles := 0
	nonStandardFiles := 0
	for _, bookFile := range bookFiles {
		// Only check for clickable references in files that contain scripture text (i.e., not just front/back matter)
		if isNTBook(bookFile.BookID) || isOTBook(bookFile.BookID) {
			if hasClickableReference(readTextFile(bookFile.Path)) {
				linkedFiles++
			}
		} else {
			nonStandardFiles++
		}
	}

	percent := (float64(linkedFiles) / float64(len(bookFiles) - nonStandardFiles)) * 100
	item.SetDetails("details.clickable_references_coverage", linkedFiles, len(bookFiles) - nonStandardFiles, fmt.Sprintf("%.1f", percent))
	if percent > 60 {
		item.Score = 1.0
		item.Status = models.StatusPass
	}
	return item
}

func (g *Grader) checkMultilingualScripts() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.multilingual_scripts_name",
		Description: "line_items.multilingual_scripts_desc",
		MaxScore:    3.0,
		Status:      models.StatusWarning,
		Details:     "details.multilingual_scripts_none",
	}

	collections := g.AppDef.Books
	if len(collections) <= 1 {
		return item
	}

	baseLanguages := make(map[string]bool)
	scriptsByBase := make(map[string]map[string]bool)
	hasAdditionalScript := false
	hasRegionalLanguage := false
	hasBackTranslation := false
	primaryBase := languageBase(collections[0].WritingSystem.Code)

	for _, collection := range collections {
		base := languageBase(collection.WritingSystem.Code)
		if base == "" {
			base = languageBase(collection.Id)
		}
		if base == "" {
			continue
		}

		baseLanguages[base] = true
		if scriptsByBase[base] == nil {
			scriptsByBase[base] = make(map[string]bool)
		}
		scriptsByBase[base][strings.ToLower(collection.WritingSystem.Code)] = true

		if base == primaryBase && len(scriptsByBase[base]) > 1 {
			hasAdditionalScript = true
		}
		if base != primaryBase && base != "eng" && base != "en" {
			hasRegionalLanguage = true
		}
		if isBackTranslationCollection(collection.Id, collection.BookCollectionName, collection.WritingSystem.Code) {
			hasBackTranslation = true
		}
	}

	switch {
	case hasBackTranslation:
		item.Score = 3.0
		item.Status = models.StatusPass
		item.SetDetails("details.multilingual_scripts_back_translation", len(collections))
	case hasRegionalLanguage:
		item.Score = 2.0
		item.Status = models.StatusPass
		item.SetDetails("details.multilingual_scripts_regional_language", len(baseLanguages), len(collections))
	case hasAdditionalScript:
		item.Score = 1.0
		item.Status = models.StatusPass
		item.SetDetails("details.multilingual_scripts_additional_scripts", len(scriptsByBase[primaryBase]), len(collections))
	}

	return item
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

	if !g.hasFeature("show-red-letters") {
		redLetterItem.SetDetails("details.red_letter_feature_missing")
		return redLetterItem
	}

	ntFilesWithWJ := 0
	for _, bookFile := range g.bookFiles() {
		if !isNTBook(bookFile.BookID) {
			continue
		}
		if strings.Contains(readTextFile(bookFile.Path), `\wj`) {
			ntFilesWithWJ++
		}
	}
	if ntFilesWithWJ > 0 {
		redLetterItem.Score = 1.0
		redLetterItem.Status = models.StatusPass
		redLetterItem.SetDetails("details.red_letter_available", ntFilesWithWJ)
	} else {
		redLetterItem.SetDetails("details.red_letter_markers_missing")
	}
	return redLetterItem
}

func hasClickableReference(content string) bool {
	return strings.Contains(content, `\x `) && (strings.Contains(content, `\xt`) || strings.Contains(content, `\ref`)) ||
		strings.Contains(content, `\ref`) && strings.Contains(content, `|`) && strings.Contains(content, `\ref*`)
}

func languageBase(code string) string {
	code = strings.TrimSpace(strings.ToLower(code))
	if code == "" {
		return ""
	}
	if idx := strings.IndexAny(code, "-_"); idx >= 0 {
		return code[:idx]
	}
	return code
}

func isBackTranslationCollection(id, name, writingSystem string) bool {
	haystack := strings.ToLower(id + " " + name + " " + writingSystem)
	return strings.Contains(haystack, "back translation") ||
		strings.Contains(haystack, "backtranslation") ||
		strings.Contains(haystack, "bt") ||
		strings.Contains(haystack, "english with")
}

func isNTBook(bookID string) bool {
	bookID = strings.ToUpper(bookID)
	for _, ntBook := range AllNTBooks {
		if bookID == ntBook {
			return true
		}
	}
	return false
}

func isOTBook(bookID string) bool {
	bookID = strings.ToUpper(bookID)
	for _, otBook := range AllOTBooks {
		if bookID == otBook {
			return true
		}
	}
	return false
}

func (g *Grader) hasFeature(name string) bool {
	if hasFeature(g.AppDef.Features.Feature, name) {
		return true
	}
	for _, collection := range g.AppDef.Books {
		if hasFeature(collection.Features.Feature, name) {
			return true
		}
		for _, book := range collection.Book {
			if hasFeature(book.Features.Feature, name) {
				return true
			}
		}
	}
	return false
}

func hasFeature(features []models.Feature, name string) bool {
	for _, feature := range features {
		if feature.Name == name {
			return true
		}
	}
	return false
}
