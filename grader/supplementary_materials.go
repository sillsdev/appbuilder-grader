package grader

import (
	"appbuilder-grader/models"
)

func (g *Grader) checkSupplementaryMaterials() models.Category {
	cat := models.Category{
		Name:        "Supplementary Materials",
		Description: "Evaluate the presence and quality of supplementary materials (ex. glossary, illustrations, reading plan, etc)",
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

	return models.LineItem{
		Name:        "Glossary",
		Description: "Check for presence and quality of glossary",
		Score:       0.0,
		MaxScore:    3.0,
		Status:      "ignored",
		Details:     "Glossary check not implemented yet",
	}
}

func (g *Grader) checkIllustrations() models.LineItem {
	// 0=No illustrations
	// 1=BW line art
	// 2=Colour illustrations

	return models.LineItem{
		Name:        "Illustrations",
		Description: "Check for presence and quality of illustrations",
		Score:       0.0,
		MaxScore:    2.0,
		Status:      "ignored",
		Details:     "Illustrations check not implemented yet",
	}
}

func (g *Grader) checkTopicalIndex() models.LineItem {
	return models.LineItem{
		Name:        "Topical Index",
		Description: "Check for presence of topical index",
		Score:       0.0,
		MaxScore:    1.0,
		Status:      "ignored",
		Details:     "Topical index check not implemented yet",
	}
}

func (g *Grader) checkReadingPlan() models.LineItem {
	return models.LineItem{
		Name:        "Reading Plan",
		Description: "Check for presence of daily reading plan",
		Score:       0.0,
		MaxScore:    1.0,
		Status:      "ignored",
		Details:     "Reading plan check not implemented yet",
	}
}

func (g *Grader) checkStudyBibleMaterial() models.LineItem {
	return models.LineItem{
		Name:        "Study Bible Material",
		Description: "Check for presence of study Bible material (ex. notes, maps, charts, etc)",
		Score:       0.0,
		MaxScore:    1.0,
		Status:      "ignored",
		Details:     "Study Bible material check not implemented yet",
	}
}