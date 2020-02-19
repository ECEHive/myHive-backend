package model

import (
	"github.com/ECEHive/myHive-backend/constants"
)

type InventoryItemClassSearchRequest struct {
	SearchKeyword string
	NameOnly      bool
	LabelOnly     bool
	InStock       bool
}

type InventoryCheckoutNewRequest struct {
	Item      string `binding:"required"`
	FirstName string `binding:"required"`
	LastName  string `binding:"required"`
	Email     string `binding:"required"`

	CheckoutPI string `binding:"required"`
}

type InventoryCheckoutUpdateRequest struct {
	Id         int                               `binding:"required"`
	NewStatus  constants.InventoryCheckoutStatus `binding:"required"`
	CheckoutPI string
}
