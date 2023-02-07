package lumberjack

import (
	"fmt"
	"time"
)

// FileLoggerInfo ...
type FileLoggerInfo struct {
	FileName string
	FilePath string
	FileDate string
}

func (f FileLoggerInfo) path() string {
	if len(f.FileName) == 0 || len(f.FilePath) == 0 {
		return ""
	}
	if f.FileDate == "" {
		f.FileDate = time.Now().Format("2006-01-02")
	}

	return fmt.Sprintf("%s/%s/%s", f.FilePath, f.FileDate, f.FileName)
}
