package domain

type Role struct {
	Entity
	Name     string `gorm:"unique"`
	Accounts []AccountRole
}
