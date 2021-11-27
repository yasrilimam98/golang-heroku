package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/yasrilimam98/grb-restapi/models"
)

func FetchAllUsers(c echo.Context) error {
	result, err := models.FetchAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreUsers(c echo.Context) error {
	// name := c.FormValue("name")
	// email := c.FormValue("email")
	// email_verified_at := c.FormValue("email_verified_at")
	// password := c.FormValue("password")
	// remember_token := c.FormValue("remember_token")
	// created_at := c.FormValue("created_at")
	// updated_at := c.FormValue("updated_at")

	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telepon := c.FormValue("telepon")

	result, err := models.StoreUsers(nama, alamat, telepon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateUsers(c echo.Context) error {
	// id := c.FormValue("id")
	// name := c.FormValue("name")
	// email := c.FormValue("email")
	// email_verified_at := c.FormValue("email_verified_at")
	// password := c.FormValue("password")
	// remember_token := c.FormValue("remember_token")
	// created_at := c.FormValue("created_at")
	// updated_at := c.FormValue("updated_at")

	id := c.FormValue("id")
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telepon := c.FormValue("telepon")

	// konversi str parsing int to string
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateUsers(conv_id, nama, alamat, telepon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteUsers(c echo.Context) error {
	id := c.FormValue("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteUsers(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
