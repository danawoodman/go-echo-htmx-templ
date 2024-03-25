# go-echo-htmx-templ

> A basic started template for a Go web app with Echo, htmx, TailwindCSS and Alpine.js and some other goodies

Tools:

- [Go](https://golang.org/)
- [Echo](https://echo.labstack.com/) - golang web framework
- [templ]() - HTML templating
- [htmx](https://htmx.org/) - hypermedia controls
- [Tailwind CSS](https://tailwindcss.com/) - styling
  - [DaisyUI](https://daisyui.com/) - Tailwind UI elements
- [Alpine.js](https://alpinejs.dev/) - UI interactions
- [Charmbracelet's log](https://charm.sh/) - logging
- [cng](https://github.com/danawoodman/cng) - file watcher

## Setup

Install the latest version of Go and run the following commands:

```bash
# download go dependencies
go mod download

# install templ CLI if you don't already have it
go install github.com/a-h/templ/cmd/templ@latest

# install tailwind + other deps
npm install
```

## Development

Now you can run the dev server with `make` or `make dev` and open <http://localhost:7878> in your browser.

Build the binary with `make build` and run it with `./bin/server`.

## License

MIT
