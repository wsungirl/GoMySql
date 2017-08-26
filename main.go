package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/wsungirl/GoMySql/db"
	"github.com/wsungirl/GoMySql/handler"

	"github.com/go-sql-driver/mysql"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type config struct {
	Mysql mysql.Config `json:"mysql"`
}

func loadConfig(filename string) *config {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	}

	var cfg config
	if err = json.Unmarshal(data, &cfg); err != nil {
		return nil
	}

	return &cfg
}

func main() {
	cfg := loadConfig("config.json")

	db, err := db.InitDB("mysql", cfg.Mysql.FormatDSN())
	if err != nil {
		log.Println(errors.New("Can't connect to DB"))
	}

	router := handler.Setup(db)

	http.Handle("/", router)

	if err := http.ListenAndServe(":6060", nil); err != nil {
		log.Println(err)
	}
}
