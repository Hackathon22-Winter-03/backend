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
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173", "https://turing.trap.games"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderSetCookie},
		AllowCredentials: true,
	}))

	e.POST("/api/ping", pingHandler)
	e.POST("/api/echo", echoHandler)
	e.POST("/api/users", getUsersHandler)
	e.POST("/api/users/create", postUserHandler)
	e.POST("/api/users/:userID", getUserHandler)
	e.POST("/api/problems", getProblemsHandler)
	e.POST("/api/problems/create", tryCreateProblemHandler)
	e.POST("/api/problems/:problemID", getProblemHandler)
	e.POST("/api/problems/:problemID/codes", getCodesHandler)
	e.POST("/api/problems/:problemID/codes/:codeID", getCodeHandler)
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
