package model

import (
	"time"
)

// Faq Struct for keeping the frequently asked questions under /admin/faq
type Faq struct {
	ID        int       `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Questions string    `json:"questions"`
}
