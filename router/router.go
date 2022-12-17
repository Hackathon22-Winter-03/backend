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
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173", "https://hackathon22-winter-03.trap.jp"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "x-master-version", "x-session"},
		AllowCredentials: true,
	}))

	e.GET("/api/ping", pingHandler)
	e.GET("/api/echo", echoHandler)
	e.GET("/api/users", getUsersHandler)
	e.POST("/api/users", postUserHandler)
	e.GET("/api/users/:userID", getUserHandler)
	e.GET("/api/problems", getProblemsHandler)
	e.GET("/api/problems/:problemID", getProblemHandler)
	e.GET("/api/problems/:problemID/codes", getCodesHandler)
	e.GET("/api/problems/:problemID/codes/:codeID", getCodeHandler)
	e.POST("/api/problems/:problemID/submit", submitCodeHandler)

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
