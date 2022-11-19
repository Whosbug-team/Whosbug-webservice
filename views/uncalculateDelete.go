package views

import (
	"errors"
	"net/http"
	. "webService_Refactoring/modules"
	"webService_Refactoring/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UnCalculateDelete 接收完数据之后删除为计算的object信息
func UnCalculateDelete(context *gin.Context) {
	//接收数据
	var t T
	if err := context.ShouldBind(&t); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": utils.GetErrMsg(http.StatusBadRequest),
		})
		return
	}
	//提取pid、version
	pid := t.Project.Pid
	version := t.Release.Version
	//以pid去找
	project := ProjectsTable{}
	if err := Db.Table("projects").Where("project_id = ?", pid).First(&project).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": utils.GetErrMsg(utils.ErrorProjectNotFound),
		})
		return
	}
	//以version去找
	release := ReleasesTable{}
	err := Db.Table("releases").Where("release_version = ? and project_table_id = ?",
		version, project.TableID).First(&release).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": utils.GetErrMsg(utils.ErrorReleaseNotFound),
		})
		return
	}
	realRelease, uncounted, commit := ReleasesTable{}, ObjectsTable{}, CommitsTable{}
	Db.Table("releases").First(&realRelease, "release_version = ?", version)
	releaseID := realRelease.TableID
	Db.Table("commits").First(&commit, "release_table_id = ?", releaseID)
	uncountedID := commit.TableID
	res5 := Db.Table("objects").Delete(&uncounted, "commit_table_id = ?", uncountedID)
	if res5.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Delete error",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"msg": utils.GetErrMsg(http.StatusOK),
	})
}
