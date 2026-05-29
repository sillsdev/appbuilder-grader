package reporter

import (
	"appbuilder-grader/loc"
	"appbuilder-grader/models"
	"bytes"
	_ "embed"
	"encoding/json"
	"html/template"
)

// ExportJSON outputs the grading report to a JSON file
func ExportJSON(report *models.Report, outputPath string) ([]byte, error) {
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return nil, err
	}
	// return os.WriteFile(outputPath, data, 0644)
	return data, nil
}

//go:embed template.html
var HTMLTemplate string

// ExportHTML outputs the grading report to an HTML file
func ExportHTML(report *models.Report, outputPath string) ([]byte, error) {
	funcMap := template.FuncMap{
		"t": loc.T,
		"safeHTML": func(s string) template.HTML {
			return template.HTML(s)
		},
	}

	tmpl, err := template.New("report").Funcs(funcMap).Parse(HTMLTemplate)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, report)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
