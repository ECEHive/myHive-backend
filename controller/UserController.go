package controller

import (
	"github.com/ECEHive/myHive-backend/entity"
	"github.com/ECEHive/myHive-backend/model"
	"github.com/ECEHive/myHive-backend/service"
	"github.com/ECEHive/myHive-backend/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ConfigureUserController(r *gin.RouterGroup) {
	r.POST("/find", handlerUserLookup)
}

func handlerUserLookup(c *gin.Context) {
	request := &entity.HiveUser{}

	if err := c.BindJSON(request); err != nil {
		c.Set("error", model.BadRequest(util.EC_INVALID_REQUEST_BODY, err, err.Error()))
		return
	}

	paginationRequest := c.MustGet("pagination").(model.PaginationRequest) // Get pagination

	users, pagination := service.HiveUserQueryWithPaginationOptions(request, &paginationRequest)

	c.JSON(http.StatusOK, model.DataObject(users, pagination))
}
