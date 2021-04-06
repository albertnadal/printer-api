package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"strconv"
	"os"
	//"printer-api/api"
	"printer-api/router"
	"printer-api/models"
	"printer-api/managers"
	//"printer-api/middleware"
)

func main() {

	var config models.Configuration
	config_filename := "server.config.yaml"

	err := config.LoadConfiguration(config_filename)
	if err != nil {
			log.Fatalf("Error loading %s: %v", config_filename, err)
			os.Exit(1)
	}

	printerManager := managers.InitPrinterManager(config)
	httpHandler := router.InitRouter(printerManager, config)

	fmt.Println("Listening on port "+itoa(config.Server.Port))
	srv := &http.Server{
		Handler:      httpHandler,
		Addr:         ":"+itoa(config.Server.Port),
		WriteTimeout: config.Server.WriteTimeout * time.Second,
		ReadTimeout:  config.Server.ReadTimeout * time.Second,
		IdleTimeout:  config.Server.IdleTimeout * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func itoa(n int32) string {
    return strconv.FormatInt(int64(n), 10)
}
