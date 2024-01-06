package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// render creates a compatible interface for rendering templ components
// using the Echo web framework.
func render(ctx echo.Context, cmp templ.Component) error {
	return cmp.Render(ctx.Request().Context(), ctx.Response())
}
