package runner

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestRunner_GetReportSummary(t *testing.T) {
	type fields struct {
		ctx    context.Context
		loads  []Load
		client LoadTestClient
		report []RequestReport
	}
	tests := []struct {
		name   string
		fields fields
		want   ReportSummary
	}{
		{
			name:   "get report summary",
			fields: fields{
				ctx:    context.Background(),
				report: []RequestReport{
					{
						responseTime: 10 * time.Second,
						success:      true,
						error:        nil,
					},
					{
						responseTime: 20 * time.Second,
						success:      true,
						error:        nil,
					},
				},
			},
			want:   ReportSummary{
				MediumResponseTime: 15 * time.Second,
				NumberOfRequests:   2,
				ErrorPercentage:    0,
			},
		},
		{
			name:   "get report summary with error",
			fields: fields{
				ctx:    context.Background(),
				report: []RequestReport{
					{
						responseTime: 10 * time.Second,
						success:      false,
						error:        errors.New("test"),
					},
					{
						responseTime: 20 * time.Second,
						success:      true,
						error:        nil,
					},
				},
			},
			want:   ReportSummary{
				MediumResponseTime: 15 * time.Second,
				NumberOfRequests:   2,
				ErrorPercentage:    50,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Runner{
				ctx:    tt.fields.ctx,
				loads:  tt.fields.loads,
				client: tt.fields.client,
				report: tt.fields.report,
			}
			if got := r.GetReportSummary(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReportSummary() = %v, want %v", got, tt.want)
			}
		})
	}
}
