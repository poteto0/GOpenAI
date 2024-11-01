package gopenai

type messageContent struct {
	messageType string `json:"message_type"`
	content     string `json:"content"`
}

type MessageContent interface{}

func NewMessageContent(messageType, content string) MessageContent {
	return &messageContent{
		messageType: messageType,
		content:     content,
	}
}
