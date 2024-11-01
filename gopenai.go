package gopenai

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type GOpenAISetting struct {
	Headers  map[string]string `json:"header"`
	Endpoint string            `json:"endpoint"`
	Method   string            `json:"method"`
}

var DefaultGOpenAISetting = GOpenAISetting{
	Headers: map[string]string{
		HEADER_CONTENT_TYPE: DEFAULT_CONTENT_TYPE,
		API_KEY:             "",
	},
	Endpoint: "",
	Method:   "POST",
}

type gOpenAI struct {
	setting GOpenAISetting `json:"setting"`
}

type GOpenAI interface {
	LoadSetting(setting GOpenAISetting)
	CreateRequest(body []byte) (*http.Request, error)
	Post(body []byte) (string, error)
}

func New() GOpenAI {
	return &gOpenAI{}
}

func (gopen *gOpenAI) LoadSetting(setting GOpenAISetting) {
	if setting.Headers == nil {
		setting.Headers = DefaultGOpenAISetting.Headers
	}

	if setting.Endpoint == "" {
		setting.Endpoint = DefaultGOpenAISetting.Endpoint
	}

	if setting.Method == "" {
		setting.Method = DefaultGOpenAISetting.Method
	}

	gopen.setting = setting
}

func (gopen *gOpenAI) CreateRequest(body []byte) (*http.Request, error) {
	var (
		endpoint = gopen.setting.Endpoint
		method   = gopen.setting.Method
		headers  = gopen.setting.Headers
	)

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Fatal create request: ", err)
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return req, nil
}

func (gopen *gOpenAI) Post(body []byte) (string, error) {
	req, err := gopen.CreateRequest(body)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("Fatal request with status code: %v %v", resp.StatusCode, body)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Fatal read response: ", err)
		return "", err
	}
	return string(respBody), nil
}
