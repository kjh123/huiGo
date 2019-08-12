package user

import "github.com/kjh123/huiGo/models"

type UserModel struct {
	models.CommonModel
	Name string `sql:"type:varchar(32);not null"`
	Password string `sql:"type:varchar(32)"`
	NickName string `sql:"type:varchar(32);"`
}

func (u *UserModel) TableName() string {
	return "users"
}
