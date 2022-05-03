package runner

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
	"log"
)

type RequestReport struct {
	responseTime time.Duration
	endTime time.Time
	success bool
	error error
}
type ReportSummary struct {
	MediumResponseTime time.Duration
	NumberOfRequests   int
	ErrorPercentage    float64
	Throughput         float64
	StartTime time.Time
	EndTime time.Time
}

func (r ReportSummary) String() string {
	return fmt.Sprintf("Total requests: %v  |  Medium response time: %v  |  Error Percentage: %v%%  |  Throughput: %v",
		r.NumberOfRequests, r.MediumResponseTime, r.ErrorPercentage, r.Throughput)
}

func (r *Runner) GetReportSummary() ReportSummary {
	report := ReportSummary{
		NumberOfRequests:   len(r.report),
		StartTime: r.startTime,
		EndTime: r.endTime,
	}

	qtdErrs := 0
	var sumResponseTime time.Duration
	for _, requestReport := range r.report {
		if !requestReport.success {
			qtdErrs++
		}

		sumResponseTime += requestReport.responseTime
	}

	if report.NumberOfRequests == 0 {
		return report
	}

	report.ErrorPercentage = float64(qtdErrs * 100 / report.NumberOfRequests)
 	report.MediumResponseTime = time.Duration(sumResponseTime.Nanoseconds() / int64(report.NumberOfRequests))
 	report.Throughput = float64(report.NumberOfRequests) / (report.EndTime.Sub(report.StartTime).Seconds())

	log.Println(report.EndTime.Sub(report.StartTime).Seconds())
	return report
}

func (r *Runner) ReportToCsv() {
	summary := r.GetReportSummary()
	content := [][]string{{
		"Number of requests",
		"Medium response time",
		"Error percentage",
		"Throughput",
	}}
	content = append(content, []string{
		fmt.Sprint(summary.NumberOfRequests),
		fmt.Sprint(summary.MediumResponseTime),
		fmt.Sprint(summary.ErrorPercentage),
		fmt.Sprint(summary.Throughput),
	})

	content = append(content, []string{"Type", "Requests per second", "Duration", ""})
	for _, load := range r.loads {
		content = append(content, []string{
			r.code,
			fmt.Sprint(load.CallsPerSecond),
			fmt.Sprint(load.Duration),
			"",
		})
	}

	content = append(content, []string{"Response time", "Success", "Error", "End time"})

	for _, report := range r.report {
		errMsg := "-"
		if report.error != nil {
			errMsg = report.error.Error()
		}

		content = append(content, []string{
			fmt.Sprint(report.responseTime),
			fmt.Sprint(report.success),
			errMsg,
			fmt.Sprint(report.endTime.Format(time.StampNano)),
		})
	}

	wr := csv.NewWriter(os.Stdout)
	wr.WriteAll(content)
	wr.Flush()
}
