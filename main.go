package main

import (
	"appbuilder-grader/grader"
	"appbuilder-grader/reporter"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	targetDir := flag.String("dir", "", "The target directory containing the build output to grade")
	outputDir := flag.String("out", "out", "The directory to save the output reports")
	
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of AppBuilder Grader:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *targetDir == "" {
		fmt.Println("Error: Target directory is required.")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("Analyzing directory: %s\n", *targetDir)

	g := grader.NewGrader(*targetDir)
	report, err := g.Evaluate()
	if err != nil {
		log.Fatalf("Failed to grade directory: %v", err)
	}

	fmt.Printf("Grading Complete! Overall Percentage: %.2f%%\n", report.Percentage)

	jsonPath := filepath.Join(*outputDir, "report.json")
	htmlPath := filepath.Join(*outputDir, "report.html")

	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Save JSON
	if err := reporter.WriteJSON(report, jsonPath); err != nil {
		log.Printf("Failed to write JSON report: %v", err)
	} else {
		fmt.Printf("JSON report saved to: %s\n", jsonPath)
	}

	// Save HTML
	if err := reporter.WriteHTML(report, htmlPath); err != nil {
		log.Printf("Failed to write HTML report: %v", err)
	} else {
		fmt.Printf("HTML report saved to: %s\n", htmlPath)
	}
}
