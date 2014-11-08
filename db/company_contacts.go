package db

import (
	"time"

	sql "github.com/aodin/aspect"
	"github.com/aodin/aspect/postgres"
)

type CompanyContact struct {
	ID        int64     `db:"id"`
	CompanyID int64     `db:"company_id"`
	Key       string    `db:"key"`
	Value     string    `db:"value"`
	IsPublic  bool      `db:"is_public"`
	CreatedAt time.Time `db:"created_at"`
}

// CompanyContacts includes all user contact information, such as phone
// numbers, addresses, and social media.
var CompanyContacts = sql.Table("company_contacts",
	sql.Column("id", postgres.Serial{NotNull: true}),
	sql.ForeignKey(
		"company_id",
		Companies.C["id"],
		sql.Integer{NotNull: true},
	),
	sql.Column("key", sql.String{}),
	sql.Column("value", sql.String{}),
	sql.Column("is_public", sql.Boolean{Default: sql.True}),
	sql.Column(
		"created_at",
		sql.Timestamp{Default: "now() at time zone 'utc'"},
	),
	sql.PrimaryKey("id"),
	sql.Unique("company_id", "key"),
)
