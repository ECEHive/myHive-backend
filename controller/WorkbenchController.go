package controller

import (
	"github.com/ECEHive/myHive-backend/entity"
	"github.com/ECEHive/myHive-backend/model"
	"github.com/ECEHive/myHive-backend/service"
	"github.com/ECEHive/myHive-backend/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ConfigureWorkbenchRoutes(r *gin.RouterGroup) {
	r.GET("/list", workbenchHandlerList)

	r.GET("/enum/record_types", handlerWorkbenchEnumRecordTypes)
}

func handlerWorkbenchEnumRecordTypes(r *gin.Context) {
	r.JSON(http.StatusOK, model.DataObject(entity.WorkbenchRecordTypes))
}

func workbenchHandlerList(ctx *gin.Context) {
	workbenches, err := service.WorkbenchListAll(ctx)
	if err != nil {
		ctx.Set("error", model.InternalServerError(util.EC_DB_ERROR, err, "Something went wrong."))
		return
	}
	ctx.JSON(http.StatusOK, model.DataObject(workbenches))
}
