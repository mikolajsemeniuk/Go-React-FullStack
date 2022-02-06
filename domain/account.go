package domain

type Account struct {
	Entity
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password []byte
	Roles    []Role `gorm:"many2many:account_roles"`
}
