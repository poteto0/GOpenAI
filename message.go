package gopenai

type message struct {
	role           string          `json:"role"`
	messageContent *messageContent `json:"message_content"`
}

type Message interface {
	GetMessageContent() string
}

func NewMessage(role, messageType, content string) Message {
	return &message{
		role:           role,
		messageContent: NewMessageContent(messageType, content),
	}
}

func (m *message) GetMessageContent() string {
	return m.messageContent.GetContent()
}
