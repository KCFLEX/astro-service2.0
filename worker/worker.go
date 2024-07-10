package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/KCFLEX/astro-service2.0/cmd/fixtures"
)

//go:generate mockgen -destination=worker/mocks/mocks.go -source=worker/worker.go . Repository
type Repository interface {
	CreateTable(ctx context.Context) error
	InsertFixtures(ctx context.Context, footballData fixtures.Fixture) (int, error)
	InsertFixturesFromResponse(ctx context.Context, response fixtures.Response) error
}

type fetchFixturesWorker struct {
	apiUrl string
	db     Repository
}

func New(apiUrl string, db Repository) *fetchFixturesWorker {
	return &fetchFixturesWorker{
		apiUrl: apiUrl,
		db:     db,
	}
}

func (ff *fetchFixturesWorker) Start() error {
	ctx := context.Background()
	err := ff.db.CreateTable(ctx)
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
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

			if resp.StatusCode != http.StatusOK {

				return fmt.Errorf("%v", resp.StatusCode)
			}

			var response fixtures.Response
			err = json.NewDecoder(resp.Body).Decode(&response)
			if err != nil {

				return fmt.Errorf("%v", err)
			}

			// for _, fixture := range response.Data {
			// 	fmt.Printf("ID: %d, Name: %s, Starting At: %s, Result Info: %s\n", fixture.ID, fixture.Name, fixture.StartingAt, fixture.ResultInfo)
			// }

			err = ff.db.InsertFixturesFromResponse(ctx, response)
			if err != nil {

				return fmt.Errorf("failed to save to db: %v", err)
			}

		}

	}

}
