package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/scukonick/randomtiger/internal/app/handlers/top"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/scukonick/randomtiger/internal/app"
	"github.com/scukonick/randomtiger/internal/app/db"
	"github.com/scukonick/randomtiger/internal/app/handlers/enlarger"
	"github.com/scukonick/randomtiger/internal/app/handlers/getter"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	log.Print("starting app")

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	connConfig, err := pgx.ParseConfig("user=tiger password=tiger host=127.0.0.1 port=5432 dbname=tiger sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	nativeDB := stdlib.OpenDB(*connConfig)
	database := sqlx.NewDb(nativeDB, "pgx")

	err = database.Ping()
	if err != nil {
		panic(err)
	}

	//bot.Debug = true
	log.Printf("autorized with: %s", bot.Self.UserName)

	storage := db.NewStorage(database)

	getHandler := getter.NewHandler(bot, storage)
	growHandler := enlarger.NewHandler(bot, storage)
	topHandler := top.NewHandler(bot, storage)

	router := app.NewRouter(bot)
	router.AddCmdRoute("get", getHandler)
	router.AddCmdRoute("enlarge", growHandler)
	router.AddCmdRoute("top", topHandler)

	a := app.NewApp(bot, router)
	a.Run()
}
