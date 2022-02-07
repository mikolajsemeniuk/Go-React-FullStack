package data

import (
	"github.com/mikolajsemeniuk/Go-React-Fullstack/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Context *gorm.DB

func init() {
	var err error
	Context, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Context.AutoMigrate(&domain.Account{}, &domain.Role{}, &domain.AccountRole{})
}

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
			{
				Name: "member",
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
			AccountId: accounts[0].Id,
			RoleId:    roles[2].Id,
		})
		accountRoles = append(accountRoles, domain.AccountRole{
			AccountId: accounts[1].Id,
			RoleId:    roles[1].Id,
		})
		accountRoles = append(accountRoles, domain.AccountRole{
			AccountId: accounts[1].Id,
			RoleId:    roles[2].Id,
		})

		Context.Create(accountRoles)
	}
}
