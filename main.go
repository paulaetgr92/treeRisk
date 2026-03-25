package main

import (
	"arvore/config"
	sqlc "arvore/db/sqlc"
	"arvore/internal/handler"
	"arvore/internal/repository"
	"arvore/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	// 🔥 CORS (necessário pro frontend)
	e.Use(middleware.CORS())

	// ---------------------------
	// DATABASE
	// ---------------------------
	dbConn := config.NewDatabase()
	queries := sqlc.New(dbConn)

	baseRepo := repository.NewBaseRepository(queries)

	// ===========================
	// 🌳 REPOSITORY (BANCO)
	// ===========================
	treeRepo := repository.NewTreeRepository(baseRepo)

	// ===========================
	// 🌳 SERVICE
	// ===========================
	treeService := service.NewTreeService(treeRepo)

	// ===========================
	// 🌳 HANDLER
	// ===========================
	treeHandler := handler.NewTreeHandler(treeService)

	// ===========================
	// 🌍 API EXTERNA (OPCIONAL)
	// ===========================
	treeExternalRepo := repository.NewTreeExternalRepository()
	treeExternalService := service.NewTreeExternalService(treeExternalRepo)
	treeExternalHandler := handler.NewTreeExternalHandler(treeExternalService)

	// ===========================
	// 🌳 ROTAS
	// ===========================

	// 🔥 PRINCIPAL (frontend usa)
	e.GET("/trees", treeHandler.GetTreesByLocation)

	// CRUD
	e.POST("/trees", treeHandler.CreateTreeHandler)
	e.GET("/trees/:id", treeHandler.GetTreeByID)
	e.PUT("/trees", treeHandler.UpdateTree)
	e.DELETE("/trees/:id", treeHandler.DeleteTree)

	// opcional
	e.GET("/trees/all", treeHandler.ListTreesHandler)

	// 🌍 externa (se quiser usar Seattle)
	e.GET("/trees/external", treeExternalHandler.GetTrees)

	// ===========================

	e.Logger.Fatal(e.Start(":8080"))
}
