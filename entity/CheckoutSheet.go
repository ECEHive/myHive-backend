package entity

type InventoryCheckoutRecord struct {
	BaseModel
	FirstName string
	LastName  string
	Email     string

	Item         string
	CheckoutDate UnixTime `gorm:"type:timestamp"`
	Returned     bool
	CheckoutPI   string

	LastEmail *UnixTime
}

type InventoryCheckoutItem struct {
	Id       int `gorm:"primary_key"`
	ItemName string
}
