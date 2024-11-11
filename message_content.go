package gopenai

type messageContent struct {
	messageType string `json:"message_type"`
	content     string `json:"content"`
}

func NewMessageContent(messageType, content string) *messageContent {
	return &messageContent{
		messageType: messageType,
		content:     content,
	}
}

func (mc *messageContent) GetContent() string {
	return mc.content
}
