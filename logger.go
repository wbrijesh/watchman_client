package watchman_client

import (
	"fmt"
	"os"
)

type LoggerProps struct {
	UUID       string `json:"uuid"`
	Level      string `json:"level"`
	Message    string `json:"message"`
	Subject    string `json:"subject"`
	SyncStatus string `json:"sync_status"`
	Timestamp  int64  `json:"timestamp"`
}

func LoggerInit() {
	CreateLogFile()
}

func CreateLogFile() {
	if _, err := os.Stat("log.json"); os.IsNotExist(err) {
		fmt.Println("Creating log.json file")
		file, _ := os.Create("log.json")
		defer file.Close()
	} else {
		fmt.Println("log.json file already exists")
	}
}

func WriteLog(log LoggerProps) {
	file, _ := os.OpenFile("log.json", os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()

	jsonStyleString := fmt.Sprintf("{\"uuid\": \"%v\", \"level\": \"%v\", \"message\": \"%v\", \"subject\": \"%v\", \"sync_status\": \"%v\", \"timestamp\": %v}", log.UUID, log.Level, log.Message, log.Subject, log.SyncStatus, log.Timestamp)
	fmt.Fprintf(file, "%v\n", jsonStyleString)
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

func Log(log LoggerProps) {
	PrintLog(log)
	WriteLog(log)
}
