package constants

type InventoryCheckoutStatus string

const (
	InventoryCheckoutStatusCheckedOut   InventoryCheckoutStatus = "CheckedOut"
	InventoryCheckoutStatusExtended     InventoryCheckoutStatus = "Extended"
	InventoryCheckoutStatusReturned     InventoryCheckoutStatus = "Returned"
	InventoryCheckoutStatusLostOrDamage InventoryCheckoutStatus = "LostOrDamage"
)

var InventoryCheckoutStatusAll = []InventoryCheckoutStatus{
	InventoryCheckoutStatusCheckedOut,
	InventoryCheckoutStatusExtended,
	InventoryCheckoutStatusReturned,
	InventoryCheckoutStatusLostOrDamage,
}
