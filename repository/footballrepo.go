package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/KCFLEX/astro-service2.0/cmd/fixtures"
	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func DbConnect(conn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Printf("db validation failed incorrect parameters: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("connection verification failed: %v", err)
		return nil, err
	}

	return db, nil
}

func New(conn string) (*Repository, error) {
	connect, err := DbConnect(conn)
	if err != nil {
		log.Fatalf("db connection failed: %v", err)

	}
	return &Repository{
		db: connect,
	}, nil
}

func (repo *Repository) Close() error {
	return repo.db.Close()
}

func (repo *Repository) CreateTable(ctx context.Context) error {
	query := `CREATE TABLE IF NOT EXISTS fixtures (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
    starting_at TIMESTAMP,
    result_info TEXT,
    leg VARCHAR(10),
    details JSONB,
    length INTEGER,
    placeholder BOOLEAN,
    has_odds BOOLEAN,
    has_premium_odds BOOLEAN,
    starting_at_timestamp INTEGER)`

	_, err := repo.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	query2 := `CREATE TABLE IF NOT EXISTS carts (id SERIAL PRIMARY KEY);`
	_, err = repo.db.ExecContext(ctx, query2)
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) InsertFixtures(ctx context.Context, footballData fixtures.Fixture) (int, error) {
	query := `
	INSERT INTO fixtures (
		name, 
		starting_at, 
		result_info, 
		leg, 
		details, 
		length, 
		placeholder, 
		has_odds, 
		starting_at_timestamp
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	var id int
	err := repo.db.QueryRowContext(ctx, query, footballData.Name, footballData.StartingAt, footballData.ResultInfo, footballData.Leg, footballData.Details, footballData.Length, footballData.Placeholder, footballData.HasOdds, footballData.StartingAtTimestamp).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}

func (repo *Repository) InsertFixturesFromResponse(ctx context.Context, response fixtures.Response) error {
	for _, fixture := range response.Data {
		if _, err := repo.InsertFixtures(ctx, fixture); err != nil {
			return err
		}
	}
	return nil
}
