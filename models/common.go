package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type CommonModel struct {
	gorm.Model
	ModelInterface
	DeletedAt *time.Time
}
