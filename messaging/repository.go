package messaging

import (
	"gorm.io/gorm"
)

type Repository interface {
	CreateConversation(conversation *Conversation) error
	GetConversationByParticipants(participantA, participantB uint) (*Conversation, error)

	CreateMessage(message *Message) error
	GetMessagesByConversation(conversationID uint) ([]Message, error)
	MarkMessageRead(messageID uint) error
}

type GormRepository struct {
	Database *gorm.DB
}

func NewGormRepository(database *gorm.DB) Repository {
	return &GormRepository{Database: database}
}

func (repository *GormRepository) CreateConversation(conversation *Conversation) error {
	return repository.Database.Create(conversation).Error
}

func (repository *GormRepository) GetConversationByParticipants(participantA, participantB uint) (*Conversation, error) {
	var conversation Conversation
	err := repository.Database.Where("(participant_a = ? AND participant_b = ?) OR (participant_a = ? AND participant_b = ?)", participantA, participantB, participantB, participantA).First(&conversation).Error
	if err != nil {
		return nil, err
	}
	return &conversation, nil
}

func (repository *GormRepository) CreateMessage(message *Message) error {
	return repository.Database.Create(message).Error
}

func (repository *GormRepository) GetMessagesByConversation(conversationID uint) ([]Message, error) {
	var messages []Message
	err := repository.Database.Where("conversation_id = ?", conversationID).Order("sent_at asc").Find(&messages).Error
	return messages, err
}

func (repository *GormRepository) MarkMessageRead(messageID uint) error {
	return repository.Database.Model(&Message{}).Where("id = ?", messageID).Update("read", true).Error
}
