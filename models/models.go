package models

// LineItem represents a specific check within a category
type LineItem struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Score       float64 `json:"score"`
	MaxScore    float64 `json:"max_score"`
	Status      string  `json:"status"`
	Details     string  `json:"details"`
}

// Category represents a specific grading area
type Category struct {
	Name             string     `json:"name"`
	Description      string     `json:"description"`
	Score            float64    `json:"score"`
	MaxScore         float64    `json:"max_score"`
	Weight           float64    `json:"weight"`
	WeightPercentage float64    `json:"weight_percentage"`
	Status           string     `json:"status"`
	Details          string     `json:"details"`
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
