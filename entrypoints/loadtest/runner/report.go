package runner

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type RequestReport struct {
	responseTime time.Duration
	success bool
	error error
}
type ReportSummary struct {
	MediumResponseTime time.Duration
	NumberOfRequests int
	ErrorPercentage float64
}

func (r ReportSummary) String() string {
	return fmt.Sprintf("Total requests: %v  |  Medium response time: %v  |  Error Percentage: %v%%",
		r.NumberOfRequests, r.MediumResponseTime, r.ErrorPercentage)
}

func (r *Runner) GetReportSummary() ReportSummary {
	report := ReportSummary{
		NumberOfRequests:   len(r.report),
	}

	qtdErrs := 0
	var sumResponseTime time.Duration
	for _, requestReport := range r.report {
		if !requestReport.success {
			qtdErrs++
		}

		sumResponseTime += requestReport.responseTime
	}

	report.ErrorPercentage = float64(qtdErrs * 100 / report.NumberOfRequests)
 	report.MediumResponseTime = time.Duration(sumResponseTime.Nanoseconds() / int64(report.NumberOfRequests))

	return report
}

func (r *Runner) ReportToCsv() {
	summary := r.GetReportSummary()
	content := [][]string{{
		"Number of requests",
		"Medium response time",
		"Error percentage",
	}}
	content = append(content, []string{
		fmt.Sprint(summary.NumberOfRequests),
		fmt.Sprint(summary.MediumResponseTime),
		fmt.Sprint(summary.ErrorPercentage),
	})

	content = append(content, []string{"Type", "Requests per second", "Duration"})
	for _, load := range r.loads {
		content = append(content, []string{
			r.code,
			fmt.Sprint(load.CallsPerSecond),
			fmt.Sprint(load.Duration),
		})
	}

	content = append(content, []string{"Response time", "Success", "Error"})

	for _, report := range r.report {
		errMsg := ""
		if report.error != nil {
			errMsg = report.error.Error()
		}

		content = append(content, []string{
			fmt.Sprint(report.responseTime),
			fmt.Sprint(report.success),
			errMsg,
		})
	}

	wr := csv.NewWriter(os.Stdout)
	wr.WriteAll(content)
	wr.Flush()
}
