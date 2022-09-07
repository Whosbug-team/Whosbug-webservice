package routes

import (
	"github.com/gin-gonic/gin"
	. "webService_Refactoring/middlewear"
	. "webService_Refactoring/utils"
	. "webService_Refactoring/views"
)

// InitRouter 初始化
func InitRouter() {
	gin.SetMode(AppMode)
	r := gin.Default()
	r.POST("/v1/api-token-auth", CreateToken)

	api := r.Group("/v1/users")
	{
		api.POST("/", UserCreate)
		r.Use(CheckToken())
		api.GET("/:id", UserRead)
		api.PUT("/:id", UpdateUser)
		api.PATCH("/:id", UpdateUserPartial)
	}

	commits := r.Group("/v1/commits")
	{
		commits.POST("/commits-info", CommitsInfoCreate)       //1
		commits.POST("/delete_uncalculate", UnCalculateDelete) //1
		commits.POST("/diffs", CommitsDiffsCreate)             //1
		//review 暂时不重构
		commits.POST("/reviewers", CommitsReviewersCreate)
		commits.POST("/rules/", CommitsRulesCreate)
		//
		commits.POST("/train_method", CommitsTrainMethodCreate) //1
		commits.POST("/upload-done", CommitsUploadDoneCreate)   //1
	}
	r.POST("/v1/create-project-release", CreateProjectRelease) //1
	//r.POST("/v1/delete_all_related", AllRelatedDelete)         //1
	r.GET("/v1/liveness", LivenessList)         //1
	r.POST("/v1/owner", OwnerCreate)            //1
	r.POST("/v1/releases/last", GetLastRelease) //1
	r.Run(HTTPPort)
}
