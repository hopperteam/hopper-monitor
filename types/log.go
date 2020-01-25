package types

import (
	"github.com/jinzhu/gorm"
	"time"
)

type LogEntry struct {
	gorm.Model
	Instance string
	Severity uint8
	Component string
	Timestamp time.Time
	Message string
}

type LogFilter struct {
	Instance string
	Severity string
	Component string
	From time.Time
	To time.Time
	Limit int32
	Skip int32
}
