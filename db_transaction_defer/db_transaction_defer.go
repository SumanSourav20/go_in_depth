package main

import (
	"context"
	"database/sql"
	"fmt"

	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "192.168.156.76"
	port     = 5432
	user     = "postgres"
	password = "suman"
	dbname   = "pg4e"
)

func doSomeTransaction(ctx context.Context, db *sql.DB) (err error) {
	tx, err := db.BeginTx(ctx, nil)

	defer func() {
		if err == nil {
			fmt.Println("commited transaction.")
			tx.Commit()
		}
		if err != nil {
			tx.Rollback()
		}
	}()

	rows, err := tx.QueryContext(ctx, "SELECT * FROM track_raw;")
	if err != nil {
		return nil
	}

	cols, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("failed to get columns: %w", err)
	}
	log.Printf("Table columns: %v", cols)

	// Create a slice of interfaces to hold the column values
	values := make([]interface{}, len(cols))
	scanArgs := make([]interface{}, len(cols))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Iterate over rows
	for rows.Next() {
		if err := rows.Scan(scanArgs...); err != nil {
			return fmt.Errorf("failed to scan row: %w", err)
		}

		// Print all column values
		for i, col := range cols {
			val := values[i]
			log.Printf("%s: %v", col, val)
		}
	}

	if err = rows.Err(); err != nil {
		return fmt.Errorf("error iterating rows: %w", err)
	}

	return nil
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Println(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Connecting...")

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

	ctx := context.Background()
	err = doSomeTransaction(ctx, db)

	if err != nil {
		fmt.Println("transaction error : ", err)

	}

	// defer rows.Close()

	// for rows.Next() {
	// 	var id int
	// 	var data string

	// 	if err := rows.Scan(&id, &data); err != nil {
	// 		log.Fatalf("Failed to scan row: %v", err)
	// 	}
	// 	fmt.Printf("Row: ID=%d, Data=%s\n", id, data)
	// }

	// if err = rows.Err(); err != nil {
	// 	log.Fatalf("Error iterating rows: %v", err)
	// }

	fmt.Println("Transaction succeeded")
}
