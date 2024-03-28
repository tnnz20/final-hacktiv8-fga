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

type CommentHandler struct {
	CommentService port.CommentService
	Validate       *validator.Validate
}

func NewCommentHandler(service port.CommentService, validate *validator.Validate) *CommentHandler {
	return &CommentHandler{
		CommentService: service,
		Validate:       validate,
	}
}

func (h *CommentHandler) Create(c echo.Context) error {
	var req domain.CreateCommentRequest

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

	res, err := h.CommentService.Create(c.Request().Context(), &req, userId)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusCreated, "Comment created successfully", res)
}

func (h *CommentHandler) Update(c echo.Context) error {
	var req domain.UpdateCommentRequest

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

	commentId, err := strconv.Atoi((c.Param("commentId")))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, domain.ErrInvalidId.Error())
	}

	userId := utils.ClaimID(c)

	res, err := h.CommentService.Update(c.Request().Context(), &req, commentId, userId)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Comment updated successfully", res)
}

func (h *CommentHandler) Delete(c echo.Context) error {
	commentId, err := strconv.Atoi((c.Param("commentId")))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, domain.ErrInvalidId.Error())
	}

	userId := utils.ClaimID(c)

	err = h.CommentService.Delete(c.Request().Context(), commentId, userId)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Comment deleted successfully", nil)
}

func (h *CommentHandler) GetComments(c echo.Context) error {
	res, err := h.CommentService.GetComments(c.Request().Context())
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Get all comment successfully", res)
}

func (h *CommentHandler) GetCommentByID(c echo.Context) error {

	commentId, err := strconv.Atoi((c.Param("commentId")))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, domain.ErrInvalidId.Error())
	}

	res, err := h.CommentService.GetCommentByID(c.Request().Context(), commentId)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, http.StatusOK, "Get comment by ID successfully", res)
}
