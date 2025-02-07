package analytics

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"time"
)

type ReportingService interface {
	GenerateCSVReport() ([]byte, error)
}

type SimpleReportingService struct {
	AnalyticsService AnalyticsService
}

func NewSimpleReportingService(analyticsService AnalyticsService) *SimpleReportingService {
	return &SimpleReportingService{
		AnalyticsService: analyticsService,
	}
}

func (service *SimpleReportingService) GenerateCSVReport() ([]byte, error) {
	events := service.AnalyticsService.GetEvents()
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	header := []string{"Timestamp", "Event Name", "Properties"}
	if err := writer.Write(header); err != nil {
		return nil, err
	}

	for _, event := range events {
		propertiesStr := fmt.Sprintf("%v", event.Properties)
		row := []string{
			event.Timestamp.Format(time.RFC3339),
			event.Name,
			propertiesStr,
		}
		if err := writer.Write(row); err != nil {
			return nil, err
		}
	}
	writer.Flush()
	return buf.Bytes(), nil
}
