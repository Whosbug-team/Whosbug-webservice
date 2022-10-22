package views

import (
	"errors"
	"fmt"
	"net/http"
	. "webService_Refactoring/modules"
	"webService_Refactoring/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CommitsDiffsCreate 在数据库中创建commitdiff
func CommitsDiffsCreate(context *gin.Context) {
	var t T4
	if err := context.ShouldBind(&t); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": utils.GetErrMsg(http.StatusBadRequest),
		})
		return
	}
	pid := t.Project.Pid
	version := t.Release.Version
	project := ProjectsTable{}
	if err := Db.Table("projects").First(&project, "project_id = ? ", pid).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": utils.GetErrMsg(utils.ErrorProjectNotFound),
		})
		return
	}
	release := ReleasesTable{}
	if err := Db.Table("releases").First(&release, "release_version = ?", version).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": utils.GetErrMsg(utils.ErrorReleaseNotFound),
		})
		return
	}
	commit := CommitsTable{}
	Db.Table("commits").First(&commit, "release_table_id = ?", release.TableID)
	n, releaseID, commitID := len(t.UncountedObject), release.TableID, commit.TableID
	for i := 0; i < n; i++ {
		temp2 := ObjectsTable{0, t.UncountedObject[i].Parameters, t.UncountedObject[i].Hash,
			t.UncountedObject[i].StartLine, t.UncountedObject[i].EndLine, t.UncountedObject[i].Path,
			t.UncountedObject[i].ObjectID, t.UncountedObject[i].OldObjectID, t.UncountedObject[i].OldLineCount,
			t.UncountedObject[i].NewLineCount, t.UncountedObject[i].DeletedLineCount, t.UncountedObject[i].AddedLineCount,
			int(releaseID), int(commitID)}
		fmt.Println(Db.Table("objects").Create(&temp2).RowsAffected)
	}
	context.JSON(http.StatusOK, gin.H{
		"msg": utils.GetErrMsg(http.StatusOK),
	})
	return
}
