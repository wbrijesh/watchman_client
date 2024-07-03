package watchman_client

import (
	"fmt"
	"os"
	"time"
)

func BackgroundLogSync(seconds int) {
	ticker := time.NewTicker(time.Duration(seconds) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if IsLogFileEmpty() {
			fmt.Println("Log file is empty")
		} else {
			fmt.Println("Log file is not empty, but sync function is not built yet")
		}
	}
}

func IsLogFileEmpty() bool {
	file, _ := os.Open("log.json")
	defer file.Close()
	fileInfo, _ := file.Stat()
	return fileInfo.Size() == 0
}
