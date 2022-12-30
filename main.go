package main

import (
	"self-payrol/config"
	"sync"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	_ "gorm.io/driver/postgres"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Infof(".env is not loaded properly")
	}
	log.Infof("read .env from file")

	config := config.NewConfig()
	serv := InitServer(config)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		serv.Run()
	}()

	wg.Wait()
}
