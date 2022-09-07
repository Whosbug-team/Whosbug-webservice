package views

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	. "webService_Refactoring/modules"
)

// OwnerCreate 接收堆栈信息后，剔除系统库和第三方库函数后，通过keyman算法返回主要责任人
func OwnerCreate(context *gin.Context) {
	var t GetConfidence
	if err := context.ShouldBind(&t); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	pid, version, temp := t.Project.Pid, t.Release.Version, ProjectsTable{}
	res := Db.Table("projects").First(&temp, "project_id = ?", pid)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Project pid " + t.Project.Pid + " not exists",
		})
		return
	}
	temp1 := ReleasesTable{}
	res1 := Db.Table("releases").First(&temp1, "release_version = ?", version)
	if errors.Is(res1.Error, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Release version " + t.Release.Version + " not exists",
		})
		return
	}
	releaseTableID, methods := temp1.TableID, t.Method
	n, jsonResult, params, objectInfos := len(methods), JSONRes{}, make([]NodesTable, 0), make([]ObjectInfo, 0)
	for i := 0; i < n; i++ {
		methodID, filePath, parameters, commitDemo, path :=
			methods[i].MethodID, methods[i].Filepath, methods[i].Parameters, CommitsTable{}, make([]NodesTable, 0)
		Db.Table("commits").First(&commitDemo, "release_table_id = ?", releaseTableID)
		commitTableID, nodes, methods2 := commitDemo.TableID, make([]NodesTable, 0), make([]NodesTable, 0)
		//数据库中查找所有符合条件的数据
		Db.Table("nodes").Find(&nodes, "commit_table_id in (?)", commitTableID)
		if len(nodes) == 0 {
			context.Status(400)
			continue
		}
		//第一次筛选
		for x := 0; x < len(nodes); x++ {
			if nodes[x].CurrentObjectID == methodID {
				methods2 = append(methods2, nodes[x])
			}
		}
		if len(methods2) == 0 {
			jsonResult.Message, jsonResult.Status = "No such objects in release: version: "+version, "may be ok"
			fmt.Println("Get objects error:", jsonResult)
			continue
		}
		//第二次筛选
		for x := 0; x < len(methods2); x++ {
			if methods2[x].ObjectPath == filePath {
				path = append(path, methods2[x])
			}
		}
		if len(path) == 0 {
			jsonResult.Message, jsonResult.Object, jsonResult.Status =
				"No such objects in path: filepath: "+filePath+" here's results with id", nodes, "may be ok"
			fmt.Println("Get objects error:", jsonResult)
			continue
		}
		//第三次筛选
		for x := 0; x < len(path); x++ {
			if path[x].ObjectParameters == parameters {
				params = append(params, path[x])
			}
		}
		if len(params) == 0 {
			jsonResult.Message, jsonResult.Object, jsonResult.Status =
				"No such objects in params: "+parameters+" here's results with path", path, "may be ok"
			fmt.Println("Get objects error:", jsonResult)
			continue
		}
	}
	for i := 0; i < len(params); i++ {
		objectInfo := ObjectInfo{params[i].CurrentObjectID, params[i].FatherObjectID,
			params[i].NewConfidence, params[i].ObjectParameters, params[i].ObjectOldLine,
			params[i].ObjectNewLine, params[i].ObjectDeleteLine, params[i].ObjectAddLine}
		objectInfos = append(objectInfos, objectInfo)
	}
	tt := returnOwner(GetBugOrigin(objectInfos))
	context.JSON(http.StatusOK, gin.H{
		"owner": tt,
	})
}

type testss struct {
	Name   string  `json:"name" binding:"required"`
	Weight float64 `json:"weight" binding:"required"`
}

func returnOwner(OriginInfo map[string]float64) []testss {
	knum, vnum := make([]string, 0), make([]float64, 0)
	for k, v := range OriginInfo {
		knum, vnum = append(knum, k), append(vnum, v)
	}
	tt, tt1 := make([]testss, 0), testss{}
	for i := 0; i < len(OriginInfo); i++ {
		tt1.Name, tt1.Weight = knum[i], vnum[i]
		tt = append(tt, tt1)
	}
	return tt
}
