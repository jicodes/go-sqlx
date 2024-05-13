package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func MustGetEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return v
}

var ConnStr = MustGetEnv("CONNECTION_STRING")

func main() {
	// Create a new connection to the database
	db := sqlx.MustConnect("postgres", ConnStr)
	defer db.Close()
  var version string
  err:= db.QueryRow("SELECT version()").Scan(&version)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(version)
}
