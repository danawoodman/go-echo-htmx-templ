package handlers

import (
	"net/http"
	"os"
	"strings"

	"github.com/danawoodman/go-echo-htmx-templ/views"
	"github.com/danawoodman/gocrazy"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	// attach any initialization fields here you need...
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() error {

	app := echo.New()

	// Middleware
	app.Use(middleware.Gzip())
	// add whatever else you need here...

	// Handle static assets:
	assetHandler := http.FileServer(views.GetPublicAssetsFileSystem())
	app.GET("/public/*", echo.WrapHandler(http.StripPrefix("/public/", assetHandler)))

	// Routes
	app.GET("/", Home)
	app.POST("/count/inc", CountIncrement)
	app.POST("/count/dec", CountDecrement)

	return app.Start(getHost())
}

// Construct the host string from environment variables.
func getHost() string {
	// An empty hostname binds to all interfaces:
	hostname := os.Getenv("HOSTNAME")
	port := gocrazy.Getenv("PORT", "7878")
	return strings.Join([]string{hostname, port}, ":")
}
