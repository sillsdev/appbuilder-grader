package grader

import (
	"appbuilder-grader/models"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestClickableReferencesRequireMoreThanSixtyPercentOfBookFiles(t *testing.T) {
	g := testProject(t, map[string]string{
		"MAT": `\id MAT` + "\n" + `\v 1 linked \x - \xo 1.1 \xt MRK 1.1\x*`,
		"MRK": `\id MRK` + "\n" + `\v 1 linked \x - \xo 1.1 \xt MAT 1.1\x*`,
		"LUK": `\id LUK` + "\n" + `\v 1 linked \x - \xo 1.1 \ref John|JHN 1:1\ref*\x*`,
		"JHN": `\id JHN` + "\n" + `\v 1 linked \ref Mark|MRK 1:1\ref*`,
		"ACT": `\id ACT` + "\n" + `\v 1 plain text`,
	}, nil)

	item := g.checkClickableReferences()
	if item.Status != models.StatusPass || item.Score != 1 {
		t.Fatalf("expected clickable references to pass above 60%%, got status=%s score=%v", item.Status, item.Score)
	}

	g = testProject(t, map[string]string{
		"MAT": `\id MAT` + "\n" + `\v 1 linked \x - \xo 1.1 \xt MRK 1.1\x*`,
		"MRK": `\id MRK` + "\n" + `\v 1 linked \x - \xo 1.1 \xt MAT 1.1\x*`,
		"LUK": `\id LUK` + "\n" + `\v 1 linked \x - \xo 1.1 \xt JHN 1.1\x*`,
		"JHN": `\id JHN` + "\n" + `\v 1 plain text`,
		"ACT": `\id ACT` + "\n" + `\v 1 plain text`,
	}, nil)

	item = g.checkClickableReferences()
	if item.Status == models.StatusPass || item.Score != 0 {
		t.Fatalf("expected clickable references not to pass at exactly 60%%, got status=%s score=%v", item.Status, item.Score)
	}
}

func TestRedLetterRequiresFeatureAndWordsOfJesusMarkersInNTFiles(t *testing.T) {
	features := []models.Feature{{Name: "show-red-letters", Value: "false"}}
	g := testProject(t, map[string]string{
		"MAT": `\id MAT` + "\n" + `\v 1 \wj Words of Jesus\wj*`,
	}, features)

	item := g.checkRedLetterText()
	if item.Status != models.StatusPass || item.Score != 1 {
		t.Fatalf("expected red letter to pass with feature and NT wj marker, got status=%s score=%v", item.Status, item.Score)
	}

	g = testProject(t, map[string]string{
		"MAT": `\id MAT` + "\n" + `\v 1 \wj Words of Jesus\wj*`,
	}, nil)
	item = g.checkRedLetterText()
	if item.Status == models.StatusPass || item.Score != 0 {
		t.Fatalf("expected red letter not to pass without feature, got status=%s score=%v", item.Status, item.Score)
	}

	g = testProject(t, map[string]string{
		"MAT": `\id MAT` + "\n" + `\v 1 plain text`,
	}, features)
	item = g.checkRedLetterText()
	if item.Status == models.StatusPass || item.Score != 0 {
		t.Fatalf("expected red letter not to pass without NT wj marker, got status=%s score=%v", item.Status, item.Score)
	}
}

func TestMultilingualScriptsScoresBackTranslationHighest(t *testing.T) {
	g := &Grader{AppDef: models.AppDef{Books: []models.Books{
		{Id: "WSG", BookCollectionName: "Gondi Telugu script", WritingSystem: models.WritingSystem{Code: "wsg-Telu"}},
		{Id: "WSGdev", BookCollectionName: "Gondi Devanagari script", WritingSystem: models.WritingSystem{Code: "wsg-Deva"}},
		{Id: "TEL", BookCollectionName: "Telugu regional language", WritingSystem: models.WritingSystem{Code: "tel"}},
		{Id: "WSGBTeng", BookCollectionName: "English with Gondi Audio", WritingSystem: models.WritingSystem{Code: "eng"}},
	}}}

	item := g.checkMultilingualScripts()
	if item.Status != models.StatusPass || item.Score != 3 {
		t.Fatalf("expected multilingual scripts to score back translation highest, got status=%s score=%v", item.Status, item.Score)
	}
}

func testProject(t *testing.T, books map[string]string, features []models.Feature) *Grader {
	t.Helper()

	targetDir := t.TempDir()
	dataDir := filepath.Join(targetDir, "Test_data", "books", "C01")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(targetDir, "Test.appDef"), []byte("<app-definition/>"), 0644); err != nil {
		t.Fatal(err)
	}

	bookItems := make([]models.Book, 0, len(books))
	for id, content := range books {
		filename := fmt.Sprintf("%s.txt", id)
		if err := os.WriteFile(filepath.Join(dataDir, filename), []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
		bookItems = append(bookItems, models.Book{Id: id, Filename: filename})
	}

	return &Grader{
		TargetDir: targetDir,
		AppDef: models.AppDef{
			Features: models.Features{Feature: features},
			Books: []models.Books{{
				Id:                   "C01",
				BookCollectionName:   "Test collection",
				WritingSystem:        models.WritingSystem{Code: "tst-Latn"},
				Book:                 bookItems,
				BookCollectionAbbrev: "TST",
			}},
		},
	}
}
