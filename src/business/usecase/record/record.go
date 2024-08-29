package record

import (
	"context"
	"mezink/src/business/domain/record"
	"mezink/src/business/entity"
)

type UsecaseItf interface {
	GetRecords(ctx context.Context, param entity.RecordParam) ([]entity.Record, error)
}

type usecase struct {
	dom record.DomainItf
}

func InitRecordUsecase(dom record.DomainItf) UsecaseItf {
	return &usecase{
		dom: dom,
	}
}
