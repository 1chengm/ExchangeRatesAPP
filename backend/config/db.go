package config

import (
	"exchangeapp/global"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)	
  
  func initDB() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	db, err := gorm.Open(mysql.Open(AppConfig.Database.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database, %v", err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetConnMaxIdleTime(time.Duration(AppConfig.Database.MaxOpenConns))
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil{
		log.Fatalf("Error connecting to database2, %v", err)
	}
	global.Db = db
  }