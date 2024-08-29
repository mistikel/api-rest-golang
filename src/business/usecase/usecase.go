package usecase

import (
	"mezink/src/business/domain"
	"mezink/src/business/usecase/record"
)

type Usecase struct {
	Record record.UsecaseItf
}

func Init(dom *domain.Domain) *Usecase {
	return &Usecase{
		Record: record.InitRecordUsecase(dom.Record),
	}
}
