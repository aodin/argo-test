package main

import (
	dbsql "database/sql"

	sql "github.com/aodin/aspect"
	"github.com/aodin/aspect/postgres"

	db "github.com/aodin/argo-test/db"
)

func Up_20141108132205(txn *dbsql.Tx) {
	conn := sql.WrapTx(txn, sql.MustGetDialect("postgres"))
	if _, err := conn.Execute(db.Users.Create()); err != nil {
		panic(err)
	}

	users := []db.User{
		{ID: 1, Email: "client@example.com", Password: "secret"},
	}
	stmt := sql.Insert(
		db.Users.C["id"],
		db.Users.C["email"],
		db.Users.C["password"],
	).Values(users)
	if _, err := conn.Execute(stmt); err != nil {
		panic(err)
	}

	// Set the auto-increments
	seqStmt := postgres.AlterSequence(
		postgres.Sequence("users_id_seq"),
	).RestartWith(len(users) + 1)
	if _, err := conn.Execute(seqStmt); err != nil {
		panic(err)
	}
}

func Down_20141108132205(txn *dbsql.Tx) {
	conn := sql.WrapTx(txn, sql.MustGetDialect("postgres"))
	if _, err := conn.Execute(db.Users.Drop().IfExists()); err != nil {
		panic(err)
	}
}
