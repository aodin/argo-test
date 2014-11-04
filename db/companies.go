package db

import (
	sql "github.com/aodin/aspect"
	"github.com/aodin/aspect/postgres"
)

var Companies = sql.Table("companies",
	sql.Column("id", postgres.Serial{NotNull: true}),
	sql.Column("name", sql.String{Length: 255, Unique: true}),
	sql.Column("is_active", sql.Boolean{NotNull: true, Default: sql.True}),
	sql.Column("created_at", sql.Timestamp{Default: "now() at time zone 'utc'"}),
	sql.PrimaryKey("id"),
)
