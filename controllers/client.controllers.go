package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/yasrilimam98/grb-restapi/models"
)

func FetchAllClient(c echo.Context) error {
	regno := c.FormValue("regno")
	serial := c.FormValue("serial")
	// Konversi int to str regno
	conv_regno, err := strconv.Atoi(regno)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	result, err := models.FetchAllClient(conv_regno, serial)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

// func FetchAllClient(c echo.Context) error {
// 	result, err := models.FetchAllClient()
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, result)
// }
