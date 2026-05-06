package grader

import (
	"appbuilder-grader/models"
	"os"
	"path/filepath"
)

// checkFileStructure verifies essential file structure
func (g *Grader) checkFileStructure() models.Category {
	cat := models.Category{
		Name:        "File Structure",
		Description: "Checks if the standard build output files are present.",
		Weight:      2.0,
		LineItems:   make([]models.LineItem, 0),
	}

	score := 0.0

	// Line Item 1
	idxItem := models.LineItem{
		Name:        "Index HTML",
		Description: "Check for index.html presence",
		MaxScore:    5.0,
	}
	if _, err := os.Stat(filepath.Join(g.TargetDir, "index.html")); err == nil {
		idxItem.Score = 5.0
		idxItem.Status = "pass"
		idxItem.Details = "Found index.html"
	} else {
		idxItem.Status = "error"
		idxItem.Details = "Missing index.html"
	}
	score += idxItem.Score
	cat.LineItems = append(cat.LineItems, idxItem)

	// Line Item 2
	jsItem := models.LineItem{
		Name:        "Main JS",
		Description: "Check for main.js presence",
		MaxScore:    5.0,
	}
	if _, err := os.Stat(filepath.Join(g.TargetDir, "main.js")); err == nil {
		jsItem.Score = 5.0
		jsItem.Status = "pass"
		jsItem.Details = "Found main.js"
	} else {
		jsItem.Status = "error"
		jsItem.Details = "Missing main.js"
	}
	score += jsItem.Score
	cat.LineItems = append(cat.LineItems, jsItem)

	// Line Item 3 (Ignored Item Example)
	ignoreItem := models.LineItem{
		Name:        "Optional Sourcemaps",
		Description: "Check for optional .map files",
		MaxScore:    5.0,
		Status:      "ignored",
		Details:     "Sourcemaps check disabled in this environment.",
	}
	cat.LineItems = append(cat.LineItems, ignoreItem)

	cat.Score = score
	if score == 10.0 { // score max calculated value
		cat.Details = "All required files found."
	} else {
		cat.Details = "Some required files are missing."
	}

	return cat
}
