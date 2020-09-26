package domain

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

func (p *Product) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("Uuid", uuid.NewV4().String())
}