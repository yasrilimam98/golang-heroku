package routes

// Import dengan parameter agar bisa banyak import pkg
import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/yasrilimam98/grb-restapi/controllers"
	"github.com/yasrilimam98/grb-restapi/middleware"
)

func Init() *echo.Echo {

	// c echo context untuk mengakses request or respon dan error untuk notif ketika jadi kesalahan
	e := echo.New()

	// Cek connection database
	e.GET("/database1", func(c echo.Context) error {
		return c.String(http.StatusOK, "Database Connection 1")
	})
	e.GET("/database2", func(c echo.Context) error {
		return c.String(http.StatusOK, "Database Connection 2")
	})

	// Database 1 table users routes
	e.GET("/users", controllers.FetchAllUsers)
	e.POST("/users", controllers.StoreUsers)
	e.PUT("/users", controllers.UpdateUsers)
	e.DELETE("/users", controllers.DeleteUsers)

	// ================BATAS===================

	// Database 2 table mclient routes
	e.GET("/client", controllers.FetchAllClient, middleware.IsAuthenticated)

	// ================BATAS===================

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)

	return e
}
