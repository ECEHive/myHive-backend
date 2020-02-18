package controller

import (
	"bufio"
	"encoding/csv"
	"github.com/ECEHive/myHive-backend/db"
	"github.com/ECEHive/myHive-backend/entity"
	"github.com/ECEHive/myHive-backend/model"
	"github.com/ECEHive/myHive-backend/service"
	"github.com/ECEHive/myHive-backend/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func ConfigureInventoryRouter(r *gin.RouterGroup) {
	r.GET("/class/list", handlerInventoryClassList)
	r.PUT("/class/upsert", handlerInventoryClassUpsert)
	r.GET("/class/find", handlerInventoryClassFind)
	r.POST("/class/import", handlerInventoryClassImport)

	r.GET("/checkout/record/list", handlerInventoryCheckoutRecords)

	// Enum Types
	r.GET("/class/enum/count_types", handlerInventoryClassEnumCountTypes)
	r.GET("/class/enum/checkout_modes", handlerInventoryClassCheckoutModes)
}

func ConfigureV1InventoryRouter(r *gin.RouterGroup) {
	r.POST("/checkout/new", handlerInventoryCheckoutNew)
	r.GET("/checkout/items", handlerInventoryCheckoutItems)
}

var inventoryLogger = util.GetLogger("inventory-controller")

func handlerInventoryCheckoutNew(c *gin.Context) {
	var data model.InventoryCheckoutNewRequest
	if err := c.BindJSON(&data); err != nil {
		c.Set("error", model.BadRequest(util.EC_INVALID_REQUEST_BODY, err, "Malformed Request"+err.Error()))
		return
	}
	conn := db.GetDB()
	obj := entity.InventoryCheckoutRecord{
		Item:         data.Item,
		FirstName:    data.FirstName,
		LastName:     data.LastName,
		Email:        data.Email,
		CheckoutDate: entity.UnixTime(time.Now()),
		CheckoutPI:   data.CheckoutPI,
	}
	if err := conn.Save(&obj).Error; err != nil {
		c.Set("error", model.InternalServerError(util.EC_DB_ERROR, err, "something went wrong"))
		return
	}
	c.JSON(http.StatusOK, model.DataObject(obj))
}

func handlerInventoryCheckoutItems(c *gin.Context) {
	var values []*entity.InventoryCheckoutItem
	conn := db.GetDB()
	if err := conn.Find(&values).Error; err != nil {
		c.Set("error", model.InternalServerError(util.EC_DB_ERROR, err, "Something went wrong"))
		return
	}
	c.JSON(http.StatusOK, model.DataObject(values))
}

func handlerInventoryClassImport(c *gin.Context) {
	var filename string
	logger := util.LocalLogger(inventoryLogger, c)
	if file, err := c.FormFile("file"); err != nil {
		c.Set("error", model.BadRequest(util.EC_INVALID_REQUEST_BODY, err, "Missing form field: file"))
		return
	} else {
		if tempFile, err := ioutil.TempFile("", "hive_inventory_upload"); err != nil {
			c.Set("error", model.InternalServerError(util.EC_FS_ERROR, err, "failed to create tmp file"))
			return
		} else {
			filename = tempFile.Name()
			logger.Infof("Saving file to %s", filename)
			_ = tempFile.Close()
			_ = c.SaveUploadedFile(file, filename)
		}
	}
	csvFile, _ := os.Open(filename)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	skip := false
	var newRecords []entity.InventoryItemClass
	for {
		line, readerr := reader.Read()
		if readerr == io.EOF {
			break
		} else if readerr != nil {
			c.Set("error", model.BadRequest(util.EC_INVALID_REQUEST_BODY, readerr, "Bad file format"))
			return
		} else {
			if !skip {
				skip = true
				continue
			}
			if line[0] == "" {
				continue
			}
			label := line[1]
			if label == "" {
				label = line[0]
			}
			if len(label) > 20 {
				label = label[0:19]
			}
			countingType, _ := strconv.Atoi(line[2])
			stock, _ := strconv.Atoi(line[3])
			checkoutType, _ := strconv.Atoi(line[6])
			newRecords = append(newRecords, entity.InventoryItemClass{
				ItemName:         &line[0],
				ItemLabel:        &line[1],
				ItemCountingType: countingType,
				ItemCount:        int64(stock),
				ItemCountInStock: int64(stock),
				ItemDescription:  line[4],
				ItemDatasheet:    line[5],
				ItemCheckoutMode: checkoutType,
				ItemParameters:   line[7],
				ItemLocation:     line[8],
			})
		}
	}
	var patched []*entity.InventoryItemClass
	var errs []error
	for i := range newRecords {
		current := newRecords[i]
		patch, err := service.InventoryClassUpsert(&current, c)
		if err != nil {
			errs = append(errs, err)
		}
		patched = append(patched, patch)
	}
	c.JSON(http.StatusOK, model.DataObject(patched))
}

func handlerInventoryClassFind(c *gin.Context) {
	searchRequest := model.InventoryItemClassSearchRequest{}

	if err := c.BindJSON(&searchRequest); err != nil {
		c.Set("error", model.BadRequest(util.EC_INVALID_REQUEST_BODY, err, "Failed to bind json"))
		return
	}

}

func handlerInventoryClassCheckoutModes(c *gin.Context) {
	c.JSON(http.StatusOK, model.DataObject(entity.InventoryClassCheckoutModes))
}

func handlerInventoryClassEnumCountTypes(c *gin.Context) {
	c.JSON(http.StatusOK, model.DataObject(entity.InventoryClassCountingTypes))
}

func handlerInventoryClassUpsert(c *gin.Context) {
	patchModel := &entity.InventoryItemClass{}

	if err := c.BindJSON(patchModel); err != nil {
		c.Set("error", model.BadRequest(util.EC_INVALID_REQUEST_BODY, err,
			"Failed to bind request to json model"))
		return
	}

	if result, err := service.InventoryClassUpsert(patchModel, c); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.Set("error", model.BadRequest(util.EC_NOT_FOUND, err, "Can not patch item that does not exist"))
		} else {
			c.Set("error", model.InternalServerError(util.EC_DB_ERROR, err, "Error while saving changes"))
		}
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

func handlerInventoryCheckoutRecords(c *gin.Context) {
	var all []*entity.InventoryCheckoutRecord
	if e := db.GetDB().Select(&all).Error; e != nil {
		c.Set("error", model.InternalServerError(util.EC_DB_ERROR, e, "Internal DB Error while Querying"))
		return
	}
	c.JSON(http.StatusOK, model.DataObject(all))

}
