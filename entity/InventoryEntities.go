package entity

const (
	CountingType_Estimate = iota
	CountintType_Exact
	CountingType_Instance
)

var InventoryClassCountingTypes = map[int]string{
	CountingType_Estimate: "Estimate",
	CountintType_Exact:    "Exact Number",
	CountingType_Instance: "Instance",
}

const (
	CheckoutMode_FreeToTake = iota
	CheckoutMode_Checkout
	CheckoutMode_NoCheckout
)

var InventoryClassCheckoutModes = map[int]string{
	CheckoutMode_FreeToTake: "Free To Take",
	CheckoutMode_Checkout:   "Checkout Required",
	CheckoutMode_NoCheckout: "Checkout Not Allowed",
}

type InventoryItemClass struct {
	BaseModel
	ItemName         *string `gorm:"NOT NULL"`
	ItemLabel        *string `gorm:"NOT NULL"`
	ItemLabelID      string  `gorm:"unique; NOT NULL"`
	ItemCountingType int     // Estimate_Counting, ExactCounting, InstanceCounting
	ItemCount        int64   // Total Stock
	ItemCountInStock int64   // Current Stock
	ItemDescription  string
	ItemDatasheet    string // URL for datasheet
	ItemCheckoutMode int    // FreeToTake, Lending, NonCheckout
	ItemParameters   string // DIP8... etc
	ItemLocation     string

	ItemInstances []*InventoryItem `gorm:"foreignkey:item_class_id"json:"-"`
}

type InventoryItem struct {
	BaseModel

	ItemClassID EntityIDType
	ItemClass   *InventoryItemClass

	ItemInstanceId string `gorm:"unique"`

	ItemInstanceNote string
}
