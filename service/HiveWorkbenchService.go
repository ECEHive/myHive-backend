package service

import (
	"errors"
	"github.com/ECEHive/myHive-backend/db"
	"github.com/ECEHive/myHive-backend/entity"
	"github.com/ECEHive/myHive-backend/model"
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

func WorkbenchCreate(request *model.WorkbenchCreationRequest, c *gin.Context) (*entity.Workbench, error) {
	var logger = util.LocalLogger(workbenchLogger, c)

	if request.Name == "" {
		return nil, errors.New("Workbench needs to have a non empty name")
	}
	conn := db.GetDB()

	bench := &entity.Workbench{
		BenchName: request.Name,
	}

	if err := conn.Save(bench).Error; err != nil {
		logger.Errorf("Error while saving new workbench: %+v", err)
		return nil, err
	}

	return bench, nil
}
