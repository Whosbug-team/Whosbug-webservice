package views

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	. "webService_Refactoring/modules"
)

// CreateProjectRelease 生成project&release
func CreateProjectRelease(context *gin.Context) {
	var t T
	err1 := context.ShouldBind(&t)
	if err1 != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err1.Error(),
		})
		return
	}
	pid := t.Project.Pid
	releaseVersion := t.Release.Version
	releaseHash := t.Release.CommitHash
	// 数据库查询pid，若存在且数据库中last_commit_hash 为传递的last_commit_hash
	// 不新建project并返回404
	project := ProjectsTable{}
	res := Db.Table("projects").Where("project_id = ?", pid).First(&project)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		project.ProjectID = pid
		fmt.Println(Db.Table("projects").Create(&project).RowsAffected)
	}
	release := ReleasesTable{}
	res2 := Db.Table("releases").Where("release_version = ? "+
		"and last_commit_hash = ?", releaseVersion, releaseHash).First(&release)
	if errors.Is(res2.Error, gorm.ErrRecordNotFound) {
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
	context.Status(201)
}
