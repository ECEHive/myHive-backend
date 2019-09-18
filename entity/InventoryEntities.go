package entity

type InventoryItem struct {
	BaseModel
	ItemName string
	ItemLabel string
	ItemLabelID string `gorm:"unique"`
	ItemCountingType int
	ItemCount int64
	ItemCountInStock int64
	ItemDescription string
	ItemDatasheet string
	ItemCheckoutMode string
	ItemPackage string
}
