package grader

import (
	"appbuilder-grader/models"
	"testing"
)

func TestAppBuilderVersionScoresAgainstHardcodedLatest(t *testing.T) {
	g := &Grader{AppDef: models.AppDef{ProgramVersion: "14.0"}}
	assertPassScore(t, g.checkAppBuilderVersion(), 5)

	g.AppDef.ProgramVersion = "14.0.1"
	assertPassScore(t, g.checkAppBuilderVersion(), 5)

	g.AppDef.ProgramVersion = "14.0-beta"
	assertPassScore(t, g.checkAppBuilderVersion(), 5)

	g.AppDef.ProgramVersion = "14.0"
	assertPassScore(t, g.checkAppBuilderVersion(), 5)

	g.AppDef.ProgramVersion = "14"
	assertPassScore(t, g.checkAppBuilderVersion(), 5)

	g.AppDef.ProgramVersion = "13.4"
	item := g.checkAppBuilderVersion()
	if item.Status != models.StatusSuggested || item.Score != 3 {
		t.Fatalf("expected version 13.4 to score one major behind, got status=%s score=%v", item.Status, item.Score)
	}

	g.AppDef.ProgramVersion = "12.1"
	item = g.checkAppBuilderVersion()
	if item.Status != models.StatusSuggested || item.Score != 1 {
		t.Fatalf("expected version 12.1 to score two majors behind, got status=%s score=%v", item.Status, item.Score)
	}

	g.AppDef.ProgramVersion = "11.9"
	item = g.checkAppBuilderVersion()
	if item.Status != models.StatusSuggested || item.Score != 0 {
		t.Fatalf("expected version 11.9 to score more than two majors behind, got status=%s score=%v", item.Status, item.Score)
	}
}

func TestShareableDetectsShareFeatures(t *testing.T) {
	g := testProject(t, map[string]string{"MAT": `\id MAT`}, []models.Feature{{Name: "share-app-link", Value: "true"}})
	assertPassScore(t, g.checkShareable(), 1)

	g = testProject(t, map[string]string{"MAT": `\id MAT`}, []models.Feature{{Name: "text-select-share", Value: "true"}})
	assertPassScore(t, g.checkShareable(), 1)
}

func TestDirectLinksDetectsPwaScriptureEarthAndWebsiteLinks(t *testing.T) {
	g := testProject(t, map[string]string{"MAT": `\id MAT`}, []models.Feature{{Name: "export-html-pwa", Value: "true"}})
	assertPassScore(t, g.checkDirectLinks(), 1)

	g = testProject(t, map[string]string{"MAT": `\id MAT`}, nil)
	g.AppDef.Publishing.ScriptureEarth.Notify = "true"
	assertPassScore(t, g.checkDirectLinks(), 1)

	g = testProject(t, map[string]string{"MAT": `\id MAT`}, nil)
	writeAppDef(t, g, `<app-definition>
  <menu-items>
    <menu-item type="website"><link><translation>https://example.org/app</translation></link></menu-item>
  </menu-items>
</app-definition>`)
	assertPassScore(t, g.checkDirectLinks(), 1)
}

func TestPlayStorePublishingScoresServiceAboveLocalEvidence(t *testing.T) {
	g := testProject(t, map[string]string{"MAT": `\id MAT`}, nil)
	g.AppDef.Publishing.Mode = "service"
	assertPassScore(t, g.checkPlayStorePublishing(), 2)

	g = testProject(t, map[string]string{"MAT": `\id MAT`}, nil)
	writeAppDef(t, g, `<app-definition>
  <publishing><google-play verify="false"/></publishing>
</app-definition>`)
	item := g.checkPlayStorePublishing()
	if item.Status != models.StatusPass || item.Score != 1 {
		t.Fatalf("expected local Play Store evidence to score 1, got status=%s score=%v", item.Status, item.Score)
	}
}
