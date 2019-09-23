package service

import (
	"github.com/ECEHive/myHive-backend/db"
	"github.com/ECEHive/myHive-backend/entity"
	"github.com/ECEHive/myHive-backend/model"
	"github.com/ECEHive/myHive-backend/util"
	"github.com/gin-gonic/gin"
)

var userServiceLogger = util.GetLogger("user_service")

func HiveUserQueryWithPaginationOptions(userQuery *entity.HiveUser, page *model.PaginationRequest) (result []*entity.HiveUser, pageInfo *model.PaginationInformation) {
	result = []*entity.HiveUser{}
	query := db.GetDB().Model(&entity.HiveUser{}).Where(userQuery)
	var count int64 = 0
	query.Count(&count)
	pageInfo = model.ComputePaginationInformation(page.Page, page.PageSize, count)
	query.Limit(pageInfo.PageSize).
		Order("last_name asc").
		Offset(pageInfo.CurrentPage * pageInfo.PageSize).
		Find(&result)
	if err := query.Error; err != nil {
		userServiceLogger.Errorf("SQL Error: %+v", err)
		return nil, nil
	}
	return
}

func UpdateModel(model *entity.HiveUser, patch *entity.HiveUser, ctx *gin.Context) {
	var logger = util.LocalLogger(userServiceLogger, ctx)
	conn := db.GetDB()

	// Remove unexpected fields from the patch
	if patch.Id != 0 || patch.UniqueIdentifier != "" || patch.CreatedAt != nil || patch.UpdatedAt != nil {
		logger.Warnf("Patch with restricted fields detected: %+v", patch)
		patch.Id = 0
		patch.UniqueIdentifier = ""
		patch.CreatedAt = nil
		patch.UpdatedAt = nil
	}

	if err := conn.Model(model).Update(patch).Error; err != nil {
		logger.Errorf("DB Error: %+v", err)
	}
}
