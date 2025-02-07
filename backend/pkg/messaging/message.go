package messaging

import (
	"time"

	"github.com/ozoli99/Praesto/pkg/models"
)

type Message struct {
	models.Base
	ConversationID uint      `json:"conversation_id"`
	SenderID       uint      `json:"sender_id"`
	ReceiverID     uint      `json:"receiver_id"`
	Content        string    `json:"content"`
	SentAt         time.Time `json:"sent_at"`
	Read           bool      `json:"read"`
}
