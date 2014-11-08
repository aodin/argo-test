package db

import (
	sql "github.com/aodin/aspect"
	"github.com/aodin/aspect/postgres"
)

type CompanyIndustry struct {
	ID         int64 `db:"id"`
	IndustryID int64 `db:"industry_id"`
	CompanyID  int64 `db:"company_id"`
}

var CompanyIndustries = sql.Table("company_industries",
	sql.Column("id", postgres.Serial{NotNull: true}),
	sql.ForeignKey(
		"industry_id",
		Industries.C["id"],
		sql.Integer{NotNull: true},
	),
	sql.ForeignKey(
		"company_id",
		Companies.C["id"],
		sql.Integer{NotNull: true},
	),
	sql.Column(
		"created_at",
		sql.Timestamp{Default: "now() at time zone 'utc'"},
	),
	sql.Unique("industry_id", "company_id"),
	sql.PrimaryKey("id"),
)
