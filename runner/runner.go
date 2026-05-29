package runner

import (
	"appbuilder-grader/grader"
	"appbuilder-grader/loc"
	"appbuilder-grader/models"
	"log"
)

func Evaluate(targetDir, lang string) (*models.Report, error) {
	if lang == "" {
		lang = "en"
	}

	if err := loc.Init(lang); err != nil {
		log.Printf("Warning: Failed to load language '%s': %v\n", lang, err)
	}

	g := grader.NewGrader(targetDir)
	report, err := g.Evaluate()
	if err != nil {
		return nil, err
	}

	return report, nil
}
