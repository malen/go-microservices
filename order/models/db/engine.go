package db

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"

	"aoisoft/order/models"
)

type Handler struct {
	DB *xorm.Engine
}

func Init(url string) Handler {
	var engine *xorm.Engine
	engine, err := xorm.NewEngine("sqlite3", url)
	if err != nil {
		log.Fatalln(err)
	}

	err = engine.Sync2(new(models.Order))
	if err != nil {
		log.Fatalln(err)
	}

	return Handler{engine}
}
