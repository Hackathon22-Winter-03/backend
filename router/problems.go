package router

import (
	"net/http"

	"github.com/Hackathon22-Winter-03/backend/model"
	"github.com/labstack/echo/v4"
)

// GET /problems
func getProblemsHandler(c echo.Context) error {
	userID := c.FormValue("userID")
	if userID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "userID is required")
	}

	problems, err := model.GetProblemsByUser(c.Request().Context(), userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, problems)
}

// GET /problems/:problemID
func getProblemHandler(c echo.Context) error {
	problemID := c.Param("problemID")
	if problemID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "problemID is required")
	}

	problem, err := model.GetProblem(c.Request().Context(), problemID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, problem)
}

// GET /problems/:problemID/codes
func getCodesHandler(c echo.Context) error {
	problemID := c.Param("problemID")
	userID := c.FormValue("userID")
	if problemID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "problemID is required")
	}
	if userID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "userID is required")
	}

	codes, err := model.GetCodes(c.Request().Context(), userID, problemID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, codes)
}

// GET /problems/:problemID/codes/:codeID
func getCodeHandler(c echo.Context) error {
	problemID := c.Param("problemID")
	codeID := c.Param("codeID")
	if problemID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "problemID is required")
	}
	if codeID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "codeID is required")
	}

	code, err := model.GetCode(c.Request().Context(), problemID, codeID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, code)
}

// POST /problems/:problemID/submit
func submitCodeHandler(c echo.Context) error {
	problemID := c.Param("problemID")
	userID := c.FormValue("userID")
	code := c.FormValue("code")
	if problemID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "problemID is required")
	}
	if userID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "userID is required")
	}
	if code == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "code is required")
	}

	_, err := model.SubmitCode(c.Request().Context(), userID, problemID, code)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}
