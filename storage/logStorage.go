package storage

import (
	"github.com/hopperteam/hopper-monitor/types"
	"github.com/jinzhu/gorm"
)


type LogStorage struct {
	db *gorm.DB
}

func (storage *LogStorage) StoreLogEntry(entry *types.LogEntry) {
	storage.db.Create(entry)
}

func (storage *LogStorage) GetLogEntries(filter *types.LogFilter) []types.LogEntry {
	limit := 100
	var entries []types.LogEntry
	storage.db.Limit(limit).Find(&entries)
	return entries
}
