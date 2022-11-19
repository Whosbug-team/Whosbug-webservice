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

// CreateProjectRelease 生成project&release
func CreateProjectRelease(context *gin.Context) {
	var t T
	if err := context.ShouldBind(&t); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": utils.GetErrMsg(http.StatusBadRequest),
		})
		return
	}
	pid := t.Project.Pid
	releaseVersion := t.Release.Version
	releaseHash := t.Release.CommitHash
	// 数据库查询pid，若存在且数据库中last_commit_hash 为传递的last_commit_hash
	// 不新建project并返回404
	project := ProjectsTable{}
	if err := Db.Table("projects").Where("project_id = ?", pid).First(&project).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		project.ProjectID = pid
		fmt.Println(Db.Table("projects").Create(&project).RowsAffected)
	}
	release := ReleasesTable{}
	err := Db.Table("releases").Where("release_version = ? "+
		"and last_commit_hash = ?", releaseVersion, releaseHash).First(&release).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		release.ProjectTableID = int(project.TableID)
		release.ReleaseVersion = releaseVersion
		release.LastCommitHash = releaseHash
		fmt.Println(Db.Table("releases").Create(&release).RowsAffected)
	} else {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "The Project and Release already exist, update the commit pid " + t.Project.Pid +
				" release: " + t.Release.Version + ", commit_hash: " + t.Release.CommitHash,
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"msg": utils.GetErrMsg(http.StatusCreated),
	})
}
