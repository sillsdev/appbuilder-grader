package main

import (
	"appbuilder-grader/reporter"
	"appbuilder-grader/runner"
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const projectsMountRoot = "/mnt/projects"

func main() {
	lambda.Start(HandleLambda)
}

/*
{
	reportId: grading.id,
	project: {
		id: grading.project.id,
		appId: grading.project.app_id,
		name: grading.project.project_name,
		languageCode: grading.project.language_code,
		s3Url: grading.project_url
	},
	publisherId: grading.publisher_id
}
*/

type Project struct {
	Id           int    `json:"id"`
	AppId        int    `json:"app_id"`
	Name         string `json:"name"`
	LanguageCode string `json:"language_code"`
	S3Url        string `json:"s3_url"`
}
type LambdaRequest struct {
	ReportId     int     `json:"report_id"`
	ReportPrefix string  `json:"report_prefix"`
	Project      Project `json:"project"`
	PublisherId  int     `json:"publisher_id"`
}

type LambdaResponse struct {
	JSONPath   string  `json:"json_path,omitempty"`
	HTMLPath   string  `json:"html_path,omitempty"`
	TotalScore float64 `json:"total_score,omitempty"`
}

func HandleLambda(ctx context.Context, request LambdaRequest) (LambdaResponse, error) {
	// s3 projects bucket is mounted at /mnt/projects
	// s3 url is of the form s3://{{bucket}}/{{key}}
	// bucket is ENV.PROJECTS_BUCKET

	bucket, objectPath, err := parseS3URL(request.Project.S3Url)
	if err != nil {
		return LambdaResponse{}, err
	}
	if bucket != os.Getenv("PROJECTS_BUCKET") {
		return LambdaResponse{}, errors.New("Invalid s3 URL: bucket does not match. Grader is not currently compatible with this bucket")
	}

	targetDir, err := projectMountPath(objectPath)
	if err != nil {
		return LambdaResponse{}, err
	}

	// Only English supported at this time
	lang := "en"

	report, err := runner.Evaluate(targetDir, lang)
	if err != nil {
		return LambdaResponse{}, err
	}
	fmt.Printf("Grading Complete! Overall Percentage: %.2f%%\n", report.Percentage)

	json, err := reporter.ExportJSON(report, "")
	if err != nil {
		return LambdaResponse{}, err
	}

	html, err := reporter.ExportHTML(report, "")
	if err != nil {
		return LambdaResponse{}, err
	}

	awsCfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return LambdaResponse{}, err
	}
	fmt.Println("Exported JSON and HTML. Uploading to S3...")
	client := s3.NewFromConfig(awsCfg)

	// Upload to S3
	// object path is /reports/{{reportId}}/report.json
	jsonPath := "reports/" + fmt.Sprintf("%d", request.ReportId) + "/report.json"
	htmlPath := "reports/" + fmt.Sprintf("%d", request.ReportId) + "/report.html"
	if err := uploadToS3(ctx, client, json, os.Getenv("ARTIFACTS_BUCKET"), jsonPath, "application/json; charset=utf-8"); err != nil {
		return LambdaResponse{}, err
	}
	if err := uploadToS3(ctx, client, html, os.Getenv("ARTIFACTS_BUCKET"), htmlPath, "text/html; charset=utf-8"); err != nil {
		return LambdaResponse{}, err
	}

	return LambdaResponse{JSONPath: jsonPath, HTMLPath: htmlPath, TotalScore: report.TotalScore}, nil
}

func projectMountPath(objectPath string) (string, error) {
	cleaned := path.Clean(path.Join(projectsMountRoot, objectPath))
	if cleaned != projectsMountRoot && !strings.HasPrefix(cleaned, projectsMountRoot+"/") {
		return "", errors.New("invalid project path")
	}
	return cleaned, nil
}

func parseS3URL(raw string) (string, string, error) {
	parsed, err := url.Parse(raw)
	if err != nil {
		return "", "", errors.New("invalid s3 URL")
	}
	if parsed.Scheme != "s3" || parsed.Host == "" || parsed.Path == "" {
		return "", "", errors.New("invalid s3 URL")
	}

	objectPath := strings.TrimPrefix(parsed.Path, "/")
	if objectPath == "" {
		return "", "", errors.New("invalid s3 URL")
	}

	return parsed.Host, objectPath, nil
}

func uploadToS3(ctx context.Context, client *s3.Client, data []byte, bucket string, key string, contentType string) error {
	fmt.Printf("Uploading %s to s3 (content size = %.2f KB)...\n", bucket + "/" + key, float64(len(data)) / 1024.0)
	_, err := client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      &bucket,
		Key:         &key,
		Body:        bytes.NewReader(data),
		ContentType: &contentType,
	})
	return err
}
