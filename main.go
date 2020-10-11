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
	var buf [17]byte
	fmt.Println(buf)
	buf = make([]byte, 17)
	fmt.Println(buf)
	
	err := "someerr"
	fmt.Println("sdf")
	if err != nil {
	}
	
	var b Bar
	name := "foo"
	b.Name = name
	fmt.Println(name)
	
	a := []string{"bar", "baz"}
	
	for _, s := range a {
		if s == "foo" {
			fmt.Println(s)
		}
	}
	
	for _, s := range a {
		if s == "foo" {
			// Catch
			fmt.Println(s)
			fmt.Println(s)
			fmt.Println(s)
			fmt.Println(s)
			fmt.Println(s)
		}
	}
	

}

func Thing() {
	id := 1
	query := "dfsdf"

	row := db.QueryRow(query, id)
	fmt.Println(row)
	row = db.QueryRowContext(ctx, query, id)
}

func ExampleDB_QueryRowContext() {
	defer fmt.Println("end")
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
