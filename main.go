package main

import (
	"RCSserver/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main(){
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	routes.UserRoutes(e)
	
	port := ":8888"
	e.Logger.Fatal(e.Start(port))
}
