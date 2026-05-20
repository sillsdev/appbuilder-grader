package runner

import (
	"appbuilder-grader/grader"
	"appbuilder-grader/loc"
	"appbuilder-grader/models"
	"appbuilder-grader/reporter"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Evaluate(targetDir, outputDir, lang string) (*models.Report, string, string, error) {
	if lang == "" {
		lang = "en"
	}

	if err := loc.Init(lang); err != nil {
		log.Printf("Warning: Failed to load language '%s': %v\n", lang, err)
	}

	g := grader.NewGrader(targetDir)
	report, err := g.Evaluate()
	if err != nil {
		return nil, "", "", err
	}

	jsonPath := ""
	htmlPath := ""
	if outputDir != "" {
		jsonPath = filepath.Join(outputDir, "report.json")
		htmlPath = filepath.Join(outputDir, "report.html")

		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return nil, "", "", fmt.Errorf("failed to create output directory: %w", err)
		}

		if err := reporter.WriteJSON(report, jsonPath); err != nil {
			log.Printf("Failed to write JSON report: %v\n", err)
		} else {
			log.Printf("JSON report saved to: %s\n", jsonPath)
		}

		if err := reporter.WriteHTML(report, htmlPath); err != nil {
			log.Printf("Failed to write HTML report: %v\n", err)
		} else {
			log.Printf("HTML report saved to: %s\n", htmlPath)
		}
	}

	return report, jsonPath, htmlPath, nil
}
