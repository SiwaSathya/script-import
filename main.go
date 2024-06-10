package main

import (
	"log"
	"script-import/db"

	"github.com/joho/godotenv"
	"github.com/pandeptwidyaop/golog"
)

func main() {
	Init()

}

func Init() {
	InitEnv()
	InitSlack()
	InitDB()
}

func InitEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Println(".env file not found")
	}
}

func InitSlack() {
	golog.New()
}

func InitDB() {
	db.NewGormClient()
}
