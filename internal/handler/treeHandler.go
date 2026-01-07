package handler

import (
	"arvore/internal/model"
	"arvore/internal/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TreeHandler struct {
	service *service.TreeService
}

func NewTreeHandler(service *service.TreeService) *TreeHandler {
	return &TreeHandler{service: service}
}

func (h *TreeHandler) CreateTreeHandler(c echo.Context) error {
	var req model.TreeRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "dados inválidos",
		})
	}

	if req.Latitude == 0 || req.Longitude == 0 || req.Height == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "latitude, longitude e altura são obrigatórios",
		})
	}

	tree, err := h.service.CreateTreeService(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, tree)
}

func (h *TreeHandler) ListTreesByBoundingBox(c echo.Context) error {
	var req model.ListTreesByBoundingBoxRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "parâmetros inválidos",
		})
	}

	trees, err := h.service.LisTreesByBoundingBoxService(
		c.Request().Context(), req,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, trees)
}

func (h *TreeHandler) GetTreeByID(c echo.Context) error {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "id inválido",
		})
	}

	tree, err := h.service.GetTreeByIdService(
		c.Request().Context(),
		id,
	)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, tree)
}

func (h *TreeHandler) UpdateTree(c echo.Context) error {
	var req model.UpdateTreeRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "dados inválidos",
		})
	}

	if req.ID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "id é obrigatório",
		})
	}

	tree, err := h.service.UpdateTreeService(
		c.Request().Context(),
		req,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, tree)
}

func (h *TreeHandler) DeleteTree(c echo.Context) error {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "id inválido",
		})
	}

	err = h.service.DeleteTreeService(
		c.Request().Context(),
		id,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "árvore removida com sucesso",
	})
}

func (h *TreeHandler) ListPotentialRiskTree(c echo.Context) error {
	tree, err := h.service.ListPotencialRiskTree(
		c.Request().Context(),
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, tree)
}
