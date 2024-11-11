package gopenai

type message struct {
	role           string         `json:"role"`
	messageContent MessageContent `json:"message_content"`
}

type Message interface{}

func NewMessageByMessageContent(role string, messageContent MessageContent) Message {
	return &message{
		role:           role,
		messageContent: messageContent,
	}
}

func NewMessage(role, messageType, content string) Message {
	return &message{
		role:           role,
		messageContent: NewMessageContent(messageType, content),
	}
}
