package entity

type InventoryItemClass struct {
	BaseModel
	ItemName         string
	ItemLabel        string
	ItemLabelID      string `gorm:"unique"`
	ItemCountingType int    // Estimate_Counting, ExactCounting, InstanceCounting
	ItemCount        int64  // Total Stock
	ItemCountInStock int64  // Current Stock
	ItemDescription  string
	ItemDatasheet    string // URL for datasheet
	ItemCheckoutMode int    // FreeToTake, Lending, NonCheckout
	ItemParameters   string // DIP8... etc
}
