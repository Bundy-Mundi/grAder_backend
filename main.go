package main

import (
	"net/http"
	"os"

	"github.com/Bundy-Mundi/graderbackend/gethome"
	"github.com/Bundy-Mundi/graderbackend/smc2018fall"
	"github.com/Bundy-Mundi/graderbackend/smc2019spring"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// HOME
func getCollege(c echo.Context) error {
	data := gethome.GETCollege()
	return c.JSONBlob(http.StatusOK, data)
}

// SMC 2019
func smc2019(c echo.Context) error {
	return c.File("templates/2019S.html")
}

func main() {
	e := echo.New()
	port := os.Getenv("PORT")

	// Middlewares
	e.Use(middleware.CORS())

	// Common API
	e.GET("/api/v1/college", getCollege)

	// SMC
	e.GET("/api/v1/smc/2019", smc2019)
	//e.GET("/api/v1/smc/2018", get2018)

	// SMC 2019 Spring
	e.GET("/api/v1/smc/2019/spring", smc2019spring.AllData)
	e.GET("/api/v1/smc/2019/spring/:id", smc2019spring.GetByID)
	e.GET("/api/v1/smc/2019/spring/prof/:name", smc2019spring.SearchByProfessor)
	e.GET("/api/v1/smc/2019/spring/class/:name", smc2019spring.SearchByClass)

	// SMC 2018 Fall
	e.GET("/api/v1/smc/2018/fall", smc2018fall.AllData)
	e.GET("/api/v1/smc/2018/fall/:id", smc2018fall.GetByID)
	e.GET("/api/v1/smc/2018/fall/prof/:name", smc2018fall.SearchByProfessor)
	e.GET("/api/v1/smc/2018/fall/class/:name", smc2018fall.SearchByClass)

	e.Logger.Fatal(e.Start(":" + port))
}
