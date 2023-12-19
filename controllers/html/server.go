package html

import (
	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
)

type HTMLController struct {
	host string
	// app  *fiber.App
}

// Creates a new server instance
func NewHTMLController(host string) *HTMLController {
	return &HTMLController{host: host}
}

func (c *HTMLController) Start() error {
	// env := os.Getenv("ENV")

	// engine := html.New(views.HTMLTemplatePath, views.HTMLTemplateExt)

	app := echo.New()

	// Middleware
	// c.app.Use(compress.New())
	app.Static("/public", "public")
	// app.Use("/public", filesystem.New(
	// 	filesystem.Config{
	// 		Root: http.FS(views.HTMLPublicAssets),

	// 		// This must match the embedded file path of
	// 		// theHTMLPublicAssets value
	// 		PathPrefix: "/html/public",

	// 		// MaxAge: 60, // in seconds
	// 	},
	// ))
	// // c.app.Use(middleware.UserToSessionMiddleware)

	// Routes
	app.Get("/", Home)

	log.Print("Starting server on", c.host)

	return app.Listen(c.host)
	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	// e.Logger.Fatal(e.Start(":1323"))
}
