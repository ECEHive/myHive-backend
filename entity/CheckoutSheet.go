package entity

type InventoryCheckoutStatus string

const (
	InventoryCheckoutStatusCheckedOut   InventoryCheckoutStatus = "CheckedOut"
	InventoryCheckoutStatusExtended     InventoryCheckoutStatus = "Extended"
	InventoryCheckoutStatusReturned     InventoryCheckoutStatus = "Returned"
	InventoryCheckoutStatusLostOrDamage InventoryCheckoutStatus = "LostOrDamage"
)

type InventoryCheckoutRecord struct {
	BaseModel
	FirstName string
	LastName  string
	Email     string

	Item         string
	CheckoutDate UnixTime `gorm:"type:timestamp"`
	Status       InventoryCheckoutStatus
	CheckoutPI   string

	LastEmail *UnixTime
}

type InventoryCheckoutItem struct {
	Id       int `gorm:"primary_key"`
	ItemName string
}
