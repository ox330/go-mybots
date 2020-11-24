package urls

import (
	"github.com/3343780376/go-mybots"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hand() *gin.Engine {
	rout := gin.New()
	gin.SetMode(gin.ReleaseMode)
	rout.POST("/commit/", func(context *gin.Context) {
		go_mybots.EventMain(context.Request.Body)
		context.JSON(http.StatusOK,gin.H{
			})
	})
	return rout
}