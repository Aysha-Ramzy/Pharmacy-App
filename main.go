package main

import (
	"github.com/Aysha-Ramzy/Pharmacy-App/bootstrap"
	"github.com/Aysha-Ramzy/Pharmacy-App/repository"
	"github.com/gofiber/fiber/v2"
)

type Repository repository.Repository

	func main(){
		app:= fiber.New()
bootstrap.InitializeApp(app)
	}
