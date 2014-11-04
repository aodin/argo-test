package db

import (
	sql "github.com/aodin/aspect"
	"github.com/aodin/aspect/postgres"
)

var CompanyIndustries = sql.Table("company_industries",
	sql.Column("id", postgres.Serial{NotNull: true}),
	sql.ForeignKey("industry_id", Industries.C["id"], sql.Integer{}),
	sql.ForeignKey("company_id", Companies.C["id"], sql.Integer{}),
	sql.Unique("industry_id", "company_id"),
	sql.PrimaryKey("id"),
)
