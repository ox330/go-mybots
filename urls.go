package go_mybots

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hand() *gin.Engine {
	rout := gin.New()
	gin.SetMode(gin.ReleaseMode)
	rout.POST("/commit/", func(context *gin.Context) {
		eventMain(context.Request.Body)
		context.JSON(http.StatusOK, gin.H{})
	})
	return rout
}
