package modules

//这里是whosbug路由组所需的结构体

// Project 项目信息
type Project struct {
	Pid string `form:"pid" json:"pid" binding:"required"`
}

// Release 版本信息
type Release struct {
	Version    string `form:"version" json:"version" binding:"required"`
	CommitHash string `form:"last_commit_hash" json:"last_commit_hash" binding:"required"`
}

// Commit commit信息
type Commit struct {
	Hash   string `form:"hash" json:"hash" binding:"required"`
	Email  string `form:"email" json:"email" binding:"required"`
	Author string `form:"author" json:"author" binding:"required"`
	Time   string `form:"time" json:"time" binding:"required"`
}

// T 项目及版本
type T struct {
	Project Project
	Release Release
}

// T2 项目及版本、commit信息
type T2 struct {
	Project Project  `json:"project"`
	Release Release  `json:"release"`
	Commit  []Commit `json:"commits"`
}

// T3 项目、版本、object
type T3 struct {
	Project Project  `json:"project"`
	Release Release  `json:"release"`
	Object  []Object `json:"objects"`
}

// ReleaseModules 版本、最新commit、项目
type ReleaseModules struct {
	Version        string
	LastCommitHash string
	Project        int //or string?
}

// UncalculatedObjectInfo 未计算的项目信息
type UncalculatedObjectInfo struct {
	Hash             string `form:"hash" json:"hash" binding:"required"`
	ObjectID         string `form:"object_id" json:"object_id" binding:"required"`
	OldObjectID      string `form:"old_object_id " json:"old_object_id" binding:"required"`
	Parameters       string `form:"parameters " json:"parameters" binding:"required"`
	StartLine        int    `form:"start_line " json:"start_line" binding:"required"`
	Path             string `form:"path" json:"path" binding:"required"`
	EndLine          int    `form:"end_line" json:"end_line" binding:"required"`
	OldLineCount     int    `form:"old_line_count" json:"old_line_count" binding:"required"`
	NewLineCount     int    `form:"new_line_count" json:"current_line_count" binding:"required"`
	DeletedLineCount int    `form:"deleted_line_count" json:"removed_line_count" binding:"required"`
	AddedLineCount   int    `form:"added_line_count" json:"added_new_line_count" binding:"required"`
}

// Object object信息
type Object struct {
	Path          string
	ObjectID      string
	OwnerInfo     string
	OldConfidence float64
	StartLine     int
	EndLine       int
	parameters    string
}

// T4 项目、版本、未计算的object
type T4 struct {
	Project         Project                  `json:"project"`
	Release         Release                  `json:"release"`
	UncountedObject []UncalculatedObjectInfo `json:"objects"`
}

// Method 方法
type Method struct {
	MethodID   string `json:"method_id"`
	Filepath   string `json:"filepath"`
	Parameters string `json:"parameters"`
}

// GetConfidence 为得到置信度而创建的结构体
type GetConfidence struct {
	Project Project  `json:"project"`
	Release Release  `json:"release"`
	Method  []Method `json:"methods"`
}

// JSONRes jsonRes
type JSONRes struct {
	Status  string
	Message string
	Object  []NodesTable
}
