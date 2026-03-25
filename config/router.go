package config

import (
	"arvore/internal/handler"

	"github.com/labstack/echo/v4"
)

func RegisterTreeRoutes(
	e *echo.Echo,
	treeHandler *handler.TreeHandler,
	treeExternalHandler *handler.TreeExternalHandler,
) {

	// ===========================
	// 🌳 GRUPO /trees
	// ===========================
	trees := e.Group("/trees")

	// ===========================
	// 🌳 CRUD BÁSICO
	// ===========================

	// ➕ Criar árvore
	trees.POST("", treeHandler.CreateTreeHandler)

	// 📋 Listar todas
	trees.GET("", treeHandler.ListTreesHandler)

	// 📍 Buscar por localização (🔥 ESSENCIAL pro mapa)
	trees.GET("/near", treeHandler.GetTreesByLocation)

	// 📦 Bounding box (opcional)
	trees.GET("/bbox", treeHandler.ListTreesByBoundingBox)

	// 🔎 Buscar por ID
	trees.GET("/:id", treeHandler.GetTreeByID)

	// ✏️ Atualizar
	trees.PUT("/:id", treeHandler.UpdateTree)

	// ❌ Deletar
	trees.DELETE("/:id", treeHandler.DeleteTree)

	// ⚠️ Árvores com risco
	trees.GET("/risk", treeHandler.ListPotentialRiskTree)

	// ===========================
	// 🌍 API EXTERNA (SEATTLE)
	// ===========================

	trees.GET("/external/nearby", treeExternalHandler.GetNearbyTrees)
}
