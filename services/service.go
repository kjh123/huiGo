package services

import (
    "github.com/jinzhu/gorm"
    "github.com/kjh123/huiGo/config"
    "github.com/kjh123/huiGo/contracts"
    "github.com/kjh123/huiGo/modules/user"
    "reflect"
)

var (
    UserService contracts.ServiceInterface
)

func Init(conf *config.Config, db *gorm.DB) error {
    if nil == reflect.TypeOf(UserService) {
        UserService = user.NewUserService(db)
    }
    return nil
}

func Modules() []contracts.ServiceInterface {
    return []contracts.ServiceInterface{
        UserService,
    }
}

func Close() {
    UserService.Close()
}
