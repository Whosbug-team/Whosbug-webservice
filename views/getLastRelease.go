package views

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	. "webService_Refactoring/modules"
)

// GetLastRelease 从数据库中获取最新的版本信息
func GetLastRelease(c *gin.Context) {
	var id ProjectID
	if err := c.ShouldBind(&id); err != nil {
		err.Error()
	}
	projectID := ProjectsTable{}
	Db.Table("projects").Where("project_id = ?", id.Pid).First(&projectID)
	projectTableID, temp := projectID.TableID, ReleasesTable{}
	res := Db.Table("releases").Where("project_table_id = ?", projectTableID).First(&temp)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		info := "The project with pid = " + id.Pid + " does not exists."
		c.JSON(http.StatusNotFound, gin.H{
			"error": info,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"lastReleaseVersion": temp.ReleaseVersion,
	})

}
