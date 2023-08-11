package controllers

import (
	"RCSserver/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetProfile(c echo.Context) error{
	profile, err := services.FetchProfile()
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error":err.Error()})
	}
	return c.JSON(http.StatusOK, profile)
}