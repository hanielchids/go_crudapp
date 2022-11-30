package main

import (
	"test/init/bootstrap"
	"test/init/repository"

	"github.com/gofiber/fiber/v2"
)

type Repository repository.Repository


func main()  {
	app:= fiber.New()
	bootstrap.InitializeApp(app)

}