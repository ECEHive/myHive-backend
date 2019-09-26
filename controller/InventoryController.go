package controller

import (
	"github.com/ECEHive/myHive-backend/entity"
	"github.com/ECEHive/myHive-backend/model"
	"github.com/ECEHive/myHive-backend/service"
	"github.com/ECEHive/myHive-backend/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ConfigureInventoryRouter(r *gin.RouterGroup) {
	r.GET("/class/list", handlerInventoryClassList)
	r.PUT("/class/upsert", handlerInventoryClassUpsert)
}

func handlerInventoryClassUpsert(c *gin.Context) {
	patchModel := &entity.InventoryItemClass{}

	if err := c.BindJSON(patchModel); err != nil {
		c.Set("error", model.BadRequest(util.EC_INVALID_REQUEST_BODY, err,
			"Failed to bind request to json model"))
		return
	}

	if result, err := service.InventoryClassUpsert(patchModel, c); err != nil {
		c.Set("error", model.InternalServerError(util.EC_DB_ERROR, err, "Error while saving changes"))
		return
	} else {
		c.JSON(http.StatusOK, model.DataObject(result))
	}
}

func handlerInventoryClassList(c *gin.Context) {
	paginationRequest := c.MustGet("pagination").(model.PaginationRequest) // Get pagination
	if result, pagination, err := service.InventoryItemClassList(&paginationRequest, c); err != nil {
		c.Set("error", model.InternalServerError(util.EC_DB_ERROR, err, "Something went wrong while querying"))
		return
	} else {
		c.JSON(http.StatusOK, model.DataObject(result, pagination))
	}
}
