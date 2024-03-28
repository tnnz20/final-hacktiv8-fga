package http

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/tnnz20/final-hacktiv8-fga/internal/core/domain"
	"github.com/tnnz20/final-hacktiv8-fga/internal/core/port"
	"github.com/tnnz20/final-hacktiv8-fga/pkg/utils"
	"github.com/tnnz20/final-hacktiv8-fga/pkg/validation"
)

type PhotoHandler struct {
	PhotoService port.PhotoService
	Validate     *validator.Validate
}

func NewPhotoHandler(service port.PhotoService, validate *validator.Validate) *PhotoHandler {
	return &PhotoHandler{
		PhotoService: service,
		Validate:     validate,
	}
}

func (h *PhotoHandler) Create(c echo.Context) error {
	var req domain.CreatePhotoRequest

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

	userId := utils.ClaimID(c)

	res, err := h.PhotoService.Create(c.Request().Context(), &req, userId)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusCreated, "Photo created successfully", res)
}

func (h *PhotoHandler) GetAll(c echo.Context) error {
	res, err := h.PhotoService.GetAll(c.Request().Context())
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Get all photo successfully", res)
}

func (h *PhotoHandler) GetByID(c echo.Context) error {

	photoId, err := strconv.Atoi((c.Param("photoId")))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, domain.ErrInvalidId.Error())
	}

	res, err := h.PhotoService.GetByID(c.Request().Context(), photoId)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Get photo by id successfully", res)
}

func (h *PhotoHandler) Update(c echo.Context) error {
	var req domain.UpdatePhotoRequest

	photoId, err := strconv.Atoi((c.Param("photoId")))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, domain.ErrInvalidId.Error())
	}

	userId := utils.ClaimID(c)

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

	res, err := h.PhotoService.Update(c.Request().Context(), &req, photoId, userId)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Photo updated successfully", res)
}

func (h *PhotoHandler) Delete(c echo.Context) error {

	photoId, err := strconv.Atoi((c.Param("photoId")))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, domain.ErrInvalidId.Error())
	}

	userId := utils.ClaimID(c)
	err = h.PhotoService.Delete(c.Request().Context(), photoId, userId)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Photo deleted successfully", nil)
}
