package watchman_client

import (
	"time"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	return uuid.New().String()
}

func TimeStampToHumanReadable(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	dateFormat := "2006-01-02 15:04:05"

	return t.Format(dateFormat)
}
