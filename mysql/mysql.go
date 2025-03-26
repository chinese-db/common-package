package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	DB   *gorm.DB
	err  error
	once sync.Once
)

type Mysql struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
}

func NewMysql(user string, password string, host string, port int, database string) *Mysql {
	return &Mysql{
		User:     user,
		Password: password,
		Host:     host,
		Port:     port,
		Database: database,
	}
}

func (m *Mysql) MysqlINit() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", m.User, m.Password, m.Host, m.Port, m.Database)
	once.Do(func() {
		DB, err = gorm.Open(mysql.Open(dsn))
		if err != nil {
			fmt.Println("mysql error:", err)
			return
		}
		sqlDB, _ := DB.DB()

		sqlDB.SetMaxIdleConns(10)

		sqlDB.SetMaxOpenConns(100)

		sqlDB.SetConnMaxLifetime(time.Hour)

		fmt.Println("mysql init success!!!")
	})
	return DB
}
