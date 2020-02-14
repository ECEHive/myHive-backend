package entity

type InventoryCheckoutRecord struct {
	BaseModel
	SheetRow uint `gorm:"unique"`
	Name     string
	Email    string

	Item         string
	CheckoutDate UnixTime `gorm:"type:timestamp"`
	Returned     bool
	CheckoutPI   string

	LastEmail *UnixTime
}
