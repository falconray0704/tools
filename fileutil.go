package tools

import (
	"os"
)

func FileExist(file string) bool {
	_, err := os.Stat(filename)
        return err == nil || os.IsExist(err)
}
