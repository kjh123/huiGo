package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kjh123/huiGo/log"
	"github.com/kjh123/huiGo/user"
	"github.com/kjh123/huiGo/util/migrations"
)

var (
	list = [] migrations.MigrationStage{
		{
			// user migration
			Model: &user.UserModel{},
			Function: func(db *gorm.DB, m ModelInterface) error {
				if err := db.CreateTable(m).Error; err != nil {
					log.ERROR.Printf("Error create table : %s", m.TableName())
					return fmt.Errorf("Error create table : %s", err)
				}
				return nil
			},
		},
	}
)

func MigrateAll(db *gorm.DB) error {
	return migrations.Migrate(db, list)
}
