package cli

import (
	"os"
	"time"
)

func CreateMigration(baseDir string) bool {
	err := os.MkdirAll(baseDir+"/"+time.Now().String(), 0o750)
	return err == nil
}
