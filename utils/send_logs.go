package utils

import "fmt"

func SendLogsToServer() {
	logs, err := ReadLogs()
	if err != nil {
		fmt.Println("Error reading logs:", err)
		return
	}

	for _, log := range logs {
		fmt.Println(log)
	}
}
