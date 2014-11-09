package db

import (
	"time"

	sql "github.com/aodin/aspect"
	"github.com/aodin/aspect/postgres"
)

type User struct {
	ID        int64     `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	IsActive  bool      `db:"is_public"`
	CreatedAt time.Time `db:"created_at"`
}

var Users = sql.Table("users",
	sql.Column("id", postgres.Serial{NotNull: true}),
	sql.Column("email", sql.String{NotNull: true}),
	sql.Column("password", sql.String{NotNull: true}),
	sql.Column("is_active", sql.Boolean{Default: sql.True}),
	sql.Column(
		"created_at",
		sql.Timestamp{Default: "now() at time zone 'utc'"},
	),
	sql.PrimaryKey("id"),
	sql.Unique("email"),
)
