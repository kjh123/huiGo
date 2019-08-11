package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kjh123/huiGo/config"
	"time"
)

func init()  {
	gorm.NowFunc = func() time.Time {
		return time.Now().UTC()
	}
}

func NewDatabase(config *config.Config) (*gorm.DB, error) {
	if config.DatabaseConfig.Type == "sqlite3" {
		arg := fmt.Sprintf("%s", config.DatabaseConfig.File)
		db, err := gorm.Open("sqlite3", arg)
		if err != nil {
			return db, err
		}
		// setMaxIdleConns
		db.DB().SetMaxIdleConns(config.DatabaseConfig.MaxIdleConns)

		// SetMaxOpenConns
		db.DB().SetMaxOpenConns(config.DatabaseConfig.MaxOpenConns)

		// Development
		db.LogMode(config.IsDevelopment)

		return db, err
	}
	return nil, fmt.Errorf("Database type %s not supported", config.DatabaseConfig.Type)
}
