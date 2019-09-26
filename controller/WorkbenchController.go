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
	r.GET("/list", handlerWorkbenchList)
	r.POST("/create", handlerWorkbenchCreate)

	r.GET("/enum/record_types", handlerWorkbenchEnumRecordTypes)
}

func handlerWorkbenchEnumRecordTypes(r *gin.Context) {
	r.JSON(http.StatusOK, model.DataObject(entity.WorkbenchRecordTypes))
}

func handlerWorkbenchList(ctx *gin.Context) {
	workbenches, err := service.WorkbenchListAll(ctx)
	if err != nil {
		ctx.Set("error", model.InternalServerError(util.EC_DB_ERROR, err, "Something went wrong."))
		return
	}
	ctx.JSON(http.StatusOK, model.DataObject(workbenches))
}

func handlerWorkbenchCreate(ctx *gin.Context) {
	benchCreateRequest := &model.WorkbenchCreationRequest{}

	if err := ctx.BindJSON(benchCreateRequest); err != nil {
		ctx.Set("error", model.BadRequest(util.EC_INVALID_REQUEST_BODY, err, "Can not parse request body"))
		return
	}

	if bench, err := service.WorkbenchCreate(benchCreateRequest, ctx); err == nil {
		ctx.JSON(http.StatusOK, model.DataObject(bench))
		return
	} else {
		ctx.Set("error", model.InternalServerError(util.EC_DB_ERROR, err, "Internal Database Error"))
		return
	}
}
