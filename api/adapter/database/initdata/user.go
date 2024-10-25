package initdata

import (
	"gitlab.com/digeon-inc/japan-association-for-clinical-engineers/e-privado/api/adapter/database/model"
)

// User returns []*model.User（例）
// このように、固定のマスタデータなどは初期投入データを関数で定義しておくと、jsonより楽な場合が多い
func User() []*model.User {
	return []*model.User{
		{
			ID:    "system-admin",
			StudentID: "k000001",
			IdmUniv:   "0000000001",
			IdmBus:    "0987654321",
			UserType:  "system-admin",
			IsDeleted: false,
		},
		{
			ID:    "admin",
			StudentID: "k000002",
			IdmUniv:   "1234567890",
			IdmBus:    "0987654321",
			UserType:  "admin",
			IsDeleted: false,
		},
		{
			ID:    "user",
			StudentID: "k000003",
			IdmUniv:   "1234567890",
			IdmBus:    "0987654321",
			UserType:  "user",
			IsDeleted: false,
		},
		{
			ID:    "user",
			StudentID: "k000004",
			IdmUniv:   "1234567890",
			IdmBus:    "0987654321",
			UserType:  "user",
			IsDeleted: false,
		},
	}
}
