package user

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kjh123/huiGo/log"
)

type UserModel struct {
	Name string 	`sql:"type:varchar(32);not null"`
	Password string `sql:"type:varchar(32)"`
	NickName string `sql:"type:varchar(32)"`
	gorm.Model
}

func (m UserModel) TableName() string {
	return "users"
}

// User Table Migrate
func Migrate(db *gorm.DB, modelName string) error {
	if err := db.CreateTable(new(UserModel)).Error; err != nil {
		log.ERROR.Printf("Error create table : %s", modelName)
		return fmt.Errorf("Error create table : %s", err)
	}
	return nil
}
