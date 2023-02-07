package main

import (
	"fmt"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	logger := lumberjack.NewLogger(lumberjack.FileLoggerInfo{
		FileName: "rotate.log",
		FilePath: ".",
	}, lumberjack.WithRotate("@every 0h0m2s"), lumberjack.WithSplitFileByDate("@every 0h0m5s"))
	for {
		time.Sleep(time.Second)
		logger.Write([]byte(fmt.Sprintf("%s\n", time.Now().String())))
	}

}
