package gopenai

import (
	"encoding/json"
	"fmt"
)

type PayloadSetting struct {
	Temperature float64 `json:"Temperature"`
	TopP        float64 `json:"top_p"`
	MaxTokens   int     `json:"max_tokens"`
}

var DefaultPayloadSetting = PayloadSetting{
	Temperature: 0.3,
	TopP:        0.95,
	MaxTokens:   2048,
}

type payload struct {
	Setting  PayloadSetting `json:"setting"`
	Messages []*message     `json:"messages"`
}

type Payload interface {
	GetMessages() []*message
	AppendMessage(messages ...Message)
	ToPayloadBytes() ([]byte, error)
}

func NewPayload(setting PayloadSetting) Payload {
	return &payload{
		Setting: setting,
	}
}

func (p *payload) GetMessages() []*message {
	return p.Messages
}

func (p *payload) AppendMessage(messages ...Message) {
	if len(messages) == 0 {
		return
	}
	for _, msg := range messages {
		p.Messages = append(p.Messages, msg.(*message))
	}
}

func (p *payload) ToPayloadBytes() ([]byte, error) {
	payloadBytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Fatal: ToPayloadBytes: ", err)
		return nil, err
	}
	return payloadBytes, nil
}
