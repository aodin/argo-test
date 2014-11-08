package db

import (
	sql "github.com/aodin/aspect"
	"github.com/aodin/aspect/postgres"
)

type Industry struct {
	ID    int64  `db:"id"`
	Name  string `db:"name"`
	About string `db:"about"`
}

var Industries = sql.Table("industries",
	sql.Column("id", postgres.Serial{NotNull: true}),
	sql.Column("name", sql.String{Unique: true}),
	sql.Column("about", sql.String{}),
	sql.Column(
		"created_at",
		sql.Timestamp{Default: "now() at time zone 'utc'"},
	),
	sql.PrimaryKey("id"),
)
