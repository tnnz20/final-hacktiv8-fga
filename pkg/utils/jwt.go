package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/tnnz20/final-hacktiv8-fga/internal/adapters/handler/http/middleware"
)

func ClaimID(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JWTCustomClaims)
	claimId := claims.ID
	return claimId
}
