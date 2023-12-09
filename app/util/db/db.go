package db

import (
	"fmt"
	"myproject/app/util/env"
	"time"

	"github.com/pkg/errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConnectOption struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
	Args     string
}

func NewDBConnect(dbNames ...string) ConnectOption {
	var dbName string
	if dbNames != nil {
		dbName = dbNames[0]
	}
	return ConnectOption{
		Host:     env.DBHost(),
		Port:     env.DBPort(),
		DBName:   dbName,
		User:     env.DBUser(),
		Password: env.DBPassword(),
		Args:     "charset=utf8mb4&parseTime=true",
	}
}

func newConnection(dbNames ...string) (*gorm.DB, error) {
	db := NewDBConnect(dbNames...)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		db.User, db.Password, db.Host, db.Port, db.DBName, db.Args)

	conn, err := gorm.Open(mysql.Open(dsn),
		&gorm.Config{})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return conn, nil
}

func Open(dbNames ...string) (*gorm.DB, error) {
	dbName := NewDBConnect(dbNames...).DBName
	conn, connErr := newConnection(dbName)
	if connErr != nil {
		return nil, connErr
	}

	// Connection pool
	sqlDb, dbErr := conn.DB()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)
	if dbErr != nil {
		return nil, dbErr
	}

	return conn, nil
}

func Close(conn *gorm.DB) {
	sqlDB, err := conn.DB()
	if err == nil {
		_ = sqlDB.Close()
	}
}
