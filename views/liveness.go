package views

import "github.com/gin-gonic/gin"

// LivenessList secure the liveness of the webservice
func LivenessList(context *gin.Context) {
	context.Status(200)
}
