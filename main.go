package main

import (
	"log"

	"github.com/strwrd/rest-scrapper/delivery/http"
	"github.com/strwrd/rest-scrapper/repository/mysql"
	"github.com/strwrd/rest-scrapper/usecase"
)

func main() {
	// Initial mysql repository object
	mysqlRepo, err := mysql.NewRepository()
	if err != nil {
		log.Fatal(err)
	}

	// Create usecase object
	ucase := usecase.NewUsecase(mysqlRepo)

	// Create server object
	server := http.NewHandler(ucase)

	// Start server
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
