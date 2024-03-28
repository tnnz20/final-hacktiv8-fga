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
	photo.GET("", h.GetPhotos)
	photo.GET("/:photoId", h.GetPhotoByID)

	// Protected Route
	protected := e.Group("/photos")
	protected.Use(echojwt.WithConfig(config))
	protected.POST("", h.Create)
	protected.PUT("/:photoId", h.Update)
	protected.DELETE("/:photoId", h.Delete)
}

func NewCommentRouter(e *echo.Echo, h *http.CommentHandler, config echojwt.Config) {

	// Public Route
	comment := e.Group("/comments")
	comment.GET("", h.GetComments)
	comment.GET("/:commentId", h.GetCommentByID)

	// Protected Route
	protected := e.Group("/comments")
	protected.Use(echojwt.WithConfig(config))
	protected.POST("", h.Create)
	protected.PUT("/:commentId", h.Update)
	protected.DELETE("/:commentId", h.Delete)
}

func NewSocialMediasRouter(e *echo.Echo, h *http.SocialMediaHandler, config echojwt.Config) {

	// Public Route
	socialMedia := e.Group("/socialmedias")
	socialMedia.GET("", h.GetSocialMedias)
	socialMedia.GET("/:socialMediaId", h.GetSocialMediaByID)

	// Protected Route
	protected := e.Group("/socialmedias")
	protected.Use(echojwt.WithConfig(config))
	protected.POST("", h.Create)
	protected.PUT("/:socialMediaId", h.Update)
	protected.DELETE("/:socialMediaId", h.Delete)
}
