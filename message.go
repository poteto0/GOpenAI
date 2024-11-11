package gopenai

type Message struct {
	role           string          `json:"role"`
	messageContent *messageContent `json:"message_content"`
}

func NewMessage(role, messageType, content string) Message {
	return Message{
		role:           role,
		messageContent: NewMessageContent(messageType, content),
	}
}

func (m *Message) GetMessageContent() string {
	return m.messageContent.GetContent()
}
