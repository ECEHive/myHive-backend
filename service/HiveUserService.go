package service

import (
	"github.com/ECEHive/myHive-backend/db"
	"github.com/ECEHive/myHive-backend/entity"
	"github.com/ECEHive/myHive-backend/model"
	"github.com/ECEHive/myHive-backend/util"
)

var userServiceLogger = util.GetLogger("user_service")

func HiveUserQueryWithPaginationOptions(userQuery *entity.HiveUser, page *model.PaginationRequest) (result []*entity.HiveUser, pageInfo *model.PaginationInformation) {
	result = []*entity.HiveUser{}
	query := db.GetDB().Where(userQuery)
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
