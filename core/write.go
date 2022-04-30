package core

import (
	"io"
	"os"
)

func (Lyric *lyric) WriteLyric(fullPath string, content string) bool {
	f, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return false
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	_, err = io.WriteString(f, content)
	if err != nil {
		return false
	}
	return true
}
