package utils

type LoggerProps struct {
	UUID       string `json:"uuid"`
	Level      string `json:"level"`
	Message    string `json:"message"`
	Subject    string `json:"subject"`
	SyncStatus string `json:"sync_status"`
	Timestamp  int64  `json:"timestamp"`
}
