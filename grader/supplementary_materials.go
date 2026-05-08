package grader

import (
	"appbuilder-grader/models"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"
)

func (g *Grader) checkSupplementaryMaterials() models.Category {
	cat := models.Category{
		Name:        "categories.supplementary_materials_name",
		Description: "categories.supplementary_materials_desc",
		Weight:      1.0,
	}

	// Glossary, Illustrations, Topical Index, Reading Plan, Study Bible Material
	cat.LineItems = append(cat.LineItems, g.checkGlossary())
	cat.LineItems = append(cat.LineItems, g.checkIllustrations())
	cat.LineItems = append(cat.LineItems, g.checkTopicalIndex())
	cat.LineItems = append(cat.LineItems, g.checkReadingPlan())
	cat.LineItems = append(cat.LineItems, g.checkStudyBibleMaterial())

	return cat
}

func (g *Grader) checkGlossary() models.LineItem {
	// 0=No glossary
	// 1=Standalone glossary
	// 2=Linked glossary
	// 3=Linked glossary with illustrations
	item := models.LineItem{
		Name:        "line_items.glossary_name",
		Description: "line_items.glossary_desc",
		MaxScore:    3.0,
		Status:      models.StatusWarning,
		Details:     "details.glossary_missing",
	}

	glossaryFiles := g.bookFilesByID("GLO")
	if len(glossaryFiles) == 0 {
		return item
	}

	item.Score = 1.0
	item.SetDetails("details.glossary_standalone", len(glossaryFiles))

	linkedFiles := 0
	for _, bookFile := range g.bookFiles() {
		if bookFile.BookID == "GLO" {
			continue
		}
		if strings.Contains(readTextFile(bookFile.Path), `\w `) {
			linkedFiles++
		}
	}
	if linkedFiles > 0 {
		item.Score = 2.0
		item.Status = models.StatusPass
		item.SetDetails("details.glossary_linked", len(glossaryFiles), linkedFiles)
	}

	glossaryIllustrationFiles := 0
	for _, glossaryFile := range glossaryFiles {
		if strings.Contains(readTextFile(glossaryFile.Path), `\fig`) {
			glossaryIllustrationFiles++
		}
	}
	if glossaryIllustrationFiles > 0 {
		item.Score = 3.0
		item.Status = models.StatusPass
		item.SetDetails("details.glossary_linked_with_illustrations", len(glossaryFiles), linkedFiles, glossaryIllustrationFiles)
	}

	return item
}

func (g *Grader) checkIllustrations() models.LineItem {
	// 0=No illustrations
	// 1=BW line art
	// 2=Colour illustrations
	item := models.LineItem{
		Name:        "line_items.illustrations_name",
		Description: "line_items.illustrations_desc",
		MaxScore:    2.0,
		Status:      models.StatusWarning,
		Details:     "details.illustrations_missing",
	}

	illustrationPaths := g.illustrationPaths()
	figFiles := 0
	for _, bookFile := range g.bookFiles() {
		if strings.Contains(readTextFile(bookFile.Path), `\fig`) {
			figFiles++
		}
	}
	if len(illustrationPaths) == 0 && figFiles == 0 {
		return item
	}

	item.Score = 1.0
	item.Status = models.StatusPass
	item.SetDetails("details.illustrations_found", len(illustrationPaths), figFiles)

	colorImages := 0
	for _, path := range illustrationPaths {
		if imageHasColor(path) {
			colorImages++
		}
	}
	if colorImages > 0 {
		item.Score = 2.0
		item.SetDetails("details.illustrations_color_found", colorImages, len(illustrationPaths), figFiles)
	}
	return item
}

func (g *Grader) checkTopicalIndex() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.topical_index_name",
		Description: "line_items.topical_index_desc",
		MaxScore:    1.0,
		Status:      models.StatusWarning,
		Details:     "details.topical_index_missing",
	}
	files := g.bookFilesByID("TDX")
	if len(files) > 0 {
		item.Score = 1.0
		item.Status = models.StatusPass
		item.SetDetails("details.topical_index_found", len(files))
	}
	return item
}

func (g *Grader) checkReadingPlan() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.reading_plan_name",
		Description: "line_items.reading_plan_desc",
		MaxScore:    1.0,
		Status:      models.StatusWarning,
		Details:     "details.reading_plan_missing",
	}

	validPlanPaths := make(map[string]bool)
	for _, plan := range g.AppDef.Plans.Plan {
		if strings.TrimSpace(plan.Filename) == "" {
			continue
		}
		path := g.resolveDataFile("plans", plan.Filename)
		if path != "" && isValidReadingPlan(readTextFile(path)) {
			validPlanPaths[path] = true
		}
	}
	for _, path := range g.filesUnderDataDir("plans") {
		if strings.EqualFold(filepath.Ext(path), ".txt") && isValidReadingPlan(readTextFile(path)) {
			validPlanPaths[path] = true
		}
	}

	if len(validPlanPaths) > 0 {
		item.Score = 1.0
		item.Status = models.StatusPass
		item.SetDetails("details.reading_plan_found", len(validPlanPaths))
	}
	return item
}

func (g *Grader) checkStudyBibleMaterial() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.study_bible_material_name",
		Description: "line_items.study_bible_material_desc",
		MaxScore:    1.0,
		Status:      models.StatusWarning,
		Details:     "details.study_bible_material_missing",
	}

	materialFiles := 0
	for _, bookFile := range g.bookFiles() {
		if hasStudyBibleMaterial(readTextFile(bookFile.Path)) {
			materialFiles++
		}
	}
	if materialFiles > 0 {
		item.Score = 1.0
		item.Status = models.StatusPass
		item.SetDetails("details.study_bible_material_found", materialFiles)
	}
	return item
}

func (g *Grader) illustrationPaths() []string {
	paths := make([]string, 0)
	for _, images := range g.AppDef.Images {
		if images.Type != "illustration" {
			continue
		}
		for _, imageItem := range images.Image {
			if resolved := g.resolveDataFile("images", "illustrations", imageItem.Value); resolved != "" {
				paths = append(paths, resolved)
			}
		}
	}
	if len(paths) == 0 {
		for _, path := range g.filesUnderDataDir("images/illustrations") {
			paths = append(paths, path)
		}
	}
	return paths
}

func isValidReadingPlan(content string) bool {
	return strings.Contains(content, `\day `) && strings.Contains(content, `\ref `)
}

func hasStudyBibleMaterial(content string) bool {
	return strings.Contains(content, `\ef `) ||
		strings.Contains(content, `\esb`) ||
		strings.Contains(content, `\jmp`)
}

func imageHasColor(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		return false
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return false
	}
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	if width == 0 || height == 0 {
		return false
	}

	stepX := max(1, width/24)
	stepY := max(1, height/24)
	colorfulPixels := 0
	sampledPixels := 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y += stepY {
		for x := bounds.Min.X; x < bounds.Max.X; x += stepX {
			r, g, b, _ := img.At(x, y).RGBA()
			r8 := int(r >> 8)
			g8 := int(g >> 8)
			b8 := int(b >> 8)
			if maxChannel(r8, g8, b8)-minChannel(r8, g8, b8) > 20 {
				colorfulPixels++
			}
			sampledPixels++
		}
	}
	return sampledPixels > 0 && float64(colorfulPixels)/float64(sampledPixels) > 0.03
}

func maxChannel(a, b, c int) int {
	return max(a, max(b, c))
}

func minChannel(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < c {
		return b
	}
	return c
}
