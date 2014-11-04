package db

import (
	sql "github.com/aodin/aspect"
	"github.com/aodin/aspect/postgres"
)

var Industries = sql.Table("industries",
	sql.Column("id", postgres.Serial{NotNull: true}),
	sql.Column("name", sql.String{Unique: true}),
	sql.Column("about", sql.String{}),
	sql.PrimaryKey("id"),
)
