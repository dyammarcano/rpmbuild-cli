package static

import (
	"embed"
	"fmt"
)

// Build VueJS app
//go:generate npm run build --prefix ../../vue-project

//go:embed all:resources
var res embed.FS

func GetRsouce(name string) ([]byte, error) {
	path := fmt.Sprintf("resources/%s", name)
	return res.ReadFile(path)
}
