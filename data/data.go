package data

import (
	"os"
	"path/filepath"
)

const DATA_BASE_RECORDS_LIMIT = 1_000

func Path() string {
	USER_HOME, _ := os.UserHomeDir()
	return filepath.Join(USER_HOME, "tt.db")
}
