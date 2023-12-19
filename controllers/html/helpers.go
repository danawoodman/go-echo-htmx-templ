package html

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func render(ctx *fiber.Ctx, component templ.Component) error {
	return adaptor.HTTPHandler(templ.Handler(component))(ctx)
}

// func GetCurrentUser(ctx context.Context) *models.User {
// 	if user, ok := ctx.Locals("currentUser").(*models.User); ok {
// 		return user
// 	}
// 	return nil
// }
