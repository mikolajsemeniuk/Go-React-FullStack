package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Entity struct {
	Id        uuid.UUID  `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}

func (*Entity) BeforeCreate(scope *gorm.DB) error {
	id, err := uuid.NewUUID()
	scope.Statement.SetColumn("Id", id)
	return err
}
