package migrations

import (
    "github.com/jinzhu/gorm"
    "github.com/kjh123/huiGo/modules/user"
)

type Migration struct {
    gorm.Model
    Name string `sql:"size:255";DEFAULT:"bootStrap_migrations"`
}

var (
    list = [] MigrationStage{
        {
            // user migration
            ModelName: user.UserModel{}.TableName(),
            Function:  user.Migrate,
        },
    }
)

func MigrateAll(db *gorm.DB) error {
    return Migrate(db, list)
}
