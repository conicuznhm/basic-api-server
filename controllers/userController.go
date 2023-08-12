package controllers

import (
	"RCSserver/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetProfile(c echo.Context) error{
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error":"Invalid id"})
	}

	profile, err := services.FetchProfile(id)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error":err.Error()})
	}
	return c.JSON(http.StatusOK, profile)
}

func GetAllProfile(c echo.Context) error{
	profiles, err := services.FetchAllProfile()
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error":err.Error()})
	}
	return c.JSON(http.StatusOK, profiles)
}