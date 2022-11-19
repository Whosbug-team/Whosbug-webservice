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

// CommitsUploadDoneCreate 上传完数据后计算置信度，然后插入数据库
func CommitsUploadDoneCreate(context *gin.Context) {
	var t T
	if err := context.ShouldBind(&t); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": utils.GetErrMsg(http.StatusBadRequest),
		})
		return
	}
	fmt.Println(t)
	pid, version, temp := t.Project.Pid, t.Release.Version, ProjectsTable{}
	if err := Db.Table("projects").First(&temp, "project_id = ? ", pid).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		context.Status(400)
		return
	}
	release := ReleasesTable{}
	if err := Db.Table("releases").First(&release, "release_version = ?", version).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		context.Status(400)
		return
	}

	//temp2 := ObjectsTable{}
	var object []ObjectsTable
	Db.Table("objects").Find(&object)
	n := len(object)

	for i := 0; i < n; i++ {
		var nodes []NodesTable
		Db.Table("nodes").Find(&nodes)
		uncalculateObject := UncalculateObjectInfo{object[i].Hash, object[i].CurrentObjectID, object[i].FatherObjectID,
			object[i].Parameters, object[i].StartLine, object[i].EndLine, object[i].OldLine,
			object[i].NewLine, object[i].DeletedLine, object[i].AddedLine}
		//var nodes1 []NodesTable
		var tnum int

		num, tnum := judgeObject(object[i], nodes)
		fmt.Println("n:", tnum)
		if num != 0 { //有object
			t := nodes[tnum].OldConfidence
			nodes[tnum].OldConfidence = nodes[tnum].NewConfidence
			if judgeChange(uncalculateObject) == 1 { //没改
				nodes[tnum].NewConfidence = HightenConfidence(t)
				fmt.Println(Db.Table("nodes").Where("table_id = ?", tnum).Update("old_confidence", nodes[tnum].OldConfidence))
				fmt.Println(Db.Table("nodes").Where("table_id = ?", tnum).Update("new_confidence", nodes[tnum].NewConfidence))
				//Db.M(&).Update("name", "hello")
			} else {
				nodes[tnum].NewConfidence = CalculateConfidence(uncalculateObject, t)
				fmt.Println(Db.Table("nodes").Where("table_id = ?", tnum).Update("old_confidence", nodes[tnum].OldConfidence))
				fmt.Println(Db.Table("nodes").Where("table_id = ?", tnum).Update("new_confidence", nodes[tnum].NewConfidence))
			}
		} else {
			node := NodesTable{0, object[i].ObjectPath, object[i].Parameters,
				object[i].CurrentObjectID, object[i].FatherObjectID, 0,
				CalculateConfidence(uncalculateObject, 0), object[i].CommitTableID,
				int(object[i].TableID), object[i].NewLine, object[i].OldLine,
				object[i].AddedLine, object[i].DeletedLine}
			fmt.Println(Db.Table("nodes").Create(&node).RowsAffected)
		}
	}
	context.Status(200)
}
