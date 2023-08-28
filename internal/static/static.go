package static

import (
	"embed"
	"io/fs"
)

// Build VueJS app
//go:generate npm run build --prefix ../../vue-project

//go:embed all:resources
var res embed.FS

func Resources() (fs.FS, error) {
	return fs.Sub(res, "resources")
}

func ResourcesAll() embed.FS {
	return res
}
