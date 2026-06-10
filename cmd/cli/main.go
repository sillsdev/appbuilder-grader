package main

import (
	"appbuilder-grader/cmd"
	"appbuilder-grader/reporter"
	"appbuilder-grader/runner"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if err := RunCLI(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func RunCLI(args []string) error {
	flags := flag.NewFlagSet("appbuilder-grader", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)

	targetDir := flags.String("dir", "", "The target directory containing the build output to grade")
	outputDir := flags.String("out", "out", "The directory to save the output reports")
	lang := flags.String("lang", "en", "The language to use for the report (default en)")

	flags.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of AppBuilder Grader:\n")
		flags.PrintDefaults()
	}

	if err := flags.Parse(args); err != nil {
		return err
	}

	fmt.Printf("Appbuilder Grader version %s\n", cmd.Version())

	if *targetDir == "" {
		flags.Usage()
		return errors.New("target directory is required")
	}

	fmt.Printf("Analyzing directory: %s\n", *targetDir)

	report, err := runner.Evaluate(*targetDir, *lang)
	if err != nil {
		return err
	}
	fmt.Printf("Grading Complete! Overall Percentage: %.2f%%\n", report.Percentage)

	jsonPath := ""
	htmlPath := ""

	if *outputDir != "" {
		jsonPath = filepath.Join(*outputDir, "report.json")
		htmlPath = filepath.Join(*outputDir, "report.html")

		if err := os.MkdirAll(*outputDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}

		if json, err := reporter.ExportJSON(report, jsonPath); err != nil {
			log.Printf("Failed to generate JSON report: %v\n", err)
		} else {
			if err := os.WriteFile(jsonPath, json, 0644); err != nil {
				log.Printf("Failed to save JSON report: %v\n", err)
			} else {
				log.Printf("JSON report saved to: %s\n", jsonPath)
			}
		}

		if html, err := reporter.ExportHTML(report, htmlPath); err != nil {
			log.Printf("Failed to generate HTML report: %v\n", err)
		} else {
			if err := os.WriteFile(htmlPath, html, 0644); err != nil {
				log.Printf("Failed to save HTML report: %v\n", err)
			} else {
				log.Printf("HTML report saved to: %s\n", htmlPath)
			}
		}
	}
	return nil
}