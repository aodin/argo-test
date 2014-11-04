package main

import (
	dbsql "database/sql"

	sql "github.com/aodin/aspect"
	_ "github.com/aodin/aspect/postgres"

	db "github.com/aodin/argo-test/db"
)

// Up is executed when this migration is applied
func Up_20141104121456(txn *dbsql.Tx) {
	conn := sql.WrapTx(txn, sql.MustGetDialect("postgres"))

	// Insert a company
	company := db.Company{
		Name: "Test Company",
	}

	industries := []db.Industry{
		{
			Name:  "Plastics",
			About: "They're the future",
		},
		{
			Name:  "ハローキティ",
			About: "KAWAII!!!!",
		},
	}

	stmt := sql.Insert(
		db.Companies.C["name"],
	).Values(company)
	if _, err := conn.Execute(stmt); err != nil {
		panic(err)
	}
	stmt = sql.Insert(
		db.Industries.C["name"],
		db.Industries.C["about"],
	).Values(industries)
	if _, err := conn.Execute(stmt); err != nil {
		panic(err)
	}

	// TODO Set auto-increment
}

// Down is executed when this migration is rolled back
func Down_20141104121456(txn *dbsql.Tx) {
	conn := sql.WrapTx(txn, sql.MustGetDialect("postgres"))

	// Delete all companies and industries!
	if _, err := conn.Execute(db.Companies.Delete()); err != nil {
		panic(err)
	}
	if _, err := conn.Execute(db.Industries.Delete()); err != nil {
		panic(err)
	}

	// TODO Reset auto-increment

}
