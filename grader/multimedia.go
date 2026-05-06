package grader

import (
	"appbuilder-grader/models"
)

func (g *Grader) checkMultimedia() models.Category {
	cat := models.Category{
		Name:        "categories.multimedia_name",
		Description: "categories.multimedia_desc",
		Weight:      1.0,
		LineItems:   make([]models.LineItem, 0),
	}

	cat.LineItems = append(cat.LineItems, g.checkAudio())
	cat.LineItems = append(cat.LineItems, g.checkStyleOfAudio())
	cat.LineItems = append(cat.LineItems, g.checkSynchronizedHighlighting())
	cat.LineItems = append(cat.LineItems, g.checkVideo())

	return cat
}

func (g *Grader) checkAudio() models.LineItem {
	return createIgnoredItem("line_items.audio_name", "line_items.audio_desc", 3.0)
}

func (g *Grader) checkStyleOfAudio() models.LineItem {
	return createIgnoredItem("line_items.style_of_audio_name", "line_items.style_of_audio_desc", 2.0)
}

func (g *Grader) checkSynchronizedHighlighting() models.LineItem {
	return createIgnoredItem("line_items.synchronized_highlighting_name", "line_items.synchronized_highlighting_desc", 1.0)
}

func (g *Grader) checkVideo() models.LineItem {
	return createIgnoredItem("line_items.video_name", "line_items.video_desc", 4.0)
}
