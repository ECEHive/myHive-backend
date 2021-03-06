package controller

import (
	"github.com/ECEHive/myHive-backend/db"
	"github.com/ECEHive/myHive-backend/entity"
	"github.com/ECEHive/myHive-backend/model"
	"github.com/ECEHive/myHive-backend/service"
	"github.com/ECEHive/myHive-backend/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ConfigureUserController(r *gin.RouterGroup) {
	r.POST("/find", handlerUserLookup)
	r.PUT("/upsert", handlerUserUpsert)
	r.GET("/enum/types", handlerUserEnumTypes)
}

var userControllerLogger = util.GetLogger("user_controller")

func handlerUserEnumTypes(c *gin.Context) {
	c.JSON(http.StatusOK, model.DataObject(entity.HiveUserTypes))
}

func handlerUserLookup(c *gin.Context) {
	request := &entity.HiveUser{}

	if err := c.BindJSON(request); err != nil {
		c.Set("error", model.BadRequest(util.EC_INVALID_REQUEST_BODY, err, err.Error()))
		return
	}

	paginationRequest := c.MustGet("pagination").(model.PaginationRequest) // Get pagination

	users, pagination := service.HiveUserQueryWithPaginationOptions(request, &paginationRequest)
	if users != nil && pagination != nil {
		c.JSON(http.StatusOK, model.DataObject(users, pagination))
		return
	}
	c.AbortWithStatus(http.StatusInternalServerError)
}

func handlerUserUpsert(c *gin.Context) {
	var logger = util.LocalLogger(userControllerLogger, c)
	request := &model.HiveUserUpsertRequest{}
	if err := c.BindJSON(request); err != nil {
		c.Set("error", model.BadRequest(util.EC_INVALID_REQUEST_BODY, err, err.Error()))
		return
	}
	if request.UniqueIdentifier == "" {
		c.Set("error", model.BadRequest(util.EC_INVALID_REQUEST_BODY, nil,
			"UniqueIdentifier should be non-empty"))
		return
	}
	conn := db.GetDB()
	currentModel := &entity.HiveUser{}
	if err := conn.Where("unique_identifier = ?", request.UniqueIdentifier).
		Find(currentModel).Error; err != nil {
		logger.Infof("Upsert user is creating user with identification %s", request.UniqueIdentifier)
		currentModel.UniqueIdentifier = request.UniqueIdentifier
		if err := conn.Save(currentModel).Error; err != nil {
			logger.Error("Error while inserting record %+v", err)
			c.Set("error", model.InternalServerError(util.EC_DB_ERROR, err, "unknown error"))
			return
		}
	}
	service.UpdateModel(currentModel, request.Patch, c)
	c.JSON(http.StatusOK, model.DataObject(currentModel, nil))
}
