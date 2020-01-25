package storage

import (
	"github.com/hopperteam/hopper-monitor/types"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)


type ConfigProvider struct {
	LogStorage *LogStorage
}

func LoadConfig() *ConfigProvider {
	db, err := ConnectDb();
	if err != nil {
		panic(err)
	}

	return &ConfigProvider{
		&LogStorage{ db },
	}
}

	func ConnectDb() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&types.LogEntry{})

	return db, nil
}
