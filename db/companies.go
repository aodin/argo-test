package db

import (
	"time"

	sql "github.com/aodin/aspect"
	"github.com/aodin/aspect/postgres"
)

type Company struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	IsActive  bool      `db:"is_active"`
	CreatedAt time.Time `db:"created_at"`
}

var Companies = sql.Table("companies",
	sql.Column("id", postgres.Serial{NotNull: true}),
	sql.Column("name", sql.String{Length: 255}),
	sql.Column("is_active", sql.Boolean{NotNull: true, Default: sql.True}),
	sql.Column(
		"created_at",
		sql.Timestamp{Default: "now() at time zone 'utc'"},
	),
	sql.PrimaryKey("id"),
	sql.Unique("name"),
)
