package main

import (
	dbsql "database/sql"

	sql "github.com/aodin/aspect"
	"github.com/aodin/aspect/postgres"

	db "github.com/aodin/argo-test/db"
)

// Up is executed when this migration is applied
func Up_20141104121456(txn *dbsql.Tx) {
	conn := sql.WrapTx(txn, sql.MustGetDialect("postgres"))

	// Insert a company
	companies := []db.Company{
		{ID: 1, Name: "Big Co."},
		{ID: 2, Name: "Startuprr"},
	}

	industries := []db.Industry{
		{ID: 1, Name: "Plastics", About: "They're the future."},
		{ID: 2, Name: "ハローキティ", About: "KAWAII!!!!"},
		{ID: 3, Name: "Battletoads", About: "Do you have them?"},
		{ID: 4, Name: "Venture Capital", About: "Money, it's what I want."},
	}

	companyIndustries := []db.CompanyIndustry{
		{CompanyID: 1, IndustryID: 1},
		{CompanyID: 1, IndustryID: 4},
	}

	stmt := sql.Insert(
		db.Companies.C["id"],
		db.Companies.C["name"],
	).Values(companies)
	if _, err := conn.Execute(stmt); err != nil {
		panic(err)
	}
	stmt = sql.Insert(
		db.Industries.C["id"],
		db.Industries.C["name"],
		db.Industries.C["about"],
	).Values(industries)
	if _, err := conn.Execute(stmt); err != nil {
		panic(err)
	}
	stmt = sql.Insert(
		db.CompanyIndustries.C["company_id"],
		db.CompanyIndustries.C["industry_id"],
	).Values(companyIndustries)
	if _, err := conn.Execute(stmt); err != nil {
		panic(err)
	}

	// Set the auto-increments
	seqStmt := postgres.AlterSequence(
		postgres.Sequence("companies_id_seq"),
	).RestartWith(len(companies) + 1)
	if _, err := conn.Execute(seqStmt); err != nil {
		panic(err)
	}

	seqStmt = postgres.AlterSequence(
		postgres.Sequence("industries_id_seq"),
	).RestartWith(len(companies) + 1)
	if _, err := conn.Execute(seqStmt); err != nil {
		panic(err)
	}

	seqStmt = postgres.AlterSequence(
		postgres.Sequence("company_industries_id_seq"),
	).RestartWith(len(companyIndustries) + 1)
	if _, err := conn.Execute(seqStmt); err != nil {
		panic(err)
	}
}

// Down is executed when this migration is rolled back
func Down_20141104121456(txn *dbsql.Tx) {
	conn := sql.WrapTx(txn, sql.MustGetDialect("postgres"))

	// Delete all companies and industries!
	if _, err := conn.Execute(db.CompanyIndustries.Delete()); err != nil {
		panic(err)
	}
	if _, err := conn.Execute(db.Companies.Delete()); err != nil {
		panic(err)
	}
	if _, err := conn.Execute(db.Industries.Delete()); err != nil {
		panic(err)
	}

	// Reset auto-increment
	// Set the auto-increments
	seqStmt := postgres.AlterSequence(
		postgres.Sequence("companies_id_seq"),
	).RestartWith(1)
	if _, err := conn.Execute(seqStmt); err != nil {
		panic(err)
	}

	seqStmt = postgres.AlterSequence(
		postgres.Sequence("industries_id_seq"),
	).RestartWith(1)
	if _, err := conn.Execute(seqStmt); err != nil {
		panic(err)
	}

	seqStmt = postgres.AlterSequence(
		postgres.Sequence("company_industries_id_seq"),
	).RestartWith(1)
	if _, err := conn.Execute(seqStmt); err != nil {
		panic(err)
	}

}
