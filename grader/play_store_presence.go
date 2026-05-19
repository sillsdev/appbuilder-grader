package grader

import (
	"appbuilder-grader/models"
	"os"
	"path/filepath"
	"strings"
)

func (g *Grader) checkPlayStorePresence() models.Category {
	cat := models.Category{
		Name:        "categories.play_store_presence_name",
		Description: "categories.play_store_presence_desc",
		Weight:      1.0,
		LineItems:   make([]models.LineItem, 0),
	}

	cat.LineItems = append(cat.LineItems, g.checkAppDescription())
	cat.LineItems = append(cat.LineItems, g.checkLocalizedDescriptionLWC())
	cat.LineItems = append(cat.LineItems, g.checkLocalizedDescriptionVernacular())
	cat.LineItems = append(cat.LineItems, g.checkPhoneScreenshots())
	cat.LineItems = append(cat.LineItems, g.checkTabletScreenshots())
	cat.LineItems = append(cat.LineItems, g.checkEnhancedScreenshots())
	cat.LineItems = append(cat.LineItems, g.checkDemoVideo())

	return cat
}

func (g *Grader) checkAppDescription() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.app_description_name",
		Description: "line_items.app_description_desc",
		MaxScore:    1.0,
		Status:      models.StatusSuggested,
		Details:     "details.app_description_missing",
	}

	for _, listing := range g.playListingLocales() {
		if listing.hasMeaningfulDescription() {
			item.Score = 1.0
			item.Status = models.StatusPass
			item.SetDetails("details.app_description_found", listing.locale)
			return item
		}
	}
	return item
}

func (g *Grader) checkLocalizedDescriptionLWC() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.localized_description_lwc_name",
		Description: "line_items.localized_description_lwc_desc",
		MaxScore:    1.0,
		Status:      models.StatusSuggested,
		Details:     "details.localized_description_lwc_missing",
	}

	primaryBase := g.primaryLanguageBase()
	for _, listing := range g.playListingLocales() {
		base := languageBase(listing.locale)
		if base != "" && base != "en" && base != primaryBase && listing.hasMeaningfulDescription() {
			item.Score = 1.0
			item.Status = models.StatusPass
			item.SetDetails("details.localized_description_lwc_found", listing.locale)
			return item
		}
	}
	return item
}

func (g *Grader) checkLocalizedDescriptionVernacular() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.localized_description_vernacular_name",
		Description: "line_items.localized_description_vernacular_desc",
		MaxScore:    1.0,
		Status:      models.StatusSuggested,
		Details:     "details.localized_description_vernacular_missing",
	}

	primaryBase := g.primaryLanguageBase()
	if primaryBase == "" {
		return item
	}
	for _, listing := range g.playListingLocales() {
		if languageBase(listing.locale) == primaryBase && listing.hasMeaningfulDescription() {
			item.Score = 1.0
			item.Status = models.StatusPass
			item.SetDetails("details.localized_description_vernacular_found", listing.locale)
			return item
		}
	}
	return item
}

func (g *Grader) checkPhoneScreenshots() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.phone_screenshots_name",
		Description: "line_items.phone_screenshots_desc",
		MaxScore:    1.0,
		Status:      models.StatusSuggested,
		Details:     "details.phone_screenshots_missing",
	}

	if count := g.countPlayListingImages("phoneScreenshots"); count > 0 {
		item.Score = 1.0
		item.Status = models.StatusPass
		item.SetDetails("details.phone_screenshots_found", count)
	}
	return item
}

func (g *Grader) checkTabletScreenshots() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.tablet_screenshots_name",
		Description: "line_items.tablet_screenshots_desc",
		MaxScore:    1.0,
		Status:      models.StatusSuggested,
		Details:     "details.tablet_screenshots_missing",
	}

	count := g.countPlayListingImages("sevenInchScreenshots") +
		g.countPlayListingImages("tenInchScreenshots")
	if count > 0 {
		item.Score = 1.0
		item.Status = models.StatusPass
		item.SetDetails("details.tablet_screenshots_found", count)
	}
	return item
}

func (g *Grader) checkEnhancedScreenshots() models.LineItem {
	return createIgnoredItem("line_items.enhanced_screenshots_name", "line_items.enhanced_screenshots_desc", 1.0)
}

func (g *Grader) checkDemoVideo() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.demo_video_name",
		Description: "line_items.demo_video_desc",
		MaxScore:    1.0,
		Status:      models.StatusSuggested,
		Details:     "details.demo_video_missing",
	}

	for _, listing := range g.playListingLocales() {
		video := strings.TrimSpace(readTextFile(filepath.Join(listing.path, "video.txt")))
		if isVideoURL(video) {
			item.Score = 1.0
			item.Status = models.StatusPass
			item.SetDetails("details.demo_video_found", listing.locale)
			return item
		}
	}
	return item
}

type playListingLocale struct {
	locale string
	path   string
}

func (g *Grader) playListingRoot() string {
	dataDir := g.dataDir()
	if dataDir == "" {
		return ""
	}
	root := filepath.Join(dataDir, "publish", "play-listing")
	if isDir(root) {
		return root
	}
	return ""
}

func (g *Grader) playListingLocales() []playListingLocale {
	root := g.playListingRoot()
	if root == "" {
		return nil
	}

	entries, err := os.ReadDir(root)
	if err != nil {
		return nil
	}
	locales := make([]playListingLocale, 0)
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		path := filepath.Join(root, entry.Name())
		if hasPlayListingText(path) || isDir(filepath.Join(path, "images")) {
			locales = append(locales, playListingLocale{locale: entry.Name(), path: path})
		}
	}
	return locales
}

func hasPlayListingText(path string) bool {
	for _, filename := range []string{"title.txt", "short_description.txt", "full_description.txt"} {
		if strings.TrimSpace(readTextFile(filepath.Join(path, filename))) != "" {
			return true
		}
	}
	return false
}

func (listing playListingLocale) hasMeaningfulDescription() bool {
	title := strings.TrimSpace(readTextFile(filepath.Join(listing.path, "title.txt")))
	shortDescription := strings.TrimSpace(readTextFile(filepath.Join(listing.path, "short_description.txt")))
	fullDescription := strings.TrimSpace(readTextFile(filepath.Join(listing.path, "full_description.txt")))
	return title != "" &&
		len([]rune(shortDescription)) >= 20 &&
		len(strings.Fields(fullDescription)) >= 25
}

func (g *Grader) countPlayListingImages(folder string) int {
	total := 0
	for _, listing := range g.playListingLocales() {
		total += countImageFiles(filepath.Join(listing.path, "images", folder))
	}
	return total
}

func countImageFiles(root string) int {
	if !isDir(root) {
		return 0
	}
	count := 0
	_ = filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		if isImageFilename(d.Name()) {
			count++
		}
		return nil
	})
	return count
}

func isImageFilename(filename string) bool {
	switch strings.ToLower(filepath.Ext(filename)) {
	case ".png", ".jpg", ".jpeg", ".webp":
		return true
	default:
		return false
	}
}

func isVideoURL(value string) bool {
	lower := strings.ToLower(strings.TrimSpace(value))
	return strings.HasPrefix(lower, "https://") || strings.HasPrefix(lower, "http://")
}
