package grader

import (
	"appbuilder-grader/models"
	"regexp"
	"strconv"
	"strings"
)

const latestAppBuilderVersion = "14.0"

func (g *Grader) checkAccessibility() models.Category {
	cat := models.Category{
		Name:        "categories.app_accessibility_name",
		Description: "categories.app_accessibility_desc",
		Weight:      1.0,
		LineItems:   make([]models.LineItem, 0),
	}

	cat.LineItems = append(cat.LineItems, g.checkAppBuilderVersion())
	cat.LineItems = append(cat.LineItems, g.checkShareable())
	cat.LineItems = append(cat.LineItems, g.checkSDCardAvailability())
	cat.LineItems = append(cat.LineItems, g.checkDirectLinks())
	cat.LineItems = append(cat.LineItems, g.checkPlayStorePublishing())

	return cat
}

func (g *Grader) checkAppBuilderVersion() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.appbuilder_version_name",
		Description: "line_items.appbuilder_version_desc",
		MaxScore:    5.0,
		Status:      models.StatusWarning,
		Details:     "details.appbuilder_version_missing",
	}

	projectVersion, ok := parseVersionParts(g.AppDef.ProgramVersion)
	if !ok {
		return item
	}
	latestVersion, _ := parseVersionParts(latestAppBuilderVersion)

	majorDiff := latestVersion.major - projectVersion.major
	switch {
	case compareVersionParts(projectVersion, latestVersion) >= 0:
		item.Score = 5
		item.Status = models.StatusPass
		item.SetDetails("details.appbuilder_version_latest", g.AppDef.ProgramVersion, latestAppBuilderVersion)
	case majorDiff == 0:
		item.Score = 4
		item.Status = models.StatusPass
		item.SetDetails("details.appbuilder_version_current_major", g.AppDef.ProgramVersion, latestAppBuilderVersion)
	case majorDiff == 1:
		item.Score = 3
		item.Status = models.StatusWarning
		item.SetDetails("details.appbuilder_version_one_major_behind", g.AppDef.ProgramVersion, latestAppBuilderVersion)
	case majorDiff == 2:
		item.Score = 1
		item.Status = models.StatusWarning
		item.SetDetails("details.appbuilder_version_two_major_behind", g.AppDef.ProgramVersion, latestAppBuilderVersion)
	default:
		item.Score = 0
		item.Status = models.StatusWarning
		item.SetDetails("details.appbuilder_version_more_than_two_behind", g.AppDef.ProgramVersion, latestAppBuilderVersion)
	}
	return item
}

func (g *Grader) checkShareable() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.shareable_one_to_one_name",
		Description: "line_items.shareable_one_to_one_desc",
		MaxScore:    1.0,
		Status:      models.StatusWarning,
		Details:     "details.shareable_missing",
	}

	if g.hasFeatureValue("share-app-link", "true") || g.hasFeatureValue("text-select-share", "true") {
		item.Score = 1.0
		item.Status = models.StatusPass
		item.SetDetails("details.shareable_found")
	}
	return item
}

func (g *Grader) checkSDCardAvailability() models.LineItem {
	return createIgnoredItem("line_items.sd_card_availability_name", "line_items.sd_card_availability_desc", 1.0)
}

func (g *Grader) checkDirectLinks() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.direct_links_name",
		Description: "line_items.direct_links_desc",
		MaxScore:    1.0,
		Status:      models.StatusWarning,
		Details:     "details.direct_links_missing",
	}

	if g.hasFeatureValue("export-html-pwa", "true") ||
		strings.EqualFold(g.AppDef.Publishing.ScriptureEarth.Notify, "true") ||
		strings.TrimSpace(g.AppDef.PwaManifest.PwaSubDir) != "" ||
		len(informationalLinks(g.furtherInfoText())) > 0 ||
		hasRawDirectLinkEvidence(g.appDefContent()) {
		item.Score = 1.0
		item.Status = models.StatusPass
		item.SetDetails("details.direct_links_found")
	}
	return item
}

func (g *Grader) checkPlayStorePublishing() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.play_store_publishing_name",
		Description: "line_items.play_store_publishing_desc",
		MaxScore:    2.0,
		Status:      models.StatusWarning,
		Details:     "details.play_store_publishing_missing",
	}

	appDef := strings.ToLower(g.appDefContent())
	if hasServicePublishingEvidence(g.AppDef, appDef) {
		item.Score = 2.0
		item.Status = models.StatusPass
		item.SetDetails("details.play_store_publishing_service")
		return item
	}
	if hasLocalPlayStoreEvidence(appDef) {
		item.Score = 1.0
		item.Status = models.StatusPass
		item.SetDetails("details.play_store_publishing_local")
	}
	return item
}

type versionParts struct {
	major int
	minor int
	patch int
}

func parseVersionParts(version string) (versionParts, bool) {
	matches := versionNumberPattern.FindStringSubmatch(version)
	if len(matches) == 0 {
		return versionParts{}, false
	}

	major, err := strconv.Atoi(matches[1])
	if err != nil {
		return versionParts{}, false
	}
	parts := versionParts{major: major}
	if matches[2] != "" {
		minor, err := strconv.Atoi(matches[2])
		if err == nil {
			parts.minor = minor
		}
	}
	if matches[3] != "" {
		patch, err := strconv.Atoi(matches[3])
		if err == nil {
			parts.patch = patch
		}
	}
	return parts, true
}

func compareVersionParts(a, b versionParts) int {
	if a.major != b.major {
		return a.major - b.major
	}
	if a.minor != b.minor {
		return a.minor - b.minor
	}
	return a.patch - b.patch
}

func hasRawDirectLinkEvidence(appDef string) bool {
	lower := strings.ToLower(appDef)
	return strings.Contains(lower, "<pwa-sub-directory>") ||
		strings.Contains(lower, "<pwa-repository") ||
		strings.Contains(lower, "scriptureearth?:  y") ||
		strings.Contains(lower, "osa webpage:   y")
}

func hasServicePublishingEvidence(appDef models.AppDef, rawAppDef string) bool {
	serviceMarkers := []string{
		"mode=\"service\"",
		"scriptoria",
		"app publishing service",
		"kalaam",
		"kalām",
		"sil app publishing",
	}
	if strings.EqualFold(appDef.Publishing.Mode, "service") {
		return true
	}
	for _, marker := range serviceMarkers {
		if strings.Contains(rawAppDef, marker) {
			return true
		}
	}
	return false
}

func hasLocalPlayStoreEvidence(rawAppDef string) bool {
	return strings.Contains(rawAppDef, "<google-play") ||
		strings.Contains(rawAppDef, "playstore") ||
		strings.Contains(rawAppDef, "play store") ||
		strings.Contains(rawAppDef, "google play")
}

var versionNumberPattern = regexp.MustCompile(`^\D*(\d+)(?:\.(\d+))?(?:\.(\d+))?`)
