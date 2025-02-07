package messaging

import (
	"time"

	"github.com/ozoli99/Praesto/models"
)

type Conversation struct {
	models.Base
	ParticipantA uint      `json:"participant_a"`
	ParticipantB uint      `json:"participant_b"`
	LastMessage  string    `json:"last_message"`
	UpdatedAt    time.Time `json:"updated_at"`
}
