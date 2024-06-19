package mysql

import (
	"dxj/configs"
	"dxj/internal/models"
	"dxj/internal/pkg/logger"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(config *configs.Mysql) *gorm.DB {
	var ormLogger gormLogger.Interface
	if gin.Mode() == "debug" {
		ormLogger = gormLogger.Default.LogMode(gormLogger.Info)
	} else {
		ormLogger = gormLogger.Default
	}
	logger := logger.LogrusLogger
	address := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.UserName,
		config.PassWord,
		config.Host,
		config.Port,
		config.DataBase,
	)
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               address,
		DefaultStringSize: 256, // default size for string fields
	}), &gorm.Config{
		//打印SQL
		Logger: ormLogger,
	})

	if err != nil {
		panic(`😫: Connected failed, check your Mysql with ` + address)
	}
	// Migrate the schema
	migrateErr := db.AutoMigrate(&models.Category{}, &models.User{}, &models.Article{})
	if migrateErr != nil {
		panic(`😫: Auto migrate failed, check your Mysql with ` + address)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  //设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) //打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	//export DB
	DB = db
	logger.Printf(`🍟: Successfully connected to Mysql at ` + address)
	return db
}
