package model

import (
	"github.com/ECEHive/myHive-backend/entity"
	"github.com/ECEHive/myHive-backend/util"
	"net/http"
	"time"
)

type StandardObject struct {
	Code string          `json:"code"`
	Time entity.UnixTime `json:"time"`
}

type ErrorResponse struct {
	StandardObject
	Error      error  `json:"-"`
	Reason     string `json:"reason"`
	StatusCode int    `json:"-"`
}

type DataResponse struct {
	StandardObject
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination"`
}

func currentTime() entity.UnixTime {
	return entity.UnixTime(time.Now().UTC())
}

func DataObject(data interface{}, pagination ...interface{}) interface{} {
	obj := DataResponse{}
	obj.Time = currentTime()
	obj.Code = "SUCCESS"
	obj.Data = data
	if len(pagination) > 0 {
		obj.Pagination = pagination[0]
	} else {
		obj.Pagination = nil
	}
	return obj
}

type PaginationInformation struct {
	TotalItems  int64 `json:"totalItems"`
	CurrentPage int64 `json:"currentPage"`
	TotalPages  int64 `json:"totalPages"`
	PageSize    int64 `json:"pageSize"`
}

func ComputePaginationInformation(page int64, pageSize int64, totalItemCount int64) (info *PaginationInformation) {
	if pageSize < 1 {
		return nil
	}
	info = &PaginationInformation{}
	info.PageSize = pageSize
	info.TotalPages = totalItemCount/pageSize + 1
	info.TotalItems = totalItemCount
	if info.TotalPages <= page {
		info.CurrentPage = info.TotalPages - 1
	} else {
		info.CurrentPage = page
	}
	return
}

type PaginationRequest struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"validate:"required"binding:"required"`
}

var (
	newErrorResponse = func(code string, err error, reason string, statusCode int) ErrorResponse {
		response := ErrorResponse{
			StatusCode: statusCode,
			Reason:     reason,
			Error:      err,
		}
		response.Code = code
		response.Time = entity.UnixTime(time.Now().UTC())
		return response
	}
	InternalServerError = func(code string, err error, reason string) ErrorResponse {
		return newErrorResponse(code, err, reason, http.StatusInternalServerError)
	}
	BadRequest = func(code string, err error, reason string) ErrorResponse {
		return newErrorResponse(code, err, reason, http.StatusBadRequest)
	}
	NotFound = func(reason string) ErrorResponse {
		return newErrorResponse(util.EC_NOT_FOUND, nil, reason, http.StatusNotFound)
	}
)
