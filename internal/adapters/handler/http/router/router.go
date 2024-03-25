package router

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/tnnz20/final-hacktiv8-fga/internal/adapters/handler/http"
	"github.com/tnnz20/final-hacktiv8-fga/internal/adapters/handler/http/middleware"
)

func NewUserRouter(e *echo.Echo, key *string, h *http.UserHandler) {

	// Public Route
	user := e.Group("/users")
	user.POST("/register", h.Register)
	user.POST("/login", h.Login)

	// Protected Route
	protected := e.Group("/users")

	config := middleware.JWTconfig(key)

	protected.Use(echojwt.WithConfig(config))
	protected.PUT("/update", h.Update)
	protected.DELETE("", h.Delete)
}
