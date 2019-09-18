package controller

import "github.com/gin-gonic/gin"

func workbenchHandlerList(ctx *gin.Context) {

}

func ConfigureWorkbenchRoutes(r *gin.RouterGroup) {
	r.GET("/list")
}
