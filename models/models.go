package models

type Status string

const (
	StatusPass    Status = "pass"
	StatusIgnored Status = "ignored"
	StatusError   Status = "error"
	StatusWarning Status = "warning"
)

// LineItem represents a specific check within a category
type LineItem struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Score       float64 `json:"score"`
	MaxScore    float64 `json:"max_score"`
	Status      Status  `json:"status"`
	Details     string  `json:"details"`
	DetailsArgs []any   `json:"details_args,omitempty"`
}

// Category represents a specific grading area
type Category struct {
	Name             string     `json:"name"`
	Description      string     `json:"description"`
	// Handled by grader automatically (summed from line items)
	Score            float64    `json:"score"`
	MaxScore         float64    `json:"max_score"`
	Weight           float64    `json:"weight"`
	// Handled by grader automatically (calculated as Score/MaxScore * Weight)
	WeightPercentage float64    `json:"weight_percentage"`
	// Optional override, handled by grader automatically if not set (pass/warning/error/ignored)
	Status           Status     `json:"status"`
	// Optional details or notes about the category
	Details          string     `json:"details"`
	DetailsArgs      []any      `json:"details_args,omitempty"`
	LineItems        []LineItem `json:"line_items"`
}

// Report represents the final grading output
type Report struct {
	TargetDirectory    string     `json:"target_directory"`
	AppName            string     `json:"app_name"`
	TotalScore         float64    `json:"total_score"`
	MaxTotalScore      float64    `json:"max_total_score"`
	UnweightedScore    float64    `json:"unweighted_score"`
	UnweightedMaxScore float64    `json:"unweighted_max_score"`
	TotalWeight        float64    `json:"total_weight"`
	Percentage         float64    `json:"percentage"`
	Categories         []Category `json:"categories"`
}
