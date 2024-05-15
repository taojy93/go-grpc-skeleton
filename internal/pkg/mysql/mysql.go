package mysql

import (
	"fmt"

	"go-grpc-skeleton/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormClient(cfg config.MySQLConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
