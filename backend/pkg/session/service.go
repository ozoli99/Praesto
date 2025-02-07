package session

import "time"

type Service interface {
	LogSession(providerID, clientID uint, serviceType string, duration int, notes string, sessionDate time.Time) (*Session, error)
	UpdateSession(session *Session) error

	CreateTreatmentPlan(providerID, clientID, sessionID uint, planDetails string, followUpDate time.Time) (*TreatmentPlan, error)
	UpdateTreatmentPlan(plan *TreatmentPlan) error
}

type SessionService struct {
	Repository Repository
}

func NewSessionService(repository Repository) Service {
	return &SessionService{Repository: repository}
}

func (service *SessionService) LogSession(providerID, clientID uint, serviceType string, duration int, notes string, sessionDate time.Time) (*Session, error) {
	session := &Session{
		ProviderID:  providerID,
		ClientID:    clientID,
		ServiceType: serviceType,
		Duration:    duration,
		Notes:       notes,
		SessionDate: sessionDate,
	}
	if err := service.Repository.CreateSession(session); err != nil {
		return nil, err
	}
	return session, nil
}

func (service *SessionService) UpdateSession(session *Session) error {
	return service.Repository.UpdateSession(session)
}

func (service *SessionService) CreateTreatmentPlan(providerID, clientID, sessionID uint, planDetails string, followUpDate time.Time) (*TreatmentPlan, error) {
	plan := &TreatmentPlan{
		ProviderID:   providerID,
		ClientID:     clientID,
		SessionID:    sessionID,
		PlanDetails:  planDetails,
		FollowUpDate: followUpDate,
		Completed:    false,
	}
	if err := service.Repository.CreateTreatmentPlan(plan); err != nil {
		return nil, err
	}
	return plan, nil
}

func (service *SessionService) UpdateTreatmentPlan(plan *TreatmentPlan) error {
	return service.Repository.UpdateTreatmentPlan(plan)
}
