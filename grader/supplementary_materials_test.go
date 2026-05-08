package grader

import (
	"appbuilder-grader/models"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"testing"
)

func TestGlossaryScoresLinkedGlossary(t *testing.T) {
	g := testProject(t, map[string]string{
		"MAT": `\id MAT` + "\n" + `\v 1 \w grace\w* in the text`,
		"GLO": `\id GLO` + "\n" + `\k grace\k* explanation`,
	}, nil)

	item := g.checkGlossary()
	if item.Status != models.StatusPass || item.Score != 2 {
		t.Fatalf("expected linked glossary to score 2, got status=%s score=%v", item.Status, item.Score)
	}
}

func TestGlossaryScoresLinkedGlossaryWithIllustrations(t *testing.T) {
	g := testProject(t, map[string]string{
		"MAT": `\id MAT` + "\n" + `\v 1 \w grace\w* in the text`,
		"GLO": `\id GLO` + "\n" + `\fig Grace image|grace.jpg\fig*`,
	}, nil)

	item := g.checkGlossary()
	if item.Status != models.StatusPass || item.Score != 3 {
		t.Fatalf("expected illustrated linked glossary to score 3, got status=%s score=%v", item.Status, item.Score)
	}
}

func TestIllustrationsScoresColorImages(t *testing.T) {
	g := testProject(t, map[string]string{
		"MAT": `\id MAT` + "\n" + `\v 1 text`,
	}, nil)
	g.AppDef.Images = []models.Images{{
		Type:  "illustration",
		Image: []models.ImageItem{{Value: "color.png"}},
	}}

	imagePath := filepath.Join(g.dataDir(), "images", "illustrations", "color.png")
	if err := os.MkdirAll(filepath.Dir(imagePath), 0755); err != nil {
		t.Fatal(err)
	}
	writeTestPNG(t, imagePath, color.RGBA{R: 220, G: 60, B: 40, A: 255})

	item := g.checkIllustrations()
	if item.Status != models.StatusPass || item.Score != 2 {
		t.Fatalf("expected color illustrations to score 2, got status=%s score=%v", item.Status, item.Score)
	}
}

func TestTopicalIndexFound(t *testing.T) {
	g := testProject(t, map[string]string{
		"TDX": `\id TDX` + "\n" + `\toc1 Topical Index`,
	}, nil)

	item := g.checkTopicalIndex()
	if item.Status != models.StatusPass || item.Score != 1 {
		t.Fatalf("expected topical index to pass, got status=%s score=%v", item.Status, item.Score)
	}
}

func TestReadingPlanFromDataDir(t *testing.T) {
	g := testProject(t, map[string]string{
		"MAT": `\id MAT` + "\n" + `\v 1 text`,
	}, nil)

	planPath := filepath.Join(g.dataDir(), "plans", "plan.txt")
	if err := os.MkdirAll(filepath.Dir(planPath), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(planPath, []byte(`\day 1`+"\n"+`\ref MAT 1:1`), 0644); err != nil {
		t.Fatal(err)
	}

	item := g.checkReadingPlan()
	if item.Status != models.StatusPass || item.Score != 1 {
		t.Fatalf("expected reading plan to pass, got status=%s score=%v", item.Status, item.Score)
	}
}

func TestStudyBibleMaterialIgnoresOrdinaryFootnotes(t *testing.T) {
	g := testProject(t, map[string]string{
		"MAT": `\id MAT` + "\n" + `\v 1 text \f + \ft ordinary footnote\f*`,
	}, nil)

	item := g.checkStudyBibleMaterial()
	if item.Status == models.StatusPass || item.Score != 0 {
		t.Fatalf("expected ordinary footnotes not to count as study material, got status=%s score=%v", item.Status, item.Score)
	}

	g = testProject(t, map[string]string{
		"MAT": `\id MAT` + "\n" + `\v 1 text \ef - \ft extended note\ef*`,
	}, nil)
	item = g.checkStudyBibleMaterial()
	if item.Status != models.StatusPass || item.Score != 1 {
		t.Fatalf("expected extended footnote to count as study material, got status=%s score=%v", item.Status, item.Score)
	}
}

func writeTestPNG(t *testing.T, path string, c color.Color) {
	t.Helper()

	img := image.NewRGBA(image.Rect(0, 0, 8, 8))
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			img.Set(x, y, c)
		}
	}
	file, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	if err := png.Encode(file, img); err != nil {
		t.Fatal(err)
	}
}
