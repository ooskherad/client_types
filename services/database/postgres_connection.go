package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var user string
var password string
var db string
var host string
var port string
var dbConn *gorm.DB

func init() {
	user = "postgres"
	password = "1234"
	db = "postgres"
	host = "192.168.1.10"
	port = "1234"
}

func GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s ", host, user, password, db, port)
}

func CreateDBConnection() error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  GetDSN(),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{NowFunc: func() time.Time {
		utc, _ := time.LoadLocation("Asia/Tehran")
		return time.Now().In(utc)
	}})

	if err != nil {
		log.Println("Error occurred while connecting with the database")
	}

	// Create the connection pool

	sqlDB, err := db.DB()

	sqlDB.SetConnMaxIdleTime(time.Minute * 5)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	dbConn = db
	return err
}

func GetDatabaseConnection() (*gorm.DB, error) {
	sqlDB, err := dbConn.DB()
	if err != nil {
		return dbConn, err
	}
	if err := sqlDB.Ping(); err != nil {
		return dbConn, err
	}
	return dbConn, nil
}

func Init() {
	err := CreateDBConnection()
	if err != nil {
		log.Println(err.Error())
	}
}
