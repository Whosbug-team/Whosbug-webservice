package routes

import (
	_ "webService_Refactoring/middlewear"
	. "webService_Refactoring/utils"
	. "webService_Refactoring/views"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// InitRouter 初始化
func InitRouter() {
	gin.SetMode(AppMode)
	r := gin.Default()
	r.POST("/v1/api-token-auth", CreateToken)

	api := r.Group("/v1/users")
	{
		//r.Use(CheckToken())
		api.POST("/", UserCreate)
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

	//prometheus
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.Run(HTTPPort)
}
