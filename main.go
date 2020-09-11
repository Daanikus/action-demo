package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

var (
	ctx context.Context
	db  *sql.DB
)

type Bar struct {
	ID    int    `json:"id"`
	Name  string `yaml:"name"`
	Thing string `json:"name" yaml:"name"`
}

func main() {
	Thing()
	foo("Hello", context.Background())

}

func Thing() {
	t := []string{}
	fmt.Println(t)
}

func ExampleDB_QueryRowContext() {
	id := 123
	var username string
	var created time.Time
	err := db.QueryRowContext(ctx, "SELECT username, created_at FROM users WHERE id=?", id).Scan(&username, &created)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with id %d\n", id)
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:
		log.Printf("username is %q, account created on %s\n", username, created)
	}
}

func ExampleDB_QueryRowContextBadOne() {
	id := 123
	var username string
	var created time.Time
	err := db.QueryRowContext(ctx, "SELECT username, created_at FROM users WHERE id=?", id).Scan(&username, &created)
	switch {
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:
		log.Printf("username is %q, account created on %s\n", username, created)
	}
}

func ExampleDB_QueryRowContextBadTwo() {
	id := 123
	var username string
	var created time.Time
	err := db.QueryRowContext(ctx, "SELECT username, created_at FROM users WHERE id=?", id).Scan(&username, &created)
	if err != nil {
		log.Fatal(err)
	}
}
