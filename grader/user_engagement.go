package grader

import (
	"appbuilder-grader/models"
	"regexp"
	"strings"
)

func (g *Grader) checkUserEngagement() models.Category {
	cat := models.Category{
		Name:        "categories.user_engagement_name",
		Description: "categories.user_engagement_desc",
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
		Name:        "line_items.analytics_name",
		Description: "line_items.analytics_desc",
		MaxScore:    1.0,
		Status:      models.StatusWarning,
		Details:     "details.analytics_missing",
	}

	appDef := strings.ToLower(g.appDefContent())
	hasAnalyticsProvider := strings.Contains(appDef, `<analytics enabled="true"`) &&
		strings.Contains(appDef, "<analytics-provider")
	hasFirebaseAnalytics := strings.Contains(appDef, `feature name="firebase-analytics" value="true"`)
	if hasAnalyticsProvider || hasFirebaseAnalytics {
		item.Score = 1.0
		item.Status = models.StatusPass
		item.SetDetails("details.analytics_found")
	}
	return item
}

func (g *Grader) checkTextFeedback() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.text_feedback_name",
		Description: "line_items.text_feedback_desc",
		MaxScore:    1.0,
		Status:      models.StatusWarning,
		Details:     "details.text_feedback_missing",
	}

	projectText := g.projectEngagementText()
	if g.hasFeatureValue("editor", "true") && hasContactEvidence(projectText) {
		item.Score = 1.0
		item.Status = models.StatusPass
		item.SetDetails("details.text_feedback_found")
	}
	return item
}

func (g *Grader) checkFurtherInfoLink() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.further_info_link_name",
		Description: "line_items.further_info_link_desc",
		MaxScore:    1.0,
		Status:      models.StatusWarning,
		Details:     "details.further_info_link_missing",
	}

	links := informationalLinks(g.furtherInfoText())
	if len(links) > 0 {
		item.Score = 1.0
		item.Status = models.StatusPass
		item.SetDetails("details.further_info_link_found", len(links))
	}
	return item
}

func (g *Grader) checkContactDetails() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.contact_details_name",
		Description: "line_items.contact_details_desc",
		MaxScore:    1.0,
		Status:      models.StatusWarning,
		Details:     "details.contact_details_missing",
	}

	if hasContactEvidence(g.projectEngagementText()) {
		item.Score = 1.0
		item.Status = models.StatusPass
		item.SetDetails("details.contact_details_found")
	}
	return item
}

func (g *Grader) checkDeepLinking() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.deep_linking_name",
		Description: "line_items.deep_linking_desc",
		MaxScore:    1.0,
		Status:      models.StatusWarning,
		Details:     "details.deep_linking_missing",
	}

	if strings.EqualFold(g.AppDef.DeepLinking.Enabled, "true") ||
		strings.TrimSpace(g.AppDef.DeepLinking.Uri.Value) != "" {
		item.Score = 1.0
		item.Status = models.StatusPass
		item.SetDetails("details.deep_linking_found")
	}
	return item
}

func (g *Grader) projectEngagementText() string {
	appDef := g.appDefContent()
	blocks := []string{g.aboutContent()}
	for _, tag := range []string{"project-description", "menu-items", "footer"} {
		blocks = append(blocks, xmlBlocks(appDef, tag)...)
	}
	return strings.Join(blocks, "\n")
}

func (g *Grader) furtherInfoText() string {
	appDef := g.appDefContent()
	blocks := []string{g.aboutContent()}
	for _, tag := range []string{"menu-items", "footer"} {
		blocks = append(blocks, xmlBlocks(appDef, tag)...)
	}
	return strings.Join(blocks, "\n")
}

func hasContactEvidence(content string) bool {
	lower := strings.ToLower(content)
	return strings.Contains(lower, "mailto:") ||
		emailPattern.MatchString(content) ||
		phonePattern.MatchString(content) ||
		len(informationalLinks(content)) > 0
}

func informationalLinks(content string) []string {
	matches := urlPattern.FindAllString(content, -1)
	links := make([]string, 0, len(matches))
	seen := make(map[string]bool)
	for _, match := range matches {
		link := strings.TrimRight(match, ".,);]")
		lower := strings.ToLower(link)
		if strings.HasPrefix(lower, "mailto:") || isMediaOrAssetLink(lower) {
			continue
		}
		if !seen[link] {
			seen[link] = true
			links = append(links, link)
		}
	}
	return links
}

func isMediaOrAssetLink(link string) bool {
	mediaHostsOrPaths := []string{
		"media.ipsapps.org",
		"api.arclight.org",
		"jesusfilm.org",
		".mp3",
		".mp4",
		".m4a",
		".wav",
		".jpg",
		".jpeg",
		".png",
	}
	for _, marker := range mediaHostsOrPaths {
		if strings.Contains(link, marker) {
			return true
		}
	}
	return false
}

var (
	emailPattern = regexp.MustCompile(`[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}`)
	phonePattern = regexp.MustCompile(`(?m)(?:\+?\d[\d .()\-]{7,}\d)`)
	urlPattern   = regexp.MustCompile(`(?i)\b(?:https?://|www\.)[^\s<>"']+`)
)

func xmlBlocks(content, tag string) []string {
	pattern := regexp.MustCompile(`(?is)<` + regexp.QuoteMeta(tag) + `\b[^>]*>.*?</` + regexp.QuoteMeta(tag) + `>`)
	return pattern.FindAllString(content, -1)
}
