package record

import (
	"context"
	"mezink/src/business/entity"
	"mezink/stdlib/log"
)

// GetRecords
// Get list of records from table with filter
func (u *usecase) GetRecords(ctx context.Context, param entity.RecordParam) ([]entity.Record, error) {
	result, err := u.dom.GetRecords(ctx, param)
	if err != nil {
		log.ErrContext(ctx, "[GetRecords] error get records: ", err.Error())
		return nil, err
	}

	return result, nil
}
