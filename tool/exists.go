package tool

import (
	"os"
)

func FileExists(path string) bool {
	f, err := os.Lstat(path)

	// check if empty file
	if f.Size() == 0 {
		_ = os.Remove(path)
		return false
	}
	return !os.IsNotExist(err)
}
