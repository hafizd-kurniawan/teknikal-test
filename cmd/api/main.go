package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"employee-management/app/departement"
	"employee-management/app/employee"
	"employee-management/app/location"
	"employee-management/app/position"
	"employee-management/config"
	"employee-management/infra/middleware"
	"employee-management/pkg/database"
)

func main() {

	app := fiber.New(fiber.Config{
		AppName: "employee management",
		Prefork: true,
	})
	app.Use(logger.New())

	err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Println("Error saat meload config: ", err.Error())
	}

	jwt := config.Cfg.JWT
	middleware.SetJWTSecretKey(jwt.Secret)

	db, err := database.ConnectSQLXPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	departement.Init(app, db)
	position.Init(app, db)
	location.Init(app, db)
	employee.Init(app, db)
	app.Listen(config.Cfg.App.Port)

}
