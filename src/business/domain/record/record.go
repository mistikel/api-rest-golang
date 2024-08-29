package record

import (
	"context"
	"database/sql"
	"mezink/src/business/entity"
)

type DomainItf interface {
	GetRecords(ctx context.Context, param entity.RecordParam) ([]entity.Record, error)
}
type record struct {
	db *sql.DB
}

func InitRecordDomain(db *sql.DB) DomainItf {
	return &record{
		db: db,
	}
}
