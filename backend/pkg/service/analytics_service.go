package service

import "log"

type AnalyticsService interface {
	TrackEvent(event string, properties map[string]interface{})
}

type SimpleAnalyticsService struct {}

func NewSimpleAnalyticsService() *SimpleAnalyticsService {
	return &SimpleAnalyticsService{}
}

func (service *SimpleAnalyticsService) TrackEvent(event string, properties map[string]interface{}) {
	log.Printf("Analytics Event: %s - Properties: %v", event, properties)
}