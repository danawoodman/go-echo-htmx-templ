package html

import (
	"github.com/danawoodman/go-echo-htmx-templ/views"
	"github.com/labstack/echo/v4"
)

func Home(ctx *echo.Context) error {
	return render(ctx, views.Home("go-echo-htmx-templ"))
}
