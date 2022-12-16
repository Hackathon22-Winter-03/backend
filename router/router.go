package router

import (
	"fmt"
	"net/http"

	"github.com/Hackathon22-Winter-03/backend/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRouting() (*echo.Echo, error) {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
		AllowHeaders: []string{"Content-Type", "x-master-version", "x-session"},
	}))

	e.GET("/ping", pingHandler)
	e.GET("/echo", echoHandler)
	e.GET("/users", getUsersHandler)
	e.GET("/users/:userID", getUserHandler)
	e.GET("/problems", getProblemsHandler)
	e.GET("/problems/:problemID", getProblemHandler)
	e.GET("/problems/:problemID/codes", getCodesHandler)
	e.GET("/problems/:problemID/codes/:codeID", getCodeHandler)
	e.POST("/problems/:problemID/submit", submitCodeHandler)

	port := utils.GetEnv("PORT", ":3000")
	err := e.Start(port)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func pingHandler(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func echoHandler(c echo.Context) error {
	header := fmt.Sprintf("%#v", c.Request().Header)
	return c.String(http.StatusOK, header)
}
