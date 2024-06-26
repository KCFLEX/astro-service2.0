package main

import (
	"log"

	"github.com/KCFLEX/astro-service2.0/repository"
	"github.com/KCFLEX/astro-service2.0/worker"
)

func main() {
	url := "https://api.sportmonks.com/v3/football/fixtures?api_token=KExOTHUa9KsM9IDMbCFwKc8maSrXoOGQ0fbZUAfnKvuqWElPqNuq3D6DMa0R"
	connStr := "user=postgres password=password dbname=postgres host=localhost port=5432 sslmode=disable"
	repo, err := repository.New(connStr)
	if err != nil {
		log.Fatalf("could not create repository: %v", err)
	}
	defer repo.Close()

	footballWorker := worker.New(url, *repo)
	footballWorker.Start()

}
