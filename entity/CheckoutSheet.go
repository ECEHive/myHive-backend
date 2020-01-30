package entity

type InventoryCheckoutRecord struct {
	BaseModel
	Name  string
	Email string
	Phone *string

	Item               string
	CheckoutDate       UnixTime
	ExpectedReturnData UnixTime
	Returned           bool
	CheckoutPI         string

	LastEmail *UnixTime
}
