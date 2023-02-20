package help

import (
	"Getir/internal/ports"
	"testing"
)

func Test_filterRecords(t *testing.T) {
	type args struct {
		records []*ports.Record
		min     int
		max     int
	}
	tests := []struct {
		name string
		args args
		want []*ports.FetchRespRecord
	}{
		{
			name: "success",
			args: args{
				records: []*ports.Record{
					{
						CreatedAt: TestTime("2016-01-30"),
						Counts:    []int{1000, 200},
					},
					{
						CreatedAt: TestTime("2016-01-28"),
						Counts:    []int{100, 200},
					},
				},
				min: 100,
				max: 500,
			},
			want: []*ports.FetchRespRecord{
				{
					CreatedAt:  TestTime("2016-01-28"),
					TotalCount: 300,
				},
			},
		},
		{
			name: "notFound",
			args: args{
				records: []*ports.Record{
					{
						CreatedAt: TestTime("2016-01-28"),
						Counts:    []int{100, 200},
					},
				},
				min: 100,
				max: 200,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecords := FilterRecords(tt.args.records, tt.args.min, tt.args.max)
			AssertRecords(t, gotRecords, tt.want)
		})
	}
}
