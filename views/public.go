package views

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed public
var publicAssets embed.FS

// GetPublicAssetsFileSystem returns a http.FileSystem for the public assets so that
// we can embed them into the binary so it is self-contained.
func GetPublicAssetsFileSystem() http.FileSystem {
	fsys, err := fs.Sub(publicAssets, "public")
	if err != nil {
		panic(err)
	}
	return http.FS(fsys)
}
