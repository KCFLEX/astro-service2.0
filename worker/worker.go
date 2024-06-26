package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/KCFLEX/astro-service2.0/cmd/fixtures"
	"github.com/KCFLEX/astro-service2.0/repository"
)

type fetchFixturesWorker struct {
	apiUrl string
	db     repository.Repository
}

func New(apiUrl string, db repository.Repository) *fetchFixturesWorker {
	return &fetchFixturesWorker{
		apiUrl: apiUrl,
		db:     db,
	}
}

func (ff *fetchFixturesWorker) Start() {
	ctx := context.Background()
	err := ff.db.CreateTable(ctx)
	if err != nil {
		log.Printf("failed to create table: %v", err)
		return
	}
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()
	client := &http.Client{}

	for {
		select {
		case <-ticker.C:

			req, err := http.NewRequest("GET", ff.apiUrl, nil)
			if err != nil {
				log.Printf("request initialization failed: %v", err)
				continue
			}

			resp, err := client.Do(req)
			if err != nil {
				log.Printf("request failed: %v", err)

				continue
			}

			var response fixtures.Response
			err = json.NewDecoder(resp.Body).Decode(&response)
			if err != nil {
				log.Print(err)
				return
			}

			for _, fixture := range response.Data {
				fmt.Printf("ID: %d, Name: %s, Starting At: %s, Result Info: %s\n", fixture.ID, fixture.Name, fixture.StartingAt, fixture.ResultInfo)
			}

			err = ff.db.InsertFixturesFromResponse(ctx, response)
			if err != nil {
				log.Printf("failed to save to db: %v", err)
				return
			}

		}

	}

}
