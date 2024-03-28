package router

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/tnnz20/final-hacktiv8-fga/internal/adapters/handler/http"
)

func NewUserRouter(e *echo.Echo, h *http.UserHandler, config echojwt.Config) {

	// Public Route
	user := e.Group("/users")
	user.POST("/register", h.Register)
	user.POST("/login", h.Login)

	// Protected Route
	protected := e.Group("/users")

	protected.Use(echojwt.WithConfig(config))
	protected.PUT("/update", h.Update)
	protected.DELETE("", h.Delete)
}

func NewPhotoRouter(e *echo.Echo, h *http.PhotoHandler, config echojwt.Config) {

	// Public Route
	photo := e.Group("/photos")
	photo.GET("", h.GetAll)
	photo.GET("/:photoId", h.GetByID)

	// Protected Route
	protected := e.Group("/photos")
	protected.Use(echojwt.WithConfig(config))
	protected.POST("", h.Create)
	protected.PUT("/:photoId", h.Update)
	protected.DELETE("/:photoId", h.Delete)
}

func NewCommentRouter(e *echo.Echo, h *http.CommentHandler, config echojwt.Config) {

	// Public Route
	// comment := e.Group("/comments")
	// comment.GET("", h.GetAll)
	// comment.GET("/:id", h.GetByID)

	// Protected Route
	protected := e.Group("/comments")
	protected.Use(echojwt.WithConfig(config))
	protected.POST("", h.Create)
	// protected.PUT("/:id", h.Update)
	// protected.DELETE("/:id", h.Delete)
}
