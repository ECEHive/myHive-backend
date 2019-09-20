package model

import "github.com/ECEHive/myHive-backend/entity"

type HiveUserUpsertRequest struct {
	UniqueIdentifier string `binding:"required"`
	Patch            *entity.HiveUser
}
