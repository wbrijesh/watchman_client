package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func WriteLog(log LoggerProps) {
	file, _ := os.OpenFile("log.json", os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()

	jsonBytes, err := json.Marshal(log)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	jsonString := string(jsonBytes)

	fmt.Fprintf(file, "%v\n", jsonString)
}

func ReadLogs() ([]LoggerProps, error) {
	file, err := os.Open("log.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var logs []LoggerProps
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var log LoggerProps
		err := json.Unmarshal(scanner.Bytes(), &log)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return logs, nil
}

func TimeStampToHumanReadable(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	dateFormat := "2006-01-02 15:04:05"

	return t.Format(dateFormat)
}
