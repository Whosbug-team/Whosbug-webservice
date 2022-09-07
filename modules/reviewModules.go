package modules

// ReviewerModules reviewer模块
type ReviewerModules struct {
	Project    string `json:"project" form:"project" binding:"required"` //or string?
	FilePath   string `json:"file_path" form:"file_path" binding:"required"`
	Reviewer   string `json:"reviewer" form:"reviewer" binding:"required"`
	ReviewRule int    `json:"review_rule" form:"review_rule" binding:"required"`
}

// RuleModules rule模块
type RuleModules struct {
	Project string `json:"project" form:"project" binding:"required"` //or string?
	File    string `json:"file" form:"file" binding:"required"`
}
