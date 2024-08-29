package domain

import (
	"database/sql"
	"mezink/src/business/domain/record"
)

type Domain struct {
	Record record.DomainItf
}

func Init(db *sql.DB) *Domain {
	return &Domain{
		Record: record.InitRecordDomain(db),
	}
}
