package main

import (
	"flag"
	"log"

	sql "github.com/aodin/aspect"
	_ "github.com/aodin/aspect/postgres"
	"github.com/aodin/volta/config"

	"github.com/aodin/argo-test/server"
)

func main() {
	var file string
	flag.StringVar(&file, "c", "./settings.json", "configuration file")
	flag.Parse()

	// Parse the given configuration
	c, err := config.ParseFile(file)
	if err != nil {
		log.Fatalf("argo-test: could not parse configuration: %s", err)
	}

	// Connect to the database
	db, err := sql.Connect(c.Database.Driver, c.Database.Credentials())
	if err != nil {
		log.Fatalf("argo-test: could not connect to the database: %s", err)
	}
	defer db.Close()

	// Configuration
	log.Fatal(server.New(c, db).ListenAndServe())
}
