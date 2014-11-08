package main

import (
	dbsql "database/sql"

	sql "github.com/aodin/aspect"
	_ "github.com/aodin/aspect/postgres"

	db "github.com/aodin/argo-test/db"
)

// Up is executed when this migration is applied
func Up_20141107092004(txn *dbsql.Tx) {
	conn := sql.WrapTx(txn, sql.MustGetDialect("postgres"))

	if _, err := conn.Execute(db.CompanyContacts.Create()); err != nil {
		panic(err)
	}

	// Insert contacts
	contacts := []db.CompanyContact{
		{CompanyID: 1, Key: "phone", Value: "555-555-5555"},
		{CompanyID: 1, Key: "email", Value: "bigco@example.com"},
		{CompanyID: 2, Key: "instagram", Value: "@startuprr"},
	}

	stmt := sql.Insert(
		db.CompanyContacts.C["company_id"],
		db.CompanyContacts.C["key"],
		db.CompanyContacts.C["value"],
	).Values(contacts)
	if _, err := conn.Execute(stmt); err != nil {
		panic(err)
	}
}

// Down is executed when this migration is rolled back
func Down_20141107092004(txn *dbsql.Tx) {
	conn := sql.WrapTx(txn, sql.MustGetDialect("postgres"))

	if _, err := conn.Execute(db.CompanyContacts.Delete()); err != nil {
		panic(err)
	}
}
