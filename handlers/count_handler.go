package handlers

import (
	"github.com/danawoodman/go-echo-htmx-templ/views"
	"github.com/labstack/echo/v4"
)

var Count = 0

// TODO: if you wanted progressive enhancement here,
// you could redirect the user to / if the request didn't
// come from htmx:

func CountIncrement(ctx echo.Context) error {
	Count++
	return render(ctx, views.Count(Count))
}

func CountDecrement(ctx echo.Context) error {
	Count--
	if Count < 0 {
		Count = 0
	}
	return render(ctx, views.Count(Count))
}
