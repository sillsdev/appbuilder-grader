package grader

import (
	"appbuilder-grader/models"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
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
		Status:      models.StatusSuggested,
		Details:     "details.menu_navigation_english_only",
	}

	primaryBase := g.primaryLanguageBase()
	regionalMappings := 0
	vernacularMappings := 0
	for _, mapping := range g.AppDef.TranslationMappings.TranslationMapping {
		if !isNavigationMapping(mapping.Id) {
			continue
		}
		for _, translation := range mapping.Translation {
			lang := strings.TrimSpace(strings.ToLower(translation.Lang))
			if lang == "" || lang == "en" || strings.TrimSpace(translation.Value) == "" {
				continue
			}
			if languageBase(lang) == primaryBase {
				vernacularMappings++
			} else {
				regionalMappings++
			}
		}
	}

	if regionalMappings > 0 {
		item.Score = 1.0
		item.Status = models.StatusPass
		item.SetDetails("details.menu_navigation_regional", regionalMappings)
	}
	if vernacularMappings > 0 {
		item.Score = 2.0
		item.Status = models.StatusPass
		item.SetDetails("details.menu_navigation_vernacular", vernacularMappings)
	}
	return item
}

func (g *Grader) checkKeyboard() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.keyboard_name",
		Description: "line_items.keyboard_desc",
		MaxScore:    3.0,
		Status:      models.StatusSuggested,
		Details:     "details.keyboard_missing",
	}

	appDef := strings.ToLower(g.appDefContent())
	hasInputButtons := strings.Contains(appDef, `feature name="input-buttons"`) &&
		!strings.Contains(appDef, `feature name="input-buttons" value=""`) &&
		!strings.Contains(appDef, `feature name="input-buttons" value="false"`)
	hasKeyman := hasVernacularKeymanKeyboard(appDef)
	hasPredictive := strings.Contains(appDef, "predictive") || strings.Contains(appDef, "word-prediction")

	if hasInputButtons {
		item.Score = 1.0
		item.Status = models.StatusPass
		item.SetDetails("details.keyboard_input_buttons")
	}
	if hasKeyman {
		item.Score = 2.0
		item.Status = models.StatusPass
		item.SetDetails("details.keyboard_keyman")
	}
	if hasPredictive {
		item.Score = 3.0
		item.Status = models.StatusPass
		item.SetDetails("details.keyboard_predictive")
	}
	return item
}

func (g *Grader) checkContentsMenu() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.contents_menu_name",
		Description: "line_items.contents_menu_desc",
		MaxScore:    1.0,
		Status:      models.StatusSuggested,
		Details:     "details.contents_menu_missing",
	}
	if g.hasTranslationMapping("Menu_Contents") ||
		g.featureValue("book-select") != "" ||
		g.hasTranslationMapping("Menu_Layout") {
		item.Score = 1.0
		item.Status = models.StatusPass
		item.SetDetails("details.contents_menu_found")
	}
	return item
}

func (g *Grader) checkCustomizedGraphics() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.customized_graphics_name",
		Description: "line_items.customized_graphics_desc",
		MaxScore:    1.0,
		Status:      models.StatusSuggested,
		Details:     "details.customized_graphics_missing",
	}

	graphics := 0
	for _, images := range g.AppDef.Images {
		switch strings.ToLower(images.Type) {
		case "launcher", "splash", "drawer", "background":
			graphics += len(images.Image)
		}
	}
	if graphics > 0 || g.hasFeatureValue("splash-screen", "true") {
		item.Score = 1.0
		item.Status = models.StatusPass
		item.SetDetails("details.customized_graphics_found", graphics)
	}
	return item
}

func (g *Grader) checkAppropriateMenuOptions() models.LineItem {
	return createIgnoredItem("line_items.appropriate_menu_options_name", "line_items.appropriate_menu_options_desc", 1.0)
}

func (g *Grader) checkStyleAdjustments() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.style_adjustments_name",
		Description: "line_items.style_adjustments_desc",
		MaxScore:    1.0,
		Status:      models.StatusSuggested,
		Details:     "details.style_adjustments_missing",
	}
	appDef := g.appDefContent()
	styleCount := strings.Count(appDef, "<style ")
	fontCount := strings.Count(appDef, "<font ")
	if styleCount > 0 || fontCount > 0 || len(g.AppDef.Colors) > 0 {
		item.Score = 1.0
		item.Status = models.StatusPass
		item.SetDetails("details.style_adjustments_found", styleCount, fontCount)
	}
	return item
}

func (g *Grader) checkTextChanges() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.text_changes_name",
		Description: "line_items.text_changes_desc",
		MaxScore:    1.0,
		Status:      models.StatusSuggested,
		Details:     "details.text_changes_missing",
	}

	changes := changeBlocks(g.appDefContent())
	for _, change := range changes {
		lower := strings.ToLower(change)
		if strings.Contains(lower, "keep") ||
			strings.Contains(lower, "nbsp") ||
			strings.Contains(lower, "no break") ||
			strings.Contains(lower, `\u00a0`) ||
			strings.Contains(lower, "~") {
			item.Score = 1.0
			item.Status = models.StatusPass
			item.SetDetails("details.text_changes_found", len(changes))
			return item
		}
	}
	return item
}

func (g *Grader) checkAboutBoxInfo() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.about_box_information_name",
		Description: "line_items.about_box_information_desc",
		MaxScore:    1.0,
		Status:      models.StatusSuggested,
		Details:     "details.about_box_information_missing",
	}

	content := g.aboutContent()
	lower := strings.ToLower(content)
	hasCopyright := strings.Contains(lower, "copyright") || strings.Contains(content, "©")
	hasContactOrVersion := strings.Contains(lower, "contact") ||
		strings.Contains(lower, "mailto:") ||
		strings.Contains(lower, "http://") ||
		strings.Contains(lower, "https://") ||
		strings.Contains(lower, "version")
	if g.aboutEnabled() && hasCopyright && hasContactOrVersion {
		item.Score = 1.0
		item.Status = models.StatusPass
		item.SetDetails("details.about_box_information_found")
	}
	return item
}

func (g *Grader) checkAboutBoxVernacular() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.about_box_in_vernacular_name",
		Description: "line_items.about_box_in_vernacular_desc",
		MaxScore:    1.0,
		Status:      models.StatusSuggested,
		Details:     "details.about_box_vernacular_missing",
	}

	nonLatin, totalLetters := nonLatinLetterStats(g.aboutContent())
	if nonLatin >= 80 && totalLetters > 0 && float64(nonLatin)/float64(totalLetters) >= 0.10 {
		item.Score = 1.0
		item.Status = models.StatusPass
		item.SetDetails("details.about_box_vernacular_found", nonLatin)
	}
	return item
}

func (g *Grader) primaryLanguageBase() string {
	for _, collection := range g.AppDef.Books {
		if base := languageBase(collection.WritingSystem.Code); base != "" {
			return base
		}
		if base := languageBase(collection.Id); base != "" {
			return base
		}
	}
	for _, ws := range g.AppDef.InterfaceLanguages.WritingSystems.WritingSystem {
		if base := languageBase(ws.Code); base != "" && base != "en" {
			return base
		}
	}
	return ""
}

func (g *Grader) hasTranslationMapping(id string) bool {
	for _, mapping := range g.AppDef.TranslationMappings.TranslationMapping {
		if mapping.Id == id {
			return true
		}
	}
	return false
}

func (g *Grader) featureValue(name string) string {
	if value, ok := featureValue(g.AppDef.Features.Feature, name); ok {
		return value
	}
	for _, collection := range g.AppDef.Books {
		if value, ok := featureValue(collection.Features.Feature, name); ok {
			return value
		}
		for _, book := range collection.Book {
			if value, ok := featureValue(book.Features.Feature, name); ok {
				return value
			}
		}
	}
	return ""
}

func (g *Grader) hasFeatureValue(name, value string) bool {
	return strings.EqualFold(g.featureValue(name), value)
}

func (g *Grader) aboutEnabled() bool {
	return strings.EqualFold(g.AppDef.About.Enabled, "true") &&
		strings.TrimSpace(g.AppDef.About.Filename) != ""
}

func (g *Grader) aboutContent() string {
	if !g.aboutEnabled() {
		return ""
	}
	return readTextFile(g.resolveDataFile("about", filepath.FromSlash(g.AppDef.About.Filename)))
}

func featureValue(features []models.Feature, name string) (string, bool) {
	for _, feature := range features {
		if feature.Name == name {
			return feature.Value, true
		}
	}
	return "", false
}

func isNavigationMapping(id string) bool {
	return strings.HasPrefix(id, "Menu_") ||
		strings.HasPrefix(id, "Settings_") ||
		strings.HasPrefix(id, "Button_")
}

func hasVernacularKeymanKeyboard(appDef string) bool {
	keyboardBlocks := regexp.MustCompile(`(?s)<keyboard>.*?</keyboard>`).FindAllString(appDef, -1)
	for _, block := range keyboardBlocks {
		if !strings.Contains(block, ".kmp") {
			continue
		}
		if strings.Contains(block, `langid" value="en"`) ||
			strings.Contains(block, "basic_kbdus") ||
			strings.Contains(block, "us basic") {
			continue
		}
		return true
	}
	return false
}

func changeBlocks(appDef string) []string {
	return regexp.MustCompile(`(?s)<change>.*?</change>`).FindAllString(appDef, -1)
}

func nonLatinLetterStats(content string) (int, int) {
	nonLatin := 0
	totalLetters := 0
	for _, r := range content {
		if !unicode.IsLetter(r) {
			continue
		}
		totalLetters++
		if r > unicode.MaxASCII {
			nonLatin++
		}
	}
	return nonLatin, totalLetters
}
