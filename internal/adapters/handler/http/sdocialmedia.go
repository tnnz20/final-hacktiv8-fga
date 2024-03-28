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

type SocialMediaHandler struct {
	SocialMediaService port.SocialMediaService
	Validate           *validator.Validate
}

func NewSocialMediaHandler(service port.SocialMediaService, validate *validator.Validate) *SocialMediaHandler {
	return &SocialMediaHandler{
		SocialMediaService: service,
		Validate:           validate,
	}
}

func (h *SocialMediaHandler) Create(c echo.Context) error {
	var req domain.CreateSocialMediaRequest

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

	res, err := h.SocialMediaService.Create(c.Request().Context(), &req, userId)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusCreated, "Social Media created successfully", res)
}

func (h *SocialMediaHandler) Update(c echo.Context) error {
	var req domain.UpdateSocialMediaRequest
	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

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

	res, err := h.SocialMediaService.Update(c.Request().Context(), &req, socialMediaId, userId)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Social Media updated successfully", res)
}

func (h *SocialMediaHandler) Delete(c echo.Context) error {
	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	userId := utils.ClaimID(c)

	err = h.SocialMediaService.Delete(c.Request().Context(), socialMediaId, userId)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Social Media deleted successfully", nil)
}

func (h *SocialMediaHandler) GetSocialMedias(c echo.Context) error {
	res, err := h.SocialMediaService.GetSocialMedias(c.Request().Context())
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Get all social media successfully", res)
}

func (h *SocialMediaHandler) GetSocialMediaByID(c echo.Context) error {
	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.SocialMediaService.GetSocialMediaByID(c.Request().Context(), socialMediaId)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Get social media by id successfully", res)
}
