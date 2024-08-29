package entity

import (
	"time"
)

type Record struct {
	ID         int64     `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Marks      string    `json:"-" db:"marks"`
	TotalMarks int64     `json:"totalMarks"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
}

type RecordParam struct {
	StartDate *time.Time `json:"startDate" validate:"required_with=EndDate"`
	EndDate   *time.Time `json:"endDate" validate:"required_with=StartDate"`
	MinCount  *int64     `json:"minCount" validate:"required_with=MaxCount"`
	MaxCount  *int64     `json:"maxCount" validate:"required_with=MinCount,gtefield=MinCount"`
}
