package grader

import (
	"appbuilder-grader/models"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPlayStoreDescriptionRequiresMeaningfulListingText(t *testing.T) {
	g := testProject(t, map[string]string{"MAT": `\id MAT`}, nil)
	writePlayListing(t, g, "en-GB", "Test App", "Too short", strings.Repeat("word ", 30), "")

	item := g.checkAppDescription()
	if item.Status == models.StatusPass || item.Score != 0 {
		t.Fatalf("expected short listing not to pass, got status=%s score=%v", item.Status, item.Score)
	}

	writePlayListing(t, g, "en-GB", "Test App", "A helpful short app description", strings.Repeat("word ", 30), "")
	assertPassScore(t, g.checkAppDescription(), 1)
}

func TestLocalizedDescriptionsMatchRegionalAndVernacularLocales(t *testing.T) {
	g := testProject(t, map[string]string{"MAT": `\id MAT`}, nil)
	writePlayListing(t, g, "hi-IN", "Test App", "A helpful short app description", strings.Repeat("word ", 30), "")
	assertPassScore(t, g.checkLocalizedDescriptionLWC(), 1)

	item := g.checkLocalizedDescriptionVernacular()
	if item.Status == models.StatusPass || item.Score != 0 {
		t.Fatalf("expected regional listing not to pass vernacular, got status=%s score=%v", item.Status, item.Score)
	}

	writePlayListing(t, g, "tst", "Test App", "A helpful short app description", strings.Repeat("word ", 30), "")
	assertPassScore(t, g.checkLocalizedDescriptionVernacular(), 1)
}

func TestPlayStoreScreenshotsAndDemoVideo(t *testing.T) {
	g := testProject(t, map[string]string{"MAT": `\id MAT`}, nil)
	writePlayListing(t, g, "en-GB", "Test App", "A helpful short app description", strings.Repeat("word ", 30), "https://example.org/demo.mp4")
	writePlayListingImage(t, g, "en-GB", "phoneScreenshots", "s1.png")
	writePlayListingImage(t, g, "en-GB", "sevenInchScreenshots", "s1.jpg")

	assertPassScore(t, g.checkPhoneScreenshots(), 1)
	assertPassScore(t, g.checkTabletScreenshots(), 1)
	assertPassScore(t, g.checkDemoVideo(), 1)
}

func TestDemoVideoRequiresURL(t *testing.T) {
	g := testProject(t, map[string]string{"MAT": `\id MAT`}, nil)
	writePlayListing(t, g, "en-GB", "Test App", "A helpful short app description", strings.Repeat("word ", 30), "coming soon")

	item := g.checkDemoVideo()
	if item.Status == models.StatusPass || item.Score != 0 {
		t.Fatalf("expected non-url demo video not to pass, got status=%s score=%v", item.Status, item.Score)
	}
}

func writePlayListing(t *testing.T, g *Grader, locale, title, shortDescription, fullDescription, video string) {
	t.Helper()
	root := filepath.Join(g.dataDir(), "publish", "play-listing", locale)
	if err := os.MkdirAll(root, 0755); err != nil {
		t.Fatal(err)
	}
	files := map[string]string{
		"title.txt":             title,
		"short_description.txt": shortDescription,
		"full_description.txt":  fullDescription,
		"video.txt":             video,
	}
	for filename, content := range files {
		if err := os.WriteFile(filepath.Join(root, filename), []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}
}

func writePlayListingImage(t *testing.T, g *Grader, locale, folder, filename string) {
	t.Helper()
	path := filepath.Join(g.dataDir(), "publish", "play-listing", locale, "images", folder, filename)
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte("image"), 0644); err != nil {
		t.Fatal(err)
	}
}
