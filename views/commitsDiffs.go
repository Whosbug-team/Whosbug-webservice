package views

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	. "webService_Refactoring/modules"
)

// CommitsDiffsCreate 在数据库中创建commitdiff
func CommitsDiffsCreate(context *gin.Context) {

	var t T4

	err := context.ShouldBind(&t)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	pid := t.Project.Pid
	version := t.Release.Version
	temp := ProjectsTable{}
	res := Db.Table("projects").First(&temp, "project_id = ? ", pid)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		context.Status(400)
		return
	}
	temp1 := ReleasesTable{}
	res1 := Db.Table("releases").First(&temp1, "release_version = ?", version)
	if errors.Is(res1.Error, gorm.ErrRecordNotFound) {
		context.Status(400)
		return
	}
	commit := CommitsTable{}
	Db.Table("commits").First(&commit, "release_table_id = ?", temp1.TableID)
	n, releaseID, commitID := len(t.UncountedObject), temp1.TableID, commit.TableID
	for i := 0; i < n; i++ {
		temp2 := ObjectsTable{0, t.UncountedObject[i].Parameters, t.UncountedObject[i].Hash,
			t.UncountedObject[i].StartLine, t.UncountedObject[i].EndLine, t.UncountedObject[i].Path,
			t.UncountedObject[i].ObjectID, t.UncountedObject[i].OldObjectID, t.UncountedObject[i].OldLineCount,
			t.UncountedObject[i].NewLineCount, t.UncountedObject[i].DeletedLineCount, t.UncountedObject[i].AddedLineCount,
			int(releaseID), int(commitID)}
		fmt.Println(Db.Table("objects").Create(&temp2).RowsAffected)
	}

	context.Status(200)

}
