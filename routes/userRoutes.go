package routes

import (
	"RCSserver/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandlePathError(err error,c echo.Context) {
	c.JSON(http.StatusNotFound, map[string]interface{}{"error":"Route not found: 404"})
}

func UserRoutes(e *echo.Echo){
	userGroup := e.Group("/user")
	
	userGroup.GET("/profile",controllers.GetProfile)
	e.HTTPErrorHandler = HandlePathError
}