package domain

type Role struct {
	Entity
	Name     string
	Accounts []Account `gorm:"many2many:account_roles"`
}
