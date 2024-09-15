package web

import (
	"github.com/labstack/echo/v4"
	"github.com/ppreeper/addictgo/web/pages"
)

func Index(c echo.Context) error {
	component := pages.Index()
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func Login(c echo.Context) error {
	component := pages.Login()
	return component.Render(c.Request().Context(), c.Response().Writer)
}
