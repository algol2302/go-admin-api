package migration

import (
	"github.com/algol2302/go-admin-api/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}
