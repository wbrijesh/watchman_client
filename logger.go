package watchman_client

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/wbrijesh/watchman_client/utils"
)

func InitialiseWatchman() {
	if _, err := os.Stat("log.json"); os.IsNotExist(err) {
		fmt.Println("Creating log.json file")
		file, _ := os.Create("log.json")
		defer file.Close()
	} else {
		fmt.Println("log.json file already exists")
	}
}

func GenerateUUID() string {
	return uuid.New().String()
}

func Log(log utils.LoggerProps) {
	utils.PrintLog(log)
	utils.WriteLog(log)
}

func BackgroundLogSync(seconds int) {
	ticker := time.NewTicker(time.Duration(seconds) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if utils.IsLogFileEmpty() {
			fmt.Println("Log file is empty")
		} else {
			utils.SendLogsToServer()
		}
	}
}
