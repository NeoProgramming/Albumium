package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Application struct {
	db       *sqlx.DB
	srv      *http.Server
	errorLog *log.Logger
	config   Configuration
}

var App Application

func main() {
	fmt.Println("Hello")
	App.errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	LoadConfig()
	InitDatabase()

	InitWeb()
	go HandleWeb()

	interruptCh := make(chan os.Signal, 1)
	signal.Notify(interruptCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-interruptCh
		fmt.Println("\nReceived interrupt signal. Exiting...")
		CloseDatabase()
		fmt.Println("Database closed")
		os.Exit(0)
	}()

	for {
		time.Sleep(time.Second)
	}

	//Scan()

	fmt.Println("good bye")
}
