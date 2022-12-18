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

// POST /problems
func tryCreateProblemHandler(c echo.Context) error {
	userID := c.FormValue("userID")
	if userID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "userID is required")
	}

	problem := model.Problem{}
	if err := c.Bind(&problem); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id, err := model.TryCreateProblem(c.Request().Context(), userID, problem.Score, problem.Title, problem.Text, problem.Language)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, id)
}

// GET /problems/:problemID
func getProblemHandler(c echo.Context) error {
	userID := c.FormValue("userID")
	if userID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "userID is required")
	}
	problemID := c.Param("problemID")
	if problemID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "problemID is required")
	}

	problem, err := model.GetProblemByUser(c.Request().Context(), problemID, userID)
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

// POST /step
func stepExecuteHandler(c echo.Context) error {
	code := c.FormValue("code")
	state := c.FormValue("state")
	problemID := c.FormValue("problemID")
	if problemID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "problemID is required")
	}
	fmt.Println(code, state, problemID)

	problem, err := model.GetProblem(c.Request().Context(), problemID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	result, err := model.StepExecute(c.Request().Context(), code, state, problem.Language)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	type Result struct {
		Output  string `json:"output"`
		IsEnded bool   `json:"isEnded"`
	}
	if result[len(result)-1] == 'T' {
		res := Result{Output: result[:len(result)-1], IsEnded: true}
		return c.JSON(http.StatusOK, res)
	} else {
		res := Result{Output: result[:len(result)-1], IsEnded: false}
		return c.JSON(http.StatusOK, res)
	}
}
