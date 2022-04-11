package repositories

import (
	"github.com/mikolajsemeniuk/Go-React-Fullstack/data"
	"github.com/mikolajsemeniuk/Go-React-Fullstack/domain"
)

var Account IAccount = &account{}

type account struct{}

type IAccount interface {
	SingleByEmail(string) *domain.Account
}

func (*account) SingleByEmail(name string) *domain.Account {
	var account domain.Account
	result := data.Context.Where("email = ?", name).Preload("Roles.Role").Take(&account)
	if result.RowsAffected == 0 {
		return nil
	}
	return &account
}
