package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/CSystem/gcuu/initializer/config"
	// Register MySQL driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// InitConnection - opening a database connection and save the reference.
func InitConnection(conf *config.Configuration) *gorm.DB {
	var err error

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		conf.DB.Username,
		conf.DB.Password,
		conf.DB.Host,
		conf.DB.Port,
		conf.DB.Name,
		conf.DB.Charset)

	// example dbURI: "monty:some_pass@tcp(192.168.50.193:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(conf.DB.Adapter, dbURI)
	if err != nil {
		fmt.Printf("mysql connect error %v\n", err)
		panic("failed to connect database")
	}

	if db.Error != nil {
		fmt.Printf("database error %v", db.Error)
	}

	// Enable or Disable SQL Logger
	db.LogMode(conf.DebugSQL)

	return db
}
