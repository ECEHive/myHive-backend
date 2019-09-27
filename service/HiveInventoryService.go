package service

import (
	"fmt"
	"github.com/ECEHive/myHive-backend/db"
	"github.com/ECEHive/myHive-backend/entity"
	"github.com/ECEHive/myHive-backend/model"
	"github.com/ECEHive/myHive-backend/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var inventoryServiceLogger = util.GetLogger("inventory-service")

func InventoryItemClassFind(findRequest *model.InventoryItemClassSearchRequest) {
}

func InventoryItemClassList(paginationRequest *model.PaginationRequest, ctx *gin.Context) ([]*entity.InventoryItemClass, *model.PaginationInformation, error) {
	var logger = util.LocalLogger(inventoryServiceLogger, ctx)

	conn := db.GetDB()

	var count int64
	var queryResult []*entity.InventoryItemClass

	if err := conn.Model(&entity.InventoryItemClass{}).Count(&count).Error; err != nil {
		logger.Errorf("Error while counting %+v", err)
		return nil, nil, err
	}
	pagination := model.ComputePaginationInformation(paginationRequest.Page, paginationRequest.PageSize, count)

	if err := conn.Find(&queryResult).
		Limit(pagination.PageSize).
		Order("id ASC").
		Offset(pagination.CurrentPage).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			logger.Errorf("Error while querying DB: %+v", err)
			return nil, nil, err
		}
	}
	return queryResult, pagination, nil
}

func InventoryClassUpsert(patchSet *entity.InventoryItemClass, c *gin.Context) (*entity.InventoryItemClass, error) {
	var logger = util.LocalLogger(inventoryServiceLogger, c)

	patchId := patchSet.Id
	conn := db.GetDB()
	finalModel := entity.InventoryItemClass{}
	if patchId != 0 {
		if err := conn.Where("id = ?", patchId).Find(&finalModel).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				logger.Warnf("Unable to find HiveInventoryItemClass with id = %d", patchId)
				return nil, err
			} else {
				logger.Errorf("DB Error: %+v", err)
				return nil, err
			}
		}
		patchSet.Id = 0
	} else {
		// Create Model
		nextSequence := SequenceIncrementAndGet(util.SequenceInventoryItemLabelId)
		if nextSequence <= 0 {
			logger.Panicf("Sequence(%s) encountered negative value: %d", util.SequenceInventoryItemLabelId,
				nextSequence)
		}
		id := fmt.Sprintf("ICLS%04d", nextSequence)
		patchSet.ItemLabelID = id
		if err := conn.Save(*patchSet).Error; err != nil {
			logger.Errorf("Error Saving: %+v", err)
			return nil, err
		}
		return patchSet, nil
	}

	// Filter not patchable section of the model
	patchSet.CreatedAt = nil
	patchSet.UpdatedAt = nil
	patchSet.ItemLabelID = ""

	if err := conn.Model(finalModel).Update(patchSet).Error; err != nil {
		logger.Errorf("Error while saving: %+v", err)
		return nil, err
	}

	if err := conn.Where("id = ?", finalModel.Id).Find(&finalModel).Error; err != nil {
		logger.Errorf("Error while requerying: %+v", err)
		return nil, err
	}

	return &finalModel, nil
}
