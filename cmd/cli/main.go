package main

import (
	"appbuilder-grader/runner"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
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

	if *targetDir == "" {
		flags.Usage()
		return errors.New("target directory is required")
	}

	fmt.Printf("Analyzing directory: %s\n", *targetDir)

	report, _, _, err := runner.Evaluate(*targetDir, *outputDir, *lang)
	if err != nil {
		return err
	}

	fmt.Printf("Grading Complete! Overall Percentage: %.2f%%\n", report.Percentage)
	return nil
}