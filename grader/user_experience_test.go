package grader

import (
	"appbuilder-grader/models"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMenuNavigationScoresVernacularTranslations(t *testing.T) {
	g := &Grader{AppDef: models.AppDef{
		Books: []models.Books{{WritingSystem: models.WritingSystem{Code: "wsg-Telu"}}},
		TranslationMappings: models.TranslationMappings{TranslationMapping: []models.TranslationMapping{{
			Id: "Menu_Contents",
			Translation: []models.TranslationItem{
				{Lang: "en", Value: "Contents"},
				{Lang: "te", Value: "Regional"},
				{Lang: "wsg-tel", Value: "Vernacular"},
			},
		}}},
	}}

	item := g.checkMenuNavigation()
	if item.Status != models.StatusPass || item.Score != 2 {
		t.Fatalf("expected vernacular menu navigation to score 2, got status=%s score=%v", item.Status, item.Score)
	}
}

func TestKeyboardScoresKeymanAboveInputButtons(t *testing.T) {
	g := testProject(t, map[string]string{"MAT": `\id MAT`}, nil)
	writeAppDef(t, g, `<app-definition>
  <features type="main">
    <feature name="input-buttons" value="ka ki ku"/>
  </features>
  <keyboards>
    <keyboard>
      <filename>gondi_gunjala.kmp</filename>
      <trait name="langId" value="wsg-Gong"/>
    </keyboard>
  </keyboards>
</app-definition>`)

	item := g.checkKeyboard()
	if item.Status != models.StatusPass || item.Score != 2 {
		t.Fatalf("expected Keyman keyboard to score 2, got status=%s score=%v", item.Status, item.Score)
	}
}

func TestContentsMenuGraphicsStylesAndTextChanges(t *testing.T) {
	g := testProject(t, map[string]string{"MAT": `\id MAT`}, nil)
	g.AppDef.Features.Feature = []models.Feature{{Name: "book-select", Value: "grid"}}
	g.AppDef.Images = []models.Images{{Type: "splash", Image: []models.ImageItem{{Value: "title.png"}}}}
	writeAppDef(t, g, `<app-definition>
  <styles><style name="p" category="text"></style></styles>
  <fonts><font family="Test"><font-name>Test</font-name></font></fonts>
  <changes type="main">
    <change><name>Keep words together</name><find>a b</find><replace>a\u00A0b</replace></change>
  </changes>
</app-definition>`)

	assertPassScore(t, g.checkContentsMenu(), 1)
	assertPassScore(t, g.checkCustomizedGraphics(), 1)
	assertPassScore(t, g.checkStyleAdjustments(), 1)
	assertPassScore(t, g.checkTextChanges(), 1)
}

func TestAboutBoxInfoAndVernacularRequireEvidence(t *testing.T) {
	g := testProject(t, map[string]string{"MAT": `\id MAT`}, nil)
	g.AppDef.About = models.About{Enabled: "true", Filename: "about.txt"}
	aboutPath := filepath.Join(g.dataDir(), "about", "about.txt")
	if err := os.MkdirAll(filepath.Dir(aboutPath), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(aboutPath, []byte("Copyright 2026\nContact test@example.org\nVersion 1"), 0644); err != nil {
		t.Fatal(err)
	}

	assertPassScore(t, g.checkAboutBoxInfo(), 1)
	item := g.checkAboutBoxVernacular()
	if item.Status == models.StatusPass || item.Score != 0 {
		t.Fatalf("expected English-only about box not to pass vernacular, got status=%s score=%v", item.Status, item.Score)
	}

	vernacular := strings.Repeat("\u0c24\u0c46\u0c32\u0c41\u0c17\u0c41 ", 40)
	if err := os.WriteFile(aboutPath, []byte("Copyright 2026\nContact test@example.org\n"+vernacular), 0644); err != nil {
		t.Fatal(err)
	}
	assertPassScore(t, g.checkAboutBoxVernacular(), 1)
}

func assertPassScore(t *testing.T, item models.LineItem, score float64) {
	t.Helper()
	if item.Status != models.StatusPass || item.Score != score {
		t.Fatalf("expected %s to pass with score %v, got status=%s score=%v", item.Name, score, item.Status, item.Score)
	}
}

func writeAppDef(t *testing.T, g *Grader, content string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(g.TargetDir, "Test.appDef"), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}
