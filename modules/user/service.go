package user

import (
    "github.com/jinzhu/gorm"
)

type Service struct {
    db *gorm.DB
}

func NewUserService(db *gorm.DB) *Service {
    return &Service{
        db: db,
    }
}

func (s Service) Prefix() string {
    return "/v1/user"
}

func (s *Service) Close() {}
