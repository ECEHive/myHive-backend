package service

import (
	"github.com/ECEHive/myHive-backend/db"
	"github.com/ECEHive/myHive-backend/entity"
	"github.com/ECEHive/myHive-backend/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var workbenchLogger = util.GetLogger("workbench-service")

func WorkbenchListAll(c *gin.Context) ([]*entity.Workbench, error) { // Context is here for logging only, and optional
	var logger = util.LocalLogger(workbenchLogger, c)
	conn := db.GetDB()
	var workbenches []*entity.Workbench

	if err := conn.Find(&workbenches).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			logger.Errorf("Error executing query %+v", err)
			return nil, err
		}
	}
	return workbenches, nil
}
