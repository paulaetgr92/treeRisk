package handler

import (
	"arvore/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TreeExternalHandler struct {
	service *service.TreeExternalService
}

func NewTreeExternalHandler(s *service.TreeExternalService) *TreeExternalHandler {
	return &TreeExternalHandler{service: s}
}

func (h *TreeExternalHandler) GetTrees(c echo.Context) error {
	trees, err := h.service.GetTrees()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao buscar árvores",
		})
	}

	return c.JSON(http.StatusOK, trees)
}
func (h *TreeExternalHandler) GetNearbyTrees(c echo.Context) error {
	lat := c.QueryParam("lat")
	lng := c.QueryParam("lng")
	distance := c.QueryParam("distance")

	if lat == "" || lng == "" {
		return c.JSON(400, map[string]string{
			"error": "lat e lng são obrigatórios",
		})
	}

	// default distance
	if distance == "" {
		distance = "100"
	}

	trees, err := h.service.GetNearbyTrees(lat, lng, distance)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, trees)
}
