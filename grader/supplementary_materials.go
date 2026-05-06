package grader

import (
	"appbuilder-grader/models"
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

	return models.LineItem{
		Name:        "line_items.glossary_name",
		Description: "line_items.glossary_desc",
		Score:       0.0,
		MaxScore:    3.0,
		Status:      "ignored",
		Details:     "details.glossary_details",
	}
}

func (g *Grader) checkIllustrations() models.LineItem {
	// 0=No illustrations
	// 1=BW line art
	// 2=Colour illustrations

	return models.LineItem{
		Name:        "line_items.illustrations_name",
		Description: "line_items.illustrations_desc",
		Score:       0.0,
		MaxScore:    2.0,
		Status:      "ignored",
		Details:     "details.illustrations_details",
	}
}

func (g *Grader) checkTopicalIndex() models.LineItem {
	return models.LineItem{
		Name:        "line_items.topical_index_name",
		Description: "line_items.topical_index_desc",
		Score:       0.0,
		MaxScore:    1.0,
		Status:      "ignored",
		Details:     "details.topical_index_details",
	}
}

func (g *Grader) checkReadingPlan() models.LineItem {
	return models.LineItem{
		Name:        "line_items.reading_plan_name",
		Description: "line_items.reading_plan_desc",
		Score:       0.0,
		MaxScore:    1.0,
		Status:      "ignored",
		Details:     "details.reading_plan_details",
	}
}

func (g *Grader) checkStudyBibleMaterial() models.LineItem {
	return models.LineItem{
		Name:        "line_items.study_bible_material_name",
		Description: "line_items.study_bible_material_desc",
		Score:       0.0,
		MaxScore:    1.0,
		Status:      "ignored",
		Details:     "details.study_bible_material_details",
	}
}