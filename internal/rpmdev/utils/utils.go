package utils

import (
	"github.com/pkg/errors"
	"io"
	"os"
)

func Contains(l []string, v string) bool {
	for _, vv := range l {
		if vv == v {
			return true
		}
	}
	return false
}

func Cp(dst, src string) error {
	s, err := os.Open(src)
	if err != nil {
		return errors.WithStack(err)
	}
	defer s.Close()
	d, err := os.Create(dst)
	if err != nil {
		return errors.WithStack(err)
	}
	if _, err := io.Copy(d, s); err != nil {
		d.Close()
		return errors.WithStack(err)
	}
	return d.Close()
}

func ReadFile(src string) string {
	c, err := os.ReadFile(src)
	if err != nil {
		return ""
	}
	return string(c)
}
