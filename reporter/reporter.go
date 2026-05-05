package reporter

import (
	"appbuilder-grader/models"
	_ "embed"
	"encoding/json"
	"html/template"
	"os"
	"path/filepath"
)

// WriteJSON outputs the grading report to a JSON file
func WriteJSON(report *models.Report, outputPath string) error {
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(outputPath, data, 0644)
}

//go:embed template.html
var HTMLTemplate string

// WriteHTML outputs the grading report to an HTML file
func WriteHTML(report *models.Report, outputPath string) error {
	tmpl, err := template.New("report").Parse(HTMLTemplate)
	if err != nil {
		return err
	}

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return err
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, report)
}
