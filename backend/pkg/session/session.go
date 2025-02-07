package session

import (
	"time"

	"github.com/ozoli99/Praesto/pkg/models"
)

type Session struct {
	models.Base
	ProviderID  uint      `json:"provider_id"`
	ClientID    uint      `json:"client_id"`
	ServiceType string    `json:"service_type"`
	Duration    int       `json:"duration"`
	Notes       string    `json:"notes"`
	SessionDate time.Time `json:"session_date"`
}

type TreatmentPlan struct {
	models.Base
	ProviderID   uint      `json:"provider_id"`
	ClientID     uint      `json:"client_id"`
	SessionID    uint      `json:"session_id"`
	PlanDetails  string    `json:"plan_details"`
	FollowUpDate time.Time `json:"follow_up_date"`
	Completed    bool      `json:"completed"`
}
