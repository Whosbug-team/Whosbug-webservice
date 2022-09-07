package views

import (
	"errors"
	"fmt"
	"net/http"
	. "webService_Refactoring/modules"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CommitsUploadDoneCreate 上传完数据后计算置信度，然后插入数据库
func CommitsUploadDoneCreate(context *gin.Context) {

	var t T

	err := context.ShouldBind(&t)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(t)
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

	//temp2 := ObjectsTable{}
	var temp2 []ObjectsTable
	Db.Table("objects").Find(&temp2)
	n := len(temp2)

	for i := 0; i < n; i++ {
		var nodes []NodesTable
		Db.Table("nodes").Find(&nodes)
		temp3 := UncalculateObjectInfo{temp2[i].Hash, temp2[i].CurrentObjectID, temp2[i].FatherObjectID,
			temp2[i].Parameters, temp2[i].StartLine, temp2[i].EndLine, temp2[i].OldLine,
			temp2[i].NewLine, temp2[i].DeletedLine, temp2[i].AddedLine}
		//var nodes1 []NodesTable
		var tnum int

		num, tnum := judgeObject(temp2[i], nodes)
		fmt.Println("n:", tnum)
		if num != 0 { //有object
			t := nodes[tnum].OldConfidence
			nodes[tnum].OldConfidence = nodes[tnum].NewConfidence
			if judgeChange(temp3) == 1 { //没改
				nodes[tnum].NewConfidence = HightenConfidence(t)
				fmt.Println(Db.Table("nodes").Where("table_id = ?", tnum).Update("old_confidence", nodes[tnum].OldConfidence))
				fmt.Println(Db.Table("nodes").Where("table_id = ?", tnum).Update("new_confidence", nodes[tnum].NewConfidence))
				//Db.M(&).Update("name", "hello")
			} else {
				nodes[tnum].NewConfidence = CalculateConfidence(temp3, t)
				fmt.Println(Db.Table("nodes").Where("table_id = ?", tnum).Update("old_confidence", nodes[tnum].OldConfidence))
				fmt.Println(Db.Table("nodes").Where("table_id = ?", tnum).Update("new_confidence", nodes[tnum].NewConfidence))
			}
		} else {
			temp4 := NodesTable{0, temp2[i].ObjectPath, temp2[i].Parameters,
				temp2[i].CurrentObjectID, temp2[i].FatherObjectID, 0,
				CalculateConfidence(temp3, 0), temp2[i].CommitTableID,
				int(temp2[i].TableID), temp2[i].NewLine, temp2[i].OldLine,
				temp2[i].AddedLine, temp2[i].DeletedLine}
			fmt.Println(Db.Table("nodes").Create(&temp4).RowsAffected)
		}
	}
	context.Status(200)
}
