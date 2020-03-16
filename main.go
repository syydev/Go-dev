package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/SeungyoungYang/Go-dev/api"
	"github.com/SeungyoungYang/Go-dev/model"
	"github.com/SeungyoungYang/Go-dev/utils"
)

func main() {
	// Echo instance
	e := echo.New()

	config := utils.GetConfig()

	// Init DB
	model.Init(config)
	model.AutoMigrate()
	defer model.Close()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", api.Index)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
