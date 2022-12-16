package router

import (
	"net/http"

	"github.com/Hackathon22-Winter-03/backend/model"
	"github.com/labstack/echo/v4"
)

// GET /users
func getUsersHandler(c echo.Context) error {
	users, err := model.GetUsers(c.Request().Context())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

// GET /users/:userID
func getUserHandler(c echo.Context) error {
	userID := c.Param("userID")

	user, err := model.GetUser(c.Request().Context(), userID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

// POST /users
func postUserHandler(c echo.Context) error {
	user := model.User{}

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := model.CreateUser(c.Request().Context(), user.ID, user.Name, user.Comment); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, nil)
}
