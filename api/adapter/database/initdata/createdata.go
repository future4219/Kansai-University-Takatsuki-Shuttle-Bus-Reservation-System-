package initdata

import (
	"fmt"

	"gorm.io/gorm"

	"gitlab.com/digeon-inc/japan-association-for-clinical-engineers/e-privado/api/adapter/database/model"
)

func CreateStgSeedData(db *gorm.DB) error {
	if err := CreateFirstUser(db); err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

func CreateDevSeedData(db *gorm.DB) error {
	if err := CreateUser(db); err != nil {
		return err
	}
	return nil
}

func CreateUser(db *gorm.DB) error {
	users := User()
	return db.Create(users).Error
}

func CreateFirstUser(db *gorm.DB) error {
	users := []*model.User{
		{
			ID:    "system-admin",
			StudentID: "k000001",
			IdmUniv:   "0000000001",
			IdmBus:    "0987654321",
			UserType:  "system-admin",
			IsDeleted: false,
		},
	}
	return db.Save(users).Error
}
