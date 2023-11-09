package gwebz

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _sqlPool *gorm.DB

func GetDB() *gorm.DB {
	return _sqlPool
}

func SetDB(db *gorm.DB) {
	_sqlPool = db
}

// init sql pool
func initDB(driverName, host, port, database, username, password, charset string) (*gorm.DB, error) {

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	// initialize the connection pool
	return gorm.Open(mysql.New(mysql.Config{DriverName: driverName, DSN: args}), &gorm.Config{})

}

// global mode delivery
func InitGlobalDB(driverName, host, port, database, username, password, charset string) error {
	dbPool, err := initDB(driverName, host, port, database, username, password, charset)
	if err != nil {
		return err
	}
	SetDB(dbPool)
	return nil
}

// passed as a parameter
func NewDB(driverName, host, port, database, username, password, charset string) (*gorm.DB, error) {
	return initDB(driverName, host, port, database, username, password, charset)
}

func SetDBPool(db *gorm.DB, maxIdleConns, maxOpenConn, connMaxLifetime int) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(maxIdleConns)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(maxOpenConn)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)
	return nil
}
