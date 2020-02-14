package entity

import "github.com/ECEHive/myHive-backend/db"

func MigrateEntities() {
	db.GetDB().AutoMigrate(
		// Workbench
		Workbench{},
		WorkbenchItem{},
		WorkbenchRentalRecord{},
		// Inventory
		InventoryItemClass{},
		InventoryItem{},
		// User & Accounts
		HiveUser{},
		Sequence{},
		InventoryCheckoutRecord{},
	)
}
