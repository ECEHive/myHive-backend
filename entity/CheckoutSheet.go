package entity

import "github.com/ECEHive/myHive-backend/constants"

type InventoryCheckoutRecord struct {
	BaseModel
	FirstName string
	LastName  string
	Email     string

	Item         int64
	CheckoutDate UnixTime `gorm:"type:timestamp"`
	Status       constants.InventoryCheckoutStatus
	CheckoutPI   string

	LastEmail *UnixTime
}

type InventoryCheckoutItem struct {
	Id       int `gorm:"primary_key"`
	ItemName string
}
