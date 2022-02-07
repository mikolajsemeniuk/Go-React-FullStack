package main

import (
	"github.com/mikolajsemeniuk/Go-React-Fullstack/application"
	"github.com/mikolajsemeniuk/Go-React-Fullstack/data"
	"github.com/mikolajsemeniuk/Go-React-Fullstack/domain"
)

func main() {
	data.Context.AutoMigrate(&domain.Account{}, &domain.Role{})
	data.Seed()
	application.Listen()
}
