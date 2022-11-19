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

// 写入数据异常

// CommitsInfoCreate 在数据库中创建commit
func CommitsInfoCreate(context *gin.Context) {
	var t T2
	if err := context.ShouldBind(&t); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": utils.GetErrMsg(http.StatusBadRequest),
		})
		return
	}
	pid := t.Project.Pid
	version := t.Release.Version
	temp := ProjectsTable{}

	if err := Db.Table("projects").First(&temp, "project_id = ? ", pid).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": utils.GetErrMsg(http.StatusBadRequest),
		})
		return
	}
	release := ReleasesTable{}
	if err := Db.Table("releases").First(&release, "release_version = ?", version).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": utils.GetErrMsg(http.StatusBadRequest),
		})
		return
	}
	for i := 0; i < len(t.Commit); i++ {
		temp2 := CommitsTable{0, t.Commit[i].Hash, t.Commit[i].Time,
			t.Commit[i].Author, t.Commit[i].Email, int(release.TableID)}
		fmt.Println(Db.Table("commits").Create(&temp2).RowsAffected)
	}
	context.JSON(http.StatusOK, gin.H{
		"msg": utils.GetErrMsg(http.StatusOK),
	})
}
