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
	return createIgnoredItem("line_items.glossary_name", "line_items.glossary_desc", 3.0)
}

func (g *Grader) checkIllustrations() models.LineItem {
	// 0=No illustrations
	// 1=BW line art
	// 2=Colour illustrations
	return createIgnoredItem("line_items.illustrations_name", "line_items.illustrations_desc", 2.0)
}

func (g *Grader) checkTopicalIndex() models.LineItem {
	return createIgnoredItem("line_items.topical_index_name", "line_items.topical_index_desc", 1.0)
}

func (g *Grader) checkReadingPlan() models.LineItem {
	return createIgnoredItem("line_items.reading_plan_name", "line_items.reading_plan_desc", 1.0)
}

func (g *Grader) checkStudyBibleMaterial() models.LineItem {
	return createIgnoredItem("line_items.study_bible_material_name", "line_items.study_bible_material_desc", 1.0)
}
