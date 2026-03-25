package handler

import "github.com/labstack/echo/v4"

type TreeHandlerInterface interface {
	CreateTreeHandler(c echo.Context) error
	ListTreesHandler(c echo.Context) error   // ✅ ADD
	GetTreesByLocation(c echo.Context) error // ✅ ADD
	ListTreesByBoundingBox(c echo.Context) error
	GetTreeByID(c echo.Context) error
	UpdateTree(c echo.Context) error
	DeleteTree(c echo.Context) error
	ListPotentialRiskTree(c echo.Context) error
}
