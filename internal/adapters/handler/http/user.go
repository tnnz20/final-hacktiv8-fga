package http

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"github.com/tnnz20/final-hacktiv8-fga/internal/adapters/handler/http/middleware"
	"github.com/tnnz20/final-hacktiv8-fga/internal/core/domain"
	"github.com/tnnz20/final-hacktiv8-fga/internal/core/port"
	"github.com/tnnz20/final-hacktiv8-fga/pkg/utils"
	"github.com/tnnz20/final-hacktiv8-fga/pkg/validation"
)

type UserHandler struct {
	UserService port.UserService
	Validate    *validator.Validate
}

func NewUserHandler(service port.UserService, validate *validator.Validate) *UserHandler {
	return &UserHandler{
		UserService: service,
		Validate:    validate,
	}
}

func (h *UserHandler) Register(c echo.Context) error {
	var req domain.CreateUserRequest

	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.Validate.Struct(req); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			err := validation.NewErrResponse(ve)
			return utils.ErrorResponse(c, http.StatusBadRequest, err)
		}
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.UserService.Register(c.Request().Context(), &req)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusCreated, "User registered successfully", res)
}

func (h *UserHandler) Login(c echo.Context) error {
	var req domain.LoginUserRequest

	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.Validate.Struct(req); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			err := validation.NewErrResponse(ve)
			return utils.ErrorResponse(c, http.StatusBadRequest, err)
		}
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.UserService.Login(c.Request().Context(), &req)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Successfully Login", res)
}

func (h *UserHandler) Update(c echo.Context) error {
	var req domain.UpdateUserRequest

	if err := h.Validate.Struct(req); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			err := validation.NewErrResponse(ve)
			return utils.ErrorResponse(c, http.StatusBadRequest, err)
		}
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JWTCustomClaims)
	claimId := claims.ID

	req.ID = int(claimId)
	res, err := h.UserService.Update(c.Request().Context(), &req)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Successfully update user", res)
}

func (h *UserHandler) Delete(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JWTCustomClaims)
	claimId := claims.ID

	err := h.UserService.Delete(c.Request().Context(), int(claimId))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Successfully delete user", nil)
}
