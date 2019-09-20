package middleware

import (
	"github.com/ECEHive/myHive-backend/model"
	"github.com/ECEHive/myHive-backend/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

var pgLogger = util.GetLogger("middlware", "pagination")

func PaginationResolver(ctx *gin.Context) {
	pagination := model.PaginationRequest{}
	pageParam := ctx.Query("page")
	page, err := strconv.ParseInt(pageParam, 10, 64)
	if err != nil {
		page = 0
	}
	pagination.Page = page

	pageSizeParam := ctx.Query("page_size")
	pageSize, err := strconv.ParseInt(pageSizeParam, 10, 64)
	if err != nil {
		pageSize = 20
	}
	pagination.PageSize = pageSize
	pgLogger.Infof("pagination: page: %s, page_size: %s", pageParam, pageSizeParam)
	ctx.Set("pagination", pagination)
	ctx.Next()
}
