package grader

import (
	"appbuilder-grader/models"
	"os"
	"path/filepath"
	"testing"
)

func TestAnalyticsDetectsProviderOrFirebase(t *testing.T) {
	g := testProject(t, map[string]string{"MAT": `\id MAT`}, nil)
	writeAppDef(t, g, `<app-definition>
  <analytics enabled="true">
    <analytics-provider type="firebase"/>
  </analytics>
</app-definition>`)
	assertPassScore(t, g.checkAnalytics(), 1)

	g = testProject(t, map[string]string{"MAT": `\id MAT`}, nil)
	writeAppDef(t, g, `<app-definition>
  <features type="main">
    <feature name="firebase-analytics" value="true"/>
  </features>
</app-definition>`)
	assertPassScore(t, g.checkAnalytics(), 1)
}

func TestTextFeedbackRequiresEditorAndContactEvidence(t *testing.T) {
	g := testProject(t, map[string]string{"MAT": `\id MAT`}, nil)
	writeAppDef(t, g, `<app-definition>
  <project-description>Feedback: editor@example.org</project-description>
</app-definition>`)

	item := g.checkTextFeedback()
	if item.Status == models.StatusPass || item.Score != 0 {
		t.Fatalf("expected contact without editor not to pass text feedback, got status=%s score=%v", item.Status, item.Score)
	}

	g.AppDef.Features.Feature = []models.Feature{{Name: "editor", Value: "true"}}
	assertPassScore(t, g.checkTextFeedback(), 1)
}

func TestFurtherInfoLinkUsesUserFacingNonMediaLinks(t *testing.T) {
	g := testProject(t, map[string]string{"MAT": `\id MAT`}, nil)
	writeAppDef(t, g, `<app-definition>
  <project-url>https://example.org/internal-project-page</project-url>
  <menu-items>
    <menu-item><url>https://media.ipsapps.org/audio/MAT.mp3</url></menu-item>
  </menu-items>
</app-definition>`)

	item := g.checkFurtherInfoLink()
	if item.Status == models.StatusPass || item.Score != 0 {
		t.Fatalf("expected media/internal links not to pass further info, got status=%s score=%v", item.Status, item.Score)
	}

	writeAppDef(t, g, `<app-definition>
  <menu-items>
    <menu-item><url>https://example.org/gondi-scripture</url></menu-item>
  </menu-items>
</app-definition>`)
	assertPassScore(t, g.checkFurtherInfoLink(), 1)
}

func TestContactDetailsFromAboutOrProjectText(t *testing.T) {
	g := testProject(t, map[string]string{"MAT": `\id MAT`}, nil)
	g.AppDef.About = models.About{Enabled: "true", Filename: "about.txt"}
	aboutPath := filepath.Join(g.dataDir(), "about", "about.txt")
	if err := os.MkdirAll(filepath.Dir(aboutPath), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(aboutPath, []byte("Contact: mailto:team@example.org"), 0644); err != nil {
		t.Fatal(err)
	}
	assertPassScore(t, g.checkContactDetails(), 1)

	g = testProject(t, map[string]string{"MAT": `\id MAT`}, nil)
	writeAppDef(t, g, `<app-definition>
  <project-description>Call +1 555 123 4567 for follow-up.</project-description>
</app-definition>`)
	assertPassScore(t, g.checkContactDetails(), 1)
}

func TestDeepLinkingRequiresEnabledOrUri(t *testing.T) {
	g := testProject(t, map[string]string{"MAT": `\id MAT`}, nil)
	item := g.checkDeepLinking()
	if item.Status == models.StatusPass || item.Score != 0 {
		t.Fatalf("expected blank deep linking not to pass, got status=%s score=%v", item.Status, item.Score)
	}

	g.AppDef.DeepLinking = models.DeepLinking{Enabled: "true"}
	assertPassScore(t, g.checkDeepLinking(), 1)

	g.AppDef.DeepLinking = models.DeepLinking{Uri: models.Uri{Value: "example://scripture"}}
	assertPassScore(t, g.checkDeepLinking(), 1)
}
