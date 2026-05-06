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
	item := models.LineItem{
		Name:        "line_items.audio_name",
		Description: "line_items.audio_desc",
		MaxScore:    3.0,
	}
	item.Score = 0.0
	item.Status = models.StatusIgnored
	item.Details = "details.not_implemented_yet"
	return item
}

func (g *Grader) checkStyleOfAudio() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.style_of_audio_name",
		Description: "line_items.style_of_audio_desc",
		MaxScore:    2.0,
	}
	item.Score = 0.0
	item.Status = models.StatusIgnored
	item.Details = "details.not_implemented_yet"
	return item
}

func (g *Grader) checkSynchronizedHighlighting() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.synchronized_highlighting_name",
		Description: "line_items.synchronized_highlighting_desc",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = models.StatusIgnored
	item.Details = "details.not_implemented_yet"
	return item
}

func (g *Grader) checkVideo() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.video_name",
		Description: "line_items.video_desc",
		MaxScore:    4.0,
	}
	item.Score = 0.0
	item.Status = models.StatusIgnored
	item.Details = "details.not_implemented_yet"
	return item
}
