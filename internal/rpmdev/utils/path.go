package utils

import "path/filepath"

func EnsureAbsolute(path string, base string) string {
	if !filepath.IsAbs(path) {
		return filepath.Join(base, path)
	}
	return path
}
