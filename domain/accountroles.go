package domain

import (
	"time"

	"github.com/google/uuid"
)

type AccountRole struct {
	AccountId uuid.UUID  `gorm:"primaryKey"`
	Account   Account    `gorm:"foreignKey:AccountId;references:Id"`
	RoleId    uuid.UUID  `gorm:"primaryKey"`
	Role      Role       `gorm:"foreignKey:RoleId;references:Id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}
