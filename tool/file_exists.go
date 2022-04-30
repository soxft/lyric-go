package tool

import "os"

func FileExists(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}
