package grader

import (
	"appbuilder-grader/models"
)

func (g *Grader) checkMultimedia() models.Category {
	cat := models.Category{
		Name:        "Multimedia",
		Description: "Checks for the presence and configuration of audio and video features",
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
		Name:        "Audio",
		Description: "Checks for the presence and extent of audio (0=No audio, 1=Some books, 2=All books, 3=Glossary audio)",
		MaxScore:    3.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}

func (g *Grader) checkStyleOfAudio() models.LineItem {
	item := models.LineItem{
		Name:        "Style of Audio",
		Description: "Checks audio style (0=No audio, 1=Single voice, 2=Multi-voice/dramatized)",
		MaxScore:    2.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}

func (g *Grader) checkSynchronizedHighlighting() models.LineItem {
	item := models.LineItem{
		Name:        "Synchronized Highlighting",
		Description: "Checks for highlighting of verses with audio (0=No, 1=Yes)",
		MaxScore:    1.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}

func (g *Grader) checkVideo() models.LineItem {
	item := models.LineItem{
		Name:        "Video",
		Description: "Checks video links/clips presence (0=No, 1=Jesus film Luke, 2=captions, 3=other gospels, 4=other links)",
		MaxScore:    4.0,
	}
	item.Score = 0.0
	item.Status = "ignored"
	item.Details = "Not implemented yet"
	return item
}
