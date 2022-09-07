package views

import (
	"errors"
	"gorm.io/gorm"
	. "webService_Refactoring/modules"
)

// CommitInfo commit信息
type CommitInfo struct {
	commitHash   string
	commitAuthor string
	commitEmail  string
	commitTime   string
}

// ObjectInfo 已经计算置信度的object信息
type ObjectInfo struct {
	objectID         string  `json:"object_id"`         //Object的函数唯一标识符
	oldObjectID      string  `json:"old_object_id"`     //Object的父类唯一表示符
	confidence       float64 `json:"confidence"`        //置信度
	parameters       string  `json:"parameters"`        //方法的参数特征
	oldlineCount     int     `json:"oldline_count"`     //旧行数
	newlineCount     int     `json:"newline_count"`     //新行数
	deletedlineCount int     `json:"deletedline_count"` //移除行数
	addedLineCount   int     `json:"added_line_count"`  //新增行数
}

// UncalculateObjectInfo 未计算置信度的object信息
type UncalculateObjectInfo struct {
	hash             string //Object所属的Commit
	objectID         string //Object的函数唯一标识符
	oldObjectID      string //Object的父类唯一表示符
	parameters       string //方法的参数特征
	startLine        int    //起始行
	endLine          int    //结束行
	oldlineCount     int    //旧行数
	newlineCount     int    //新行数
	deletedlineCount int    //移除行数
	addedLineCount   int    //新增行数
}

// OwnerInfo 责任人信息
type OwnerInfo struct {
	author string  //责任人名称
	email  string  //邮箱
	weight float64 //权重
}

// ObjectHistoryInfo object的历史变更比例
type ObjectHistoryInfo struct {
	oldlineCount     int //旧行数
	newlineCount     int //新行数
	deletedlineCount int //移除行数
	addedLineCount   int //新增行数
}

// HistoryInfo object的历史修改记录
type HistoryInfo struct {
	commitHistory CommitInfo
	objectHistory ObjectHistoryInfo
}

type bugOriginInfo struct {
	object    ObjectInfo         `json:"object"`
	wrongRate float64            `json:"wrong_rate"`
	owners    map[string]float64 `json:"owners"`
}

// TreeNode 定义链的根节点
type TreeNode struct {
	object ObjectInfo
	childs []TreeNode
}

//  @param objectId
//  @return []historyInfo
//  返回的切片要按时间顺序排，最新的commit及其对应object放在索引0
func getHistory(objectID string) (result []HistoryInfo) {
	var temp []NodesTable
	res2 := Db.Table("nodes").Where("current_object_id = ? ", objectID).Find(&temp)
	if errors.Is(res2.Error, gorm.ErrRecordNotFound) {
		return
	}
	n := len(temp)
	for i := 0; i < n; i++ {
		var temp1 CommitsTable
		tableID := temp[i].CommitTableID
		res := Db.Table("commits").Where("table_id = ? ", tableID).First(&temp1)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return
		}
		temp3 := ObjectHistoryInfo{temp[i].ObjectOldLine, temp[i].ObjectNewLine,
			temp[i].ObjectDeleteLine, temp[i].ObjectAddLine}
		temp4 := CommitInfo{temp1.Hash, temp1.Author, temp1.Email, temp1.Time}
		temp5 := HistoryInfo{temp4, temp3}
		result = append(result, temp5)
	}

	return result

}

//  @param objectId 函数的id
//  @return	chainNode 该函数所在的定义链的根结点
func getChain(objectID string) (node TreeNode) {
	temp := ObjectsTable{}
	Db.Table("objects").First(&temp, "current_object_id = ?", objectID)
	node.object = ObjectInfo{temp.CurrentObjectID, temp.FatherObjectID, 0,
		temp.Parameters, temp.OldLine, temp.NewLine,
		temp.DeletedLine, temp.AddedLine}
	var tempChilds []ObjectsTable
	Db.Table("objects").Find(&tempChilds, "father_object_id in (?)", objectID)
	for i := range tempChilds {
		node.childs = append(node.childs, getChain(tempChilds[i].CurrentObjectID))
	}
	return
}
