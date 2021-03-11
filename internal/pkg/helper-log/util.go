package helper_log

import (
	"os"
	"path/filepath"
)

func mkLogDir(logPath string) error {
	dir, _ := filepath.Split(logPath)
	if len(dir) > 0 {
		return os.MkdirAll(dir, os.ModePerm)
	}
	return nil
}
