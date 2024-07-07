package utils

import (
	"fmt"
	"os"
)

func IsLogFileEmpty() bool {
	file, _ := os.Open("log.json")
	defer file.Close()
	fileInfo, _ := file.Stat()
	return fileInfo.Size() == 0
}

func PrintLog(log LoggerProps) {
	fmt.Printf(
		"%v [%v] %v: %v\n",
		TimeStampToHumanReadable(log.Timestamp),
		log.Level,
		log.Subject,
		log.Message,
	)
}
