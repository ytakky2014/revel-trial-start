package controllers

import (
	"log"
	"revel-trial-start/app/models"
	"strconv"

	"github.com/revel/revel"
)

type Hello struct {
	App
}

func (c Hello) Index(id int64) revel.Result {

	DbMap.Insert(&models.Hello{id, strconv.FormatInt(id, 10)})

	count, _ := DbMap.SelectInt("select count(*) from hello")
	log.Println(count)
	return c.Render(id, count)
}
