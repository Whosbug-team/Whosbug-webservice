package views

import (
	"errors"
	"github.com/cheggaaa/pb"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	. "webService_Refactoring/modules"
)

// CommitsTrainMethodCreate 训练参数数据集
func CommitsTrainMethodCreate(context *gin.Context) {

	var t T

	err := context.ShouldBind(&t)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	pid, version, temp := t.Project.Pid, t.Release.Version, ProjectsTable{}
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
	temp3 := ObjectsTable{}
	lastCommitHash := t.Release.CommitHash
	errs := Db.Table("objects").First(&temp3, "release_version = ? and hash = ?", version, lastCommitHash)
	if errs != nil {
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
