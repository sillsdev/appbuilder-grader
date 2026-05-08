package grader

import (
	"appbuilder-grader/models"
	"fmt"
	"path/filepath"
	"strings"
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
		MaxScore:    4.0,
		Status:      models.StatusWarning,
		Details:     "details.audio_missing",
	}

	audioBooks := g.audioBookIDs()
	standardAudioBooks := 0
	for bookID := range audioBooks {
		if isNTBook(bookID) || isOTBook(bookID) {
			standardAudioBooks++
		}
	}
	if standardAudioBooks == 0 {
		return item
	}

	item.Score = 1.0
	item.Status = models.StatusPass
	item.SetDetails("details.audio_some_books", standardAudioBooks)

	if allBooksHaveAudio(audioBooks, AllNTBooks) {
		item.Score = 2.0
		item.SetDetails("details.audio_all_nt", len(AllNTBooks))
	}
	if allBooksHaveAudio(audioBooks, standardBibleBooks()) {
		item.Score = 3.0
		item.SetDetails("details.audio_all_standard", len(AllOTBooks)+len(AllNTBooks))
	}
	if audioBooks["GLO"] {
		item.Score = 4.0
		item.SetDetails("details.audio_glossary", standardAudioBooks)
	}

	return item
}

func (g *Grader) checkStyleOfAudio() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.style_of_audio_name",
		Description: "line_items.style_of_audio_desc",
		MaxScore:    2.0,
		Status:      models.StatusWarning,
		Details:     "details.audio_style_missing",
	}

	if g.audioEntryCount() == 0 {
		return item
	}

	item.Score = 1.0
	item.Status = models.StatusPass
	item.SetDetails("details.audio_style_single_voice")
	if g.hasDramatizedAudioEvidence() {
		item.Score = 2.0
		item.SetDetails("details.audio_style_dramatized")
	}
	return item
}

func (g *Grader) checkSynchronizedHighlighting() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.synchronized_highlighting_name",
		Description: "line_items.synchronized_highlighting_desc",
		MaxScore:    1.0,
		Status:      models.StatusWarning,
		Details:     "details.timing_missing",
	}

	totalAudio, timedAudio := g.audioTimingCounts()
	if totalAudio == 0 {
		return item
	}

	percent := (float64(timedAudio) / float64(totalAudio)) * 100
	item.SetDetails("details.timing_coverage", timedAudio, totalAudio, fmt.Sprintf("%.1f", percent))
	if percent > 60 {
		item.Score = 1.0
		item.Status = models.StatusPass
	}
	return item
}

func (g *Grader) checkVideo() models.LineItem {
	item := models.LineItem{
		Name:        "line_items.video_name",
		Description: "line_items.video_desc",
		MaxScore:    4.0,
		Status:      models.StatusWarning,
		Details:     "details.video_missing",
	}

	lukeJesusFilm := 0
	otherGospelJesusFilm := 0
	otherVideo := 0
	for _, video := range g.AppDef.Videos.Video {
		if strings.TrimSpace(video.OnlineUrl) == "" && strings.TrimSpace(video.Id) == "" {
			continue
		}
		bookID := videoBookID(video)
		if isJesusFilmVideo(video) {
			if bookID == "LUK" {
				lukeJesusFilm++
			} else if bookID == "MAT" || bookID == "MRK" || bookID == "JHN" {
				otherGospelJesusFilm++
			}
		} else {
			otherVideo++
		}
	}

	switch {
	case otherGospelJesusFilm > 0 && otherVideo > 0:
		item.Score = 4.0
		item.Status = models.StatusPass
		item.SetDetails("details.video_other_links", lukeJesusFilm, otherGospelJesusFilm, otherVideo)
	case otherGospelJesusFilm > 0:
		item.Score = 3.0
		item.Status = models.StatusPass
		item.SetDetails("details.video_other_gospels", lukeJesusFilm, otherGospelJesusFilm)
	case lukeJesusFilm > 0:
		item.Score = 1.0
		item.Status = models.StatusPass
		item.SetDetails("details.video_luke_jesus_film", lukeJesusFilm)
	case otherVideo > 0:
		item.Score = 1.0
		item.Status = models.StatusPass
		item.SetDetails("details.video_generic_links", otherVideo)
	}
	return item
}

func (g *Grader) audioBookIDs() map[string]bool {
	audioBooks := make(map[string]bool)
	for _, collection := range g.AppDef.Books {
		for _, book := range collection.Book {
			bookID := strings.ToUpper(book.Id)
			if bookID != "" && bookHasAudio(book) {
				audioBooks[bookID] = true
			}
		}
	}
	return audioBooks
}

func (g *Grader) audioEntryCount() int {
	count := 0
	for _, collection := range g.AppDef.Books {
		for _, book := range collection.Book {
			for _, audio := range book.Audio {
				if audioEntryExists(audio) {
					count++
				}
			}
		}
	}
	return count
}

func (g *Grader) audioTimingCounts() (int, int) {
	totalAudio := 0
	timedAudio := 0
	for _, collection := range g.AppDef.Books {
		for _, book := range collection.Book {
			for _, audio := range book.Audio {
				if !audioEntryExists(audio) {
					continue
				}
				totalAudio++
				if g.timingExists(audio.TimingFilename) {
					timedAudio++
				}
			}
		}
	}
	return totalAudio, timedAudio
}

func (g *Grader) timingExists(filename string) bool {
	filename = strings.TrimSpace(filename)
	if filename == "" {
		return false
	}
	if resolved := g.resolveDataFile("timings", filename); resolved != "" {
		return true
	}
	for _, path := range g.filesUnderDataDir("timings") {
		if strings.EqualFold(filepath.Base(path), filepath.Base(filename)) {
			return true
		}
	}
	return false
}

func (g *Grader) hasDramatizedAudioEvidence() bool {
	for _, source := range g.AppDef.AudioSources.AudioSource {
		haystack := strings.ToLower(strings.Join([]string{
			source.Id,
			source.Type,
			source.Name,
			source.Folder,
			source.Address,
			source.DamID,
		}, " "))
		if strings.Contains(haystack, "drama") ||
			strings.Contains(haystack, "dramati") ||
			strings.Contains(haystack, "multi-voice") ||
			strings.Contains(haystack, "multivoice") ||
			strings.HasSuffix(strings.ToUpper(strings.TrimSpace(source.DamID)), "DA") {
			return true
		}
	}
	return false
}

func bookHasAudio(book models.Book) bool {
	for _, audio := range book.Audio {
		if audioEntryExists(audio) {
			return true
		}
	}
	return false
}

func audioEntryExists(audio models.Audio) bool {
	return strings.TrimSpace(audio.Chapter) != "" ||
		strings.TrimSpace(audio.Filename.Value) != "" ||
		strings.TrimSpace(audio.TimingFilename) != ""
}

func allBooksHaveAudio(audioBooks map[string]bool, requiredBooks []string) bool {
	for _, bookID := range requiredBooks {
		if !audioBooks[bookID] {
			return false
		}
	}
	return true
}

func standardBibleBooks() []string {
	books := make([]string, 0, len(AllOTBooks)+len(AllNTBooks))
	books = append(books, AllOTBooks...)
	books = append(books, AllNTBooks...)
	return books
}

func isJesusFilmVideo(video models.Video) bool {
	haystack := strings.ToLower(strings.Join([]string{
		video.Id,
		video.Title,
		video.Thumbnail,
		video.OnlineUrl,
	}, " "))
	return strings.Contains(haystack, "jesusfilm") ||
		strings.Contains(haystack, "arclight") ||
		strings.Contains(haystack, "jf")
}

func videoBookID(video models.Video) string {
	ref := strings.ToUpper(video.Placement.Ref)
	if separator := strings.Index(ref, "|"); separator >= 0 {
		ref = ref[separator+1:]
	}
	if dot := strings.Index(ref, "."); dot >= 0 {
		return ref[:dot]
	}
	return ""
}
