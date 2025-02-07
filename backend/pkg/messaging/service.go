package messaging

import (
	"time"
)

type Service interface {
	StartConversation(participantA, participantB uint) (*Conversation, error)
	SendMessage(conversationID, senderID, receiverID uint, content string) (*Message, error)
	GetConversationMessages(conversationID uint) ([]Message, error)
	MarkMessageAsRead(messageID uint) error
}

type MessagingService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &MessagingService{repository: repository}
}

func (service *MessagingService) StartConversation(participantA, participantB uint) (*Conversation, error) {
	conversation, err := service.repository.GetConversationByParticipants(participantA, participantB)
	if err == nil {
		return conversation, nil
	}
	newConversation := &Conversation{
		ParticipantA: participantA,
		ParticipantB: participantB,
		UpdatedAt:    time.Now(),
	}
	if err := service.repository.CreateConversation(newConversation); err != nil {
		return nil, err
	}
	return newConversation, nil
}

func (service *MessagingService) SendMessage(conversationID, senderID, receiverID uint, content string) (*Message, error) {
	message := &Message{
		ConversationID: conversationID,
		SenderID:       senderID,
		ReceiverID:     receiverID,
		Content:        content,
		SentAt:         time.Now(),
		Read:           false,
	}
	if err := service.repository.CreateMessage(message); err != nil {
		return nil, err
	}
	return message, nil
}

func (service *MessagingService) GetConversationMessages(conversationID uint) ([]Message, error) {
	return service.repository.GetMessagesByConversation(conversationID)
}

func (service *MessagingService) MarkMessageAsRead(messageID uint) error {
	return service.repository.MarkMessageRead(messageID)
}
