package main

import (
	"github.com/mikolajsemeniuk/Go-React-Fullstack/application"
	"github.com/mikolajsemeniuk/Go-React-Fullstack/data"
)

func main() {
	data.Seed()
	application.Listen()
}
