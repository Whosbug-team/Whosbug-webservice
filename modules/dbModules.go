package modules

// 这里时需要与数据库交互的结构体

import (
	"github.com/google/uuid"
)

// UsersTable 用户表单
type UsersTable struct {
	UserID        uuid.UUID `gorm:"primaryKey;column:user_id"`
	UserName      string    `gorm:"type:varchar(150)"`
	UserFirstName string    `gorm:"type:varchar(150)"`
	UserLastName  string    `gorm:"type:varchar(150)"`
	UserEmail     string    `gorm:"type:varchar(255)"`
	UserToken     string    `gorm:"type:varchar(40)"`
	UserPassword  string    `gorm:"type:varchar(128)"`
}

// CommitsTable commit表单
type CommitsTable struct {
	TableID        uint64 `gorm:"primaryKey;auto_increment;type:bigserial;column:table_id"`
	Hash           string `gorm:"type:varchar(1000)"`
	Time           string `gorm:"type:varchar(1000)"`
	Author         string `gorm:"type:varchar(1000)"`
	Email          string `gorm:"type:varchar(1000)"`
	ReleaseTableID int    `gorm:"type:int;column:release_table_id"`
}

// ProjectsTable 项目表单
type ProjectsTable struct {
	TableID   uint64 `gorm:"primaryKey;auto_increment;type:serial;column:table_id"`
	ProjectID string `gorm:"type:varchar(200);column:project_id"`
}

// ReleasesTable 版本表单
type ReleasesTable struct {
	TableID        uint64 `gorm:"primaryKey;auto_increment;type:serial;column:table_id"`
	ReleaseVersion string `gorm:"type:varchar(200)"`
	LastCommitHash string `gorm:"type:varchar(1000)"`
	ProjectTableID int    `gorm:"type:int;column:project_table_id"`
}

// ObjectsTable 未计算置信度的object表单
type ObjectsTable struct {
	TableID         uint64 `gorm:"primaryKey;auto_increment;type:serial;column:table_id"`
	Parameters      string `gorm:"type:varchar(10000)"`
	Hash            string `gorm:"type:varchar(1000)"`
	StartLine       int    `gorm:"type:int"`
	EndLine         int    `gorm:"type:int"`
	ObjectPath      string `gorm:"type:varchar(1000)"`
	CurrentObjectID string `gorm:"type:varchar(1000);column:current_object_id"`
	FatherObjectID  string `gorm:"type:varchar(1000);column:father_object_id"`
	OldLine         int    `gorm:"type:int"`
	NewLine         int    `gorm:"type:int"`
	DeletedLine     int    `gorm:"type:int"`
	AddedLine       int    `gorm:"type:int"`
	ReleaseTableID  int    `gorm:"type:int;column:release_table_id"`
	CommitTableID   int    `gorm:"type:varchar(200);column:commit_table_id"`
}

// NodesTable 已经计算置信度的object表单
type NodesTable struct {
	TableID          uint64  `gorm:"primaryKey;auto_increment;type:serial;column:table_id"`
	ObjectPath       string  `gorm:"type:varchar(1000)"`
	ObjectParameters string  `gorm:"type:varchar(10000)"`
	CurrentObjectID  string  `gorm:"type:varchar(1000);column:current_object_id"`
	FatherObjectID   string  `gorm:"type:varchar(1000);column:father_object_id"`
	OldConfidence    float64 `gorm:"type:double"`
	NewConfidence    float64 `gorm:"type:double"`
	CommitTableID    int     `gorm:"type:int;column:commit_table_id"`
	ObjectTableID    int     `gorm:"type:int;column:object_table_id"`
	ObjectNewLine    int     `gorm:"type:int"`
	ObjectOldLine    int     `gorm:"type:int"`
	ObjectAddLine    int     `gorm:"type:int;column:object_ad_line"`
	ObjectDeleteLine int     `gorm:"type:int;column:object_de_line"`
}
