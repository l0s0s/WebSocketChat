package client

import (
	"time"
)

// Message represents a single message.
type Message struct {
	Name    string
	Message string
	When    time.Time
}
