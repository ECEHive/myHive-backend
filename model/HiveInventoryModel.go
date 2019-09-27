package model

type InventoryItemClassSearchRequest struct {
	SearchKeyword string
	NameOnly      bool
	LabelOnly     bool
	InStock       bool
}
