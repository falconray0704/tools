package tools

import (
	"os"
)

func FileExist(file string) bool {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
