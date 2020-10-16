package env

import (
	"log"

	"github.com/axolotlteam/thunder/db/mongov2"
)

var (
	// 服務註冊 port
	PORT = "5000"
)

func Step() {
	connMongo()
}

func connMongo() {
	if err := mongov2.Con(mongov2.Config{
		Host:     "localhost:27017",
		User:     "",
		Password: "",
		Database: "local",
		AppName:  "",
	}); err != nil {
		panic(err.Error())
	}
	log.Println("mongoDB connection succeeded")
}
