package analytics

import (
	"log"
	"sync"
	"time"
)

type Event struct {
	Timestamp  time.Time
	Name       string
	Properties map[string]interface{}
}

type AnalyticsService interface {
	TrackEvent(event string, properties map[string]interface{})
	GetEvents() []Event
}

type SimpleAnalyticsService struct {
	events []Event
	mu     sync.RWMutex
}

func NewSimpleAnalyticsService() *SimpleAnalyticsService {
	return &SimpleAnalyticsService{
		events: make([]Event, 0),
	}
}

func (service *SimpleAnalyticsService) TrackEvent(event string, properties map[string]interface{}) {
	service.mu.Lock()
	defer service.mu.Unlock()
	ev := Event{
		Timestamp:  time.Now(),
		Name:       event,
		Properties: properties,
	}
	service.events = append(service.events, ev)
	log.Printf("Analytics Event: %s - Properties: %v", event, properties)
}

func (service *SimpleAnalyticsService) GetEvents() []Event {
	service.mu.RLock()
	defer service.mu.RUnlock()
	return service.events
}