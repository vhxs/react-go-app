package server

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx/v4"
)

func GetAllComments() string {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), "postgres://db-user:shittypassword@db:5432/testdb")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT comment FROM comments")
	if err != nil {
		return "query failed: " + err.Error()
	}

	comments := make([]string, 0, 10)
	for rows.Next() {
		var comment string
		rows.Scan(&comment)
		comments = append(comments, comment)
	}

	return strings.Join(comments, ",")
}

func MakeComment(comment string) {
	// TODO: make a struct that keeps a pgx connection
	conn, err := pgx.Connect(context.Background(), "postgres://db-user:shittypassword@db:5432/testdb")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), "INSERT INTO comments VALUES ($1)", comment)
	if err != nil {
		fmt.Println("query failed: " + err.Error())
		return
	}
}