package record_test

import (
	"context"
	"errors"
	"fmt"
	"mezink/src/business/entity"
	"mezink/src/business/usecase/record"
	mock_record "mezink/src/mock/business/domin/record"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/mock/gomock"
)

var (
	usecase          record.UsecaseItf
	mockRecordDomain *mock_record.MockDomainItf
)

func provideTest(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRecordDomain = mock_record.NewMockDomainItf(ctrl)

	usecase = record.InitRecordUsecase(mockRecordDomain)

	return func() {}
}

func TestGetRecord(t *testing.T) {
	testDep := provideTest(t)
	defer testDep()

	type args struct {
		ctx   context.Context
		param entity.RecordParam
	}

	ctx := context.Background()
	tests := []struct {
		testID   int
		testType string
		testDesc string
		args     args
		mockFunc func()
	}{
		{
			testID:   1,
			testType: "P",
			testDesc: "Success Get Records with empty param",
			args: args{
				ctx: ctx,
			},
			mockFunc: func() {
				mockRecordDomain.EXPECT().GetRecords(ctx, entity.RecordParam{}).Return([]entity.Record{
					{
						ID:         1,
						Marks:      "[1,2,3,4]",
						TotalMarks: 10,
						CreatedAt:  time.Now(),
					},
				}, nil)
			},
		},
		{
			testID:   2,
			testType: "N",
			testDesc: "Failed Get Records",
			args: args{
				ctx: ctx,
			},
			mockFunc: func() {
				mockRecordDomain.EXPECT().GetRecords(ctx, entity.RecordParam{}).Return([]entity.Record{}, errors.New("errMock"))
			},
		},
	}

	Convey("TestGetRecord", t, func() {
		for _, tt := range tests {
			tt.mockFunc()
			_, err := usecase.GetRecords(tt.args.ctx, tt.args.param)
			Convey(fmt.Sprintf("%d - [%s] : %s", tt.testID, tt.testType, tt.testDesc), func() {
				if tt.testType == "P" {
					So(err, ShouldBeNil)
					return
				}
				So(err, ShouldNotBeNil)
			})
		}
	})
}
