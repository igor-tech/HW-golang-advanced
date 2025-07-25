package db

import (
	"order/api/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(conf.DbConfig.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &Db{DB: db}
}
