package domain

type Role struct {
	Entity
	Name     string    `gorm:"unique"`
	Accounts []Account `gorm:"many2many:account_roles;"`
}
