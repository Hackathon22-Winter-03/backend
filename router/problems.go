package router

import "github.com/labstack/echo/v4"

// GET /problems
func getProblemsHandler(c echo.Context) error {
	return nil
}

// GET /problems/:problemID
func getProblemHandler(c echo.Context) error {
	return nil
}

// GET /problems/:problemID/codes
func getCodesHandler(c echo.Context) error {
	return nil
}

// GET /problems/:problemID/codes/:codeID
func getCodeHandler(c echo.Context) error {
	return nil
}

// POST /problems/:problemID/submit
func submitProblemHandler(c echo.Context) error {
	return nil
}
