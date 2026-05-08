package grader

import (
	"appbuilder-grader/models"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestAudioScoresFullNewTestament(t *testing.T) {
	books := make(map[string]string)
	for _, bookID := range AllNTBooks {
		books[bookID] = fmt.Sprintf(`\id %s`, bookID)
	}
	g := testProject(t, books, nil)
	addAudioToBooks(g, AllNTBooks...)

	item := g.checkAudio()
	if item.Status != models.StatusPass || item.Score != 2 {
		t.Fatalf("expected full NT audio to score 2, got status=%s score=%v", item.Status, item.Score)
	}
}

func TestAudioScoresFullStandardBibleAndGlossary(t *testing.T) {
	books := make(map[string]string)
	for _, bookID := range standardBibleBooks() {
		books[bookID] = fmt.Sprintf(`\id %s`, bookID)
	}
	books["GLO"] = `\id GLO`
	g := testProject(t, books, nil)
	addAudioToBooks(g, append(standardBibleBooks(), "GLO")...)

	item := g.checkAudio()
	if item.Status != models.StatusPass || item.Score != 4 {
		t.Fatalf("expected glossary audio to score 4, got status=%s score=%v", item.Status, item.Score)
	}
}

func TestAudioStyleDetectsDramatizedFcbhDamID(t *testing.T) {
	g := testProject(t, map[string]string{
		"MAT": `\id MAT`,
	}, nil)
	addAudioToBooks(g, "MAT")
	g.AppDef.AudioSources.AudioSource = []models.AudioSource{{Type: "fcbh", DamID: "WSGWYIN1DA"}}

	item := g.checkStyleOfAudio()
	if item.Status != models.StatusPass || item.Score != 2 {
		t.Fatalf("expected dramatized audio style to score 2, got status=%s score=%v", item.Status, item.Score)
	}
}

func TestTimingRequiresMoreThanSixtyPercentCoverage(t *testing.T) {
	g := testProject(t, map[string]string{
		"MAT": `\id MAT`,
	}, nil)
	g.AppDef.Books[0].Book[0].Audio = []models.Audio{
		{Chapter: "1", TimingFilename: "MAT-01-timing.txt"},
		{Chapter: "2", TimingFilename: "MAT-02-timing.txt"},
		{Chapter: "3", TimingFilename: "MAT-03-timing.txt"},
		{Chapter: "4", TimingFilename: "MAT-04-timing.txt"},
		{Chapter: "5", TimingFilename: "MAT-05-timing.txt"},
	}
	writeTimingFile(t, g, "MAT-01-timing.txt")
	writeTimingFile(t, g, "MAT-02-timing.txt")
	writeTimingFile(t, g, "MAT-03-timing.txt")

	item := g.checkSynchronizedHighlighting()
	if item.Status == models.StatusPass || item.Score != 0 {
		t.Fatalf("expected exactly 60%% timing coverage not to pass, got status=%s score=%v", item.Status, item.Score)
	}

	writeTimingFile(t, g, "MAT-04-timing.txt")
	item = g.checkSynchronizedHighlighting()
	if item.Status != models.StatusPass || item.Score != 1 {
		t.Fatalf("expected timing coverage above 60%% to pass, got status=%s score=%v", item.Status, item.Score)
	}
}

func TestVideoScoresLukeAndOtherGospelJesusFilmWithoutCaptions(t *testing.T) {
	g := testProject(t, map[string]string{
		"LUK": `\id LUK`,
		"JHN": `\id JHN`,
	}, nil)
	g.AppDef.Videos.Video = []models.Video{{
		Id:        "jf6101",
		OnlineUrl: "https://api.arclight.org/videoPlayerUrl?refId=1_jf6101",
		Placement: models.Placement{Ref: "C01|LUK.1.1"},
	}}

	item := g.checkVideo()
	if item.Status != models.StatusPass || item.Score != 1 {
		t.Fatalf("expected Luke Jesus Film video to score 1, got status=%s score=%v", item.Status, item.Score)
	}

	g.AppDef.Videos.Video = append(g.AppDef.Videos.Video, models.Video{
		Id:        "jf6201",
		OnlineUrl: "https://api.arclight.org/videoPlayerUrl?refId=1_jf6201",
		Placement: models.Placement{Ref: "C01|JHN.1.1"},
	})
	item = g.checkVideo()
	if item.Status != models.StatusPass || item.Score != 3 {
		t.Fatalf("expected other-gospel Jesus Film video to score 3, got status=%s score=%v", item.Status, item.Score)
	}
}

func TestVideoScoresOtherLinksAboveJesusFilm(t *testing.T) {
	g := testProject(t, map[string]string{
		"LUK": `\id LUK`,
		"JHN": `\id JHN`,
		"ACT": `\id ACT`,
	}, nil)
	g.AppDef.Videos.Video = []models.Video{
		{Id: "jf6101", OnlineUrl: "https://api.arclight.org/videoPlayerUrl?refId=1_jf6101", Placement: models.Placement{Ref: "C01|LUK.1.1"}},
		{Id: "jf6201", OnlineUrl: "https://api.arclight.org/videoPlayerUrl?refId=1_jf6201", Placement: models.Placement{Ref: "C01|JHN.1.1"}},
		{Id: "acts-overview", OnlineUrl: "https://example.org/acts-overview.mp4", Placement: models.Placement{Ref: "C01|ACT.1.1"}},
	}

	item := g.checkVideo()
	if item.Status != models.StatusPass || item.Score != 4 {
		t.Fatalf("expected other video links above Jesus Film to score 4, got status=%s score=%v", item.Status, item.Score)
	}
}

func addAudioToBooks(g *Grader, bookIDs ...string) {
	want := make(map[string]bool)
	for _, bookID := range bookIDs {
		want[bookID] = true
	}
	for collectionIndex := range g.AppDef.Books {
		for bookIndex := range g.AppDef.Books[collectionIndex].Book {
			bookID := g.AppDef.Books[collectionIndex].Book[bookIndex].Id
			if want[bookID] {
				g.AppDef.Books[collectionIndex].Book[bookIndex].Audio = []models.Audio{{Chapter: "1"}}
			}
		}
	}
}

func writeTimingFile(t *testing.T, g *Grader, filename string) {
	t.Helper()

	path := filepath.Join(g.dataDir(), "timings", filename)
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte("0.000\t1.000\t1\n"), 0644); err != nil {
		t.Fatal(err)
	}
}
