package record

import (
	"context"
	"mezink/src/business/entity"
)

// GetRecords
// Get list of records based on filter startDate and endDate
// Avoid filter min and max count in query due to query performance
func (r *record) GetRecords(ctx context.Context, param entity.RecordParam) ([]entity.Record, error) {
	return r.getSQLRecords(ctx, param)
}
