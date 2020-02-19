package util

import (
	"github.com/ECEHive/myHive-backend/constants"
)

func CheckoutStatusInAllStatus(str constants.InventoryCheckoutStatus, slice []constants.InventoryCheckoutStatus) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
