package config

import (
	"arvore/internal/handler"

	"github.com/labstack/echo/v4"
)

func RegisterTreeRoutes(e *echo.Echo, treeHandler *handler.TreeHandler) {

	trees := e.Group("/trees")

	trees.POST("", treeHandler.CreateTreeHandler)

	trees.GET("", treeHandler.ListTreesHandler)

	// 🔵 FILTRO POR ÁREA
	trees.GET("/bbox", treeHandler.ListTreesByBoundingBox)

	trees.GET("/:id", treeHandler.GetTreeByID)
	trees.PUT("", treeHandler.UpdateTree)
	trees.DELETE("/:id", treeHandler.DeleteTree)
	trees.GET("/risk", treeHandler.ListPotentialRiskTree)
}
