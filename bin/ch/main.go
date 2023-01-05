package main

import (
	"context"
	"crypto/tls"
	"log"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/davecgh/go-spew/spew"
)

var (
	hostname string          = "fylr7mz9tk.us-east-2.aws.clickhouse.cloud:9440"
	username string          = "default"
	database string          = "default"
	password string          = "LNcihzmRCNQb"
	ctx      context.Context = context.Background()
)

func main() {

	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{hostname},
		Auth: clickhouse.Auth{
			Database: database,
			Username: username,
			Password: password,
		},
		TLS: &tls.Config{},
	})

	if err != nil {
		log.Fatal("unable to create connection to ", hostname)
	}

	defer closeDb(conn)
	ping(conn)

	query(conn, "select 1, 2, 'bob'")
}

func ping(conn clickhouse.Conn) {
	if err := conn.Ping(ctx); err != nil {
		log.Fatal("ping error: ", err)
	} else {
		log.Print("able to ping ", hostname)
	}
}

func query(conn clickhouse.Conn, query string) {
	rows, err := conn.Query(ctx, query)
	if err != nil {
		log.Print("unable to execute query: ", query)
	}
	spew.Dump(rows.Columns())
}

// Close out the connection
func closeDb(conn clickhouse.Conn) {
	log.Print("closing database: ", hostname)
	conn.Close()

}
