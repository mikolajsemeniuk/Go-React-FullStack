package domain

type Account struct {
	Entity
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password []byte
	Roles    []AccountRole
}
