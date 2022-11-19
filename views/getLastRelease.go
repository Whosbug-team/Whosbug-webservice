package views

import (
	"errors"
	"net/http"
	. "webService_Refactoring/modules"
	"webService_Refactoring/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetLastRelease 从数据库中获取最新的版本信息
func GetLastRelease(c *gin.Context) {
	var id ProjectID
	if err := c.ShouldBind(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": utils.GetErrMsg(http.StatusBadRequest),
		})
		return
	}
	projectID := ProjectsTable{}
	Db.Table("projects").Where("project_id = ?", id.Pid).First(&projectID)
	projectTableID, temp := projectID.TableID, ReleasesTable{}
	if err := Db.Table("releases").Where("project_table_id = ?", projectTableID).First(&temp).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": utils.GetErrMsg(http.StatusNotFound),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"msg": utils.GetErrMsg(http.StatusCreated),
		"lastReleaseVersion": temp.ReleaseVersion,
	})

}
