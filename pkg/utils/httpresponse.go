package utils

import "github.com/labstack/echo/v4"

func ErrorResponse(c echo.Context, statusCode int, message interface{}) error {
	return c.JSON(statusCode, echo.Map{
		"status":  "error",
		"message": message,
	})
}

func SuccessResponse(c echo.Context, statusCode int, message, data interface{}) error {
	if data != nil {
		return c.JSON(statusCode, echo.Map{
			"status":  "success",
			"message": message,
			"data":    data,
		})
	}
	return c.JSON(statusCode, echo.Map{
		"status":  "success",
		"message": message,
	})
}
