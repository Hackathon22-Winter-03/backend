package router

import (
	"fmt"
	"net/http"

	"github.com/Hackathon22-Winter-03/backend/model"
	"github.com/Hackathon22-Winter-03/backend/utils"
	"github.com/labstack/echo/v4"
)

// GET /problems
func getProblemsHandler(c echo.Context) error {
	userID, err := c.Cookie("userID")
	if err != nil || userID.Value == "" {
		fmt.Println(err, userID)
		return echo.NewHTTPError(http.StatusBadRequest, "userID is required")
	}

	problems, err := model.GetProblemsByUser(c.Request().Context(), userID.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, problems)
}

// POST /problems
func tryCreateProblemHandler(c echo.Context) error {
	userID, err := c.Cookie("userID")
	if err != nil || userID.Value == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "userID is required")
	}

	problem := model.Problem{}
	if err := c.Bind(&problem); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id, err := model.TryCreateProblem(c.Request().Context(), userID.Value, problem.Score, problem.Title, problem.Text, problem.Language)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, id)
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
	userID, err := c.Cookie("userID")
	if problemID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "problemID is required")
	}
	if err != nil || userID.Value == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "userID is required")
	}

	codes, err := model.GetCodes(c.Request().Context(), userID.Value, problemID)
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
	if err == utils.ErrNotFound {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, code)
}

// POST /problems/:problemID/submit
func submitCodeHandler(c echo.Context) error {
	problemID := c.Param("problemID")
	userID, err := c.Cookie("userID")
	code := c.FormValue("code")
	if problemID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "problemID is required")
	}
	if err != nil || userID.Value == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "userID is required")
	}
	if code == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "code is required")
	}

	_, err = model.SubmitCode(c.Request().Context(), userID.Value, problemID, code)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

// POST /step
func stepExecuteHandler(c echo.Context) error {
	code := c.FormValue("code")
	if code == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "code is required")
	}
	state := c.FormValue("state")
	if state == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "state is required")
	}
	problemID := c.FormValue("problemID")
	if problemID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "problemID is required")
	}

	problem, err := model.GetProblem(c.Request().Context(), problemID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	result, err := model.StepExecute(c.Request().Context(), code, state, problem.Language)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
