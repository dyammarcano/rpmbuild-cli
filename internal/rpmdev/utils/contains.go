package utils

type (
	fileItems []fileItem

	fileItem struct {
		Path string
		Type string
	}
)

func (f fileItems) Contains(path string) bool {
	for _, item := range f {
		if item.Path == path {
			return true
		}
	}
	return false
}
