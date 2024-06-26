package main

import (
	"log"

	"github.com/KCFLEX/astro-service2.0/repository"
	"github.com/KCFLEX/astro-service2.0/worker"
)

// func Worker(apiUrl string) (fixtures.Response, error) {
// 	resp, err := http.Get(apiUrl)
// 	if err != nil {
// 		return fixtures.Response{}, fmt.Errorf("error fetching data from API: %w", err)
// 	}

// 	respppp, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return fixtures.Response{}, err
// 	}
// 	//fmt.Println(string(respppp))
// 	var response fixtures.Response
// 	err = json.Unmarshal(respppp, &response)
// 	if err != nil {
// 		return fixtures.Response{}, fmt.Errorf("failed to decode JSON response: %w", err)
// 	}
// 	for _, fixture := range response.Data {
// 		fmt.Printf("ID: %d, Name: %s, Starting At: %s, Result Info: %s\n", fixture.ID, fixture.Name, fixture.StartingAt, fixture.ResultInfo)
// 	}
// 	return response, nil
// }

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

	// for {
	// 	fmt.Println(Worker(url))
	// 	time.Sleep(24 * time.Hour)
	// }

}
