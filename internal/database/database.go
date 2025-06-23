//database/database.go

package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	DB            *pgxpool.Pool
	DB_CONNECTION string = "postgres://postgres:password@localhost:5432/app"
)

const (
	maxRetries    = 10
	retryInterval = 5 * time.Second
)

func InitDB() {
	var err error
	DB, err = ConnectDbWithRetry(DB_CONNECTION)

	if err != nil {
		fmt.Println(err)
	}
}

func ConnectDbWithRetry(connString string) (*pgxpool.Pool, error) {
	var err error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		// This is the important line, everything else is gravy!
		DB, err = pgxpool.New(context.Background(), connString)
		if err == nil {
			// Test the connection
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			err = DB.Ping(ctx)
			if err == nil {
				log.Printf("Connected to database on attempt %d", attempt)
				return DB, nil
			}

			DB.Close() // Clean up failed pool
		}

		log.Printf("Attempt %d: Failed to connect to database: %v", attempt, err)

		if attempt < maxRetries {
			time.Sleep(retryInterval)
		}
	}

	return nil, fmt.Errorf("could not connect to database after %d attempts: %w", maxRetries, err)
}

func InsertCars(brand string, model string, year int) {
	var err error
	insertSQL := `INSERT INTO cars (brand, model,year) VALUES ($1, $2, $3);`
	_, err = DB.Exec(context.Background(), insertSQL, brand, model, year)
	if err != nil {
		fmt.Println(err)
	}
}

//func DeleteCars() {
//	deleteSQL := `DELETE FROM cars WHERE Brand="Nissan"`
//	rows, err := DB.Query(context.Background(), deleteSQL)
//}

func GetCars() {
	querySQL := `SELECT * FROM cars`
	rows, err := DB.Query(context.Background(), querySQL)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var carList []car

	for rows.Next() {
		var c car
		err2 := rows.Scan(&c.Brand, &c.Model, &c.Year)
		if err2 != nil {
			fmt.Println(err2)
		}
		carList = append(carList, c)
	}
	fmt.Println(carList)
}

func CloseDb() {
	DB.Close()
}

type car struct {
	Brand string
	Model string
	Year  string
}
