package data

import (
	"github.com/mikolajsemeniuk/Go-React-Fullstack/domain"
	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	result := Context.Exec("PRAGMA foreign_keys = ON", nil)
	if result.Error != nil {
		panic(result.Error)
	}

	var roles []domain.Role
	if Context.Find(&roles).RowsAffected == 0 {
		roles = []domain.Role{
			{
				Name: "admin",
			},
			{
				Name: "moderator",
			},
		}
		Context.Create(roles)
	}

	var accounts []domain.Account
	if Context.Find(&accounts).RowsAffected == 0 {

		adminPassword, err := bcrypt.GenerateFromPassword([]byte("P@ssw0rd"), 14)
		if err != nil {
			panic(err)
		}

		moderatorPassword, err := bcrypt.GenerateFromPassword([]byte("P@ssw0rd"), 14)
		if err != nil {
			panic(err)
		}

		accounts = []domain.Account{
			{
				Email:    "admin@example.com",
				Username: "admin",
				Password: adminPassword,
			},
			{
				Email:    "moderator@example.com",
				Username: "moderator",
				Password: moderatorPassword,
			},
		}

		Context.Create(accounts)
	}

	var accountRoles []domain.AccountRole
	if Context.Find(&accountRoles).RowsAffected == 0 {
		accountRoles = append(accountRoles, domain.AccountRole{
			AccountId: accounts[0].Id,
			RoleId:    roles[0].Id,
		})
		accountRoles = append(accountRoles, domain.AccountRole{
			AccountId: accounts[0].Id,
			RoleId:    roles[1].Id,
		})
		accountRoles = append(accountRoles, domain.AccountRole{
			AccountId: accounts[1].Id,
			RoleId:    roles[1].Id,
		})

		Context.Create(accountRoles)
	}
}
