package record

import (
	"context"
	"encoding/json"
	"mezink/src/business/entity"
	errors "mezink/stdlib/error"
)

func (r *record) getSQLRecords(ctx context.Context, param entity.RecordParam) ([]entity.Record, error) {
	getRecordQuery := `SELECT id, name, marks, created_at FROM records WHERE 1=1 `
	var args []interface{}
	if param.EndDate != nil && param.StartDate != nil {
		getRecordQuery += " AND created_at BETWEEN ? AND ? "
		args = append(args, *param.StartDate, *param.EndDate)
	}

	rows, err := r.db.QueryContext(ctx, getRecordQuery, args...)
	if err != nil {
		return nil, errors.NewDatabaseError("Failed to fetch records from database", err)
	}
	defer rows.Close()

	var records []entity.Record
	for rows.Next() {
		var record entity.Record
		if err := rows.Scan(&record.ID, &record.Name, &record.Marks, &record.CreatedAt); err != nil {
			return nil, errors.NewDatabaseError("Failed to scan record", err)
		}

		var marks []int64
		err := json.Unmarshal([]byte(record.Marks), &marks)
		if err != nil {
			return nil, err
		}

		for _, m := range marks {
			record.TotalMarks += m
		}

		if param.MinCount != nil && param.MaxCount != nil {
			if record.TotalMarks < *param.MaxCount && record.TotalMarks > *param.MinCount {
				records = append(records, record)
			}
		} else {
			records = append(records, record)
		}

	}

	return records, nil
}
