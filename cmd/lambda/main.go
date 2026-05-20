package main

import (
	"appbuilder-grader/models"
	"appbuilder-grader/runner"
	"context"
	"errors"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(HandleLambda)
}

type LambdaRequest struct {
	TargetDir string `json:"target_dir"`
	OutputDir string `json:"output_dir,omitempty"`
	Lang      string `json:"lang,omitempty"`
}

type LambdaResponse struct {
	Report   *models.Report `json:"report"`
	JSONPath string         `json:"json_path,omitempty"`
	HTMLPath string         `json:"html_path,omitempty"`
}

func HandleLambda(ctx context.Context, request LambdaRequest) (LambdaResponse, error) {
	_ = ctx

	request.TargetDir = strings.TrimSpace(request.TargetDir)
	request.OutputDir = strings.TrimSpace(request.OutputDir)
	request.Lang = strings.TrimSpace(request.Lang)

	if request.TargetDir == "" {
		return LambdaResponse{}, errors.New("target_dir is required")
	}

	report, jsonPath, htmlPath, err := runner.Evaluate(request.TargetDir, request.OutputDir, request.Lang)
	if err != nil {
		return LambdaResponse{}, err
	}

	return LambdaResponse{Report: report, JSONPath: jsonPath, HTMLPath: htmlPath}, nil
}