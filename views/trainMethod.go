package views

import (
	"errors"
	"net/http"
	. "webService_Refactoring/modules"
	"webService_Refactoring/utils"

	"github.com/cheggaaa/pb"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CommitsTrainMethodCreate 训练参数数据集
func CommitsTrainMethodCreate(context *gin.Context) {
	var t T
	if err := context.ShouldBind(&t); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": utils.GetErrMsg(http.StatusBadRequest),
		})
		return
	}
	pid, version, temp := t.Project.Pid, t.Release.Version, ProjectsTable{}
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
	object := ObjectsTable{}
	lastCommitHash := t.Release.CommitHash
	if err := Db.Table("objects").First(&object, "release_version = ? and hash = ?", version, lastCommitHash); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Delete error",
		})
		return
	}
	context.Status(200)
	count := 100

	// 创建进度条并开始
	bar := pb.StartNew(count)

	for i := 0; i < count; i++ {
		bar.Increment()
		//time.Sleep(50 * time.Microsecond)
	}

	// 结束进度条
	bar.Finish()

}
