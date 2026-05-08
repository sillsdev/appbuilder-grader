package grader

import (
	"os"
	"path/filepath"
	"strings"
)

type projectBookFile struct {
	CollectionID   string
	CollectionName string
	WritingSystem  string
	BookID         string
	Path           string
}

func (g *Grader) dataDir() string {
	if g.TargetDir == "" {
		return ""
	}

	appDefFiles, err := filepath.Glob(filepath.Join(g.TargetDir, "*.appDef"))
	if err == nil && len(appDefFiles) > 0 {
		base := strings.TrimSuffix(filepath.Base(appDefFiles[0]), filepath.Ext(appDefFiles[0]))
		candidate := filepath.Join(g.TargetDir, base+"_data")
		if isDir(candidate) {
			return candidate
		}
	}

	entries, err := os.ReadDir(g.TargetDir)
	if err != nil {
		return ""
	}
	for _, entry := range entries {
		if entry.IsDir() && strings.HasSuffix(entry.Name(), "_data") {
			return filepath.Join(g.TargetDir, entry.Name())
		}
	}
	return ""
}

func (g *Grader) bookFiles() []projectBookFile {
	dataDir := g.dataDir()
	if dataDir == "" {
		return nil
	}

	files := make([]projectBookFile, 0)
	for _, collection := range g.AppDef.Books {
		for _, book := range collection.Book {
			if strings.TrimSpace(book.Filename) == "" {
				continue
			}
			path := filepath.Join(dataDir, "books", collection.Id, filepath.FromSlash(book.Filename))
			if !fileExists(path) {
				path = findFile(filepath.Join(dataDir, "books"), book.Filename)
			}
			if path == "" {
				continue
			}
			files = append(files, projectBookFile{
				CollectionID:   collection.Id,
				CollectionName: collection.BookCollectionName,
				WritingSystem:  collection.WritingSystem.Code,
				BookID:         strings.ToUpper(book.Id),
				Path:           path,
			})
		}
	}
	return files
}

func readTextFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(content)
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func findFile(root, filename string) string {
	filename = filepath.Base(filepath.FromSlash(filename))
	var found string
	_ = filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || found != "" {
			return nil
		}
		if strings.EqualFold(d.Name(), filename) {
			found = path
		}
		return nil
	})
	return found
}
