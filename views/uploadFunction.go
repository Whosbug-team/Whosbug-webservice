package views

import (
	"fmt"
	. "webService_Refactoring/modules"
)

func judgeObject(temp2 ObjectsTable, nodes []NodesTable) (int, int) {

	// temp2 := ObjectsTable{}
	// db.Table("objects").Model(&ObjectsTable{}).Find(&temp2)
	//var nodes1 []NodesTable
	//第一次筛选
	var methods []NodesTable
	i := 0
	var nodesnum, pathsnum, paramsnum []int
	var tnum int
	for x := 0; x < len(nodes); x++ { //0-9
		if nodes[x].CurrentObjectID == temp2.CurrentObjectID {
			methods, nodesnum = append(methods, nodes[x]), append(nodesnum, x)
		} //0 1 2 3
	}
	if len(methods) == 0 {
		fmt.Println("Get1 objects error:")
		return 0, 0
	}
	//第二次筛选
	var path []NodesTable
	for x := 0; x < len(methods); x++ {
		if methods[x].ObjectPath == temp2.ObjectPath {
			path, pathsnum = append(path, methods[x]), append(pathsnum, x)
		} //0 1
	}
	if len(path) == 0 {
		fmt.Println("Get2 objects error:")
		return 0, 0

	}
	//第三次筛选
	var params []NodesTable
	for x := 0; x < len(path); x++ {
		if path[x].ObjectParameters == temp2.Parameters {
			params, paramsnum, i = append(params, path[x]), append(paramsnum, x), x
		}
	}
	if len(params) == 0 {
		fmt.Println("Get3 objects error:")
		return 0, 0
	} else {
		tnum = nodesnum[pathsnum[i]]
	}

	return i, tnum

}

func judgeChange(object UncalculateObjectInfo) int {
	if object.addedLineCount == 0 && object.deletedlineCount == 0 {
		return 1
	} else {
		return 0
	}

}
