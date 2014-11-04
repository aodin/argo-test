package main

import (
	dbsql "database/sql"

	sql "github.com/aodin/aspect"
	_ "github.com/aodin/aspect/postgres"

	db "github.com/aodin/argo-test/db"
)

// Up is executed when this migration is applied
func Up_20141103232645(txn *dbsql.Tx) {
	conn := sql.WrapTx(txn, sql.MustGetDialect("postgres"))
	if _, err := conn.Execute(db.Companies.Create()); err != nil {
		panic(err)
	}
	if _, err := conn.Execute(db.Industries.Create()); err != nil {
		panic(err)
	}
	if _, err := conn.Execute(db.CompanyIndustries.Create()); err != nil {
		panic(err)
	}
}

// Down is executed when this migration is rolled back
func Down_20141103232645(txn *dbsql.Tx) {
	conn := sql.WrapTx(txn, sql.MustGetDialect("postgres"))
	if _, err := conn.Execute(db.CompanyIndustries.Drop().IfExists()); err != nil {
		panic(err)
	}
	if _, err := conn.Execute(db.Industries.Drop().IfExists()); err != nil {
		panic(err)
	}
	if _, err := conn.Execute(db.Companies.Drop().IfExists()); err != nil {
		panic(err)
	}
}
