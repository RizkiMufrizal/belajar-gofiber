package main

import (
	"github.com/RizkiMufrizal/belajar-gofiber/configuration"
	"github.com/RizkiMufrizal/belajar-gofiber/controller"
	"github.com/RizkiMufrizal/belajar-gofiber/exception"
	repository "github.com/RizkiMufrizal/belajar-gofiber/repository/impl"
	service "github.com/RizkiMufrizal/belajar-gofiber/service/impl"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	//setup configuration
	config := configuration.New()
	database := configuration.NewDatabase(config)

	//repository
	productRepository := repository.NewProductRepositoryImpl(database)

	//service
	productService := service.NewProductServiceImpl(&productRepository)

	//controller
	productController := controller.NewProductController(&productService, config)

	//setup fiber
	app := fiber.New(configuration.NewFiberConfiguration())
	app.Use(recover.New())
	app.Use(cors.New())

	//routing
	productController.Route(app)

	//start app
	err := app.Listen(config.Get("SERVER.PORT"))
	exception.PanicLogging(err)
}
