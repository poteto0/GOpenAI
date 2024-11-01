# GOpenAI

AzureOpenAI wrapper for GoLang.

```sh
go get something
```

You can call OpenAI easily.

```go
func main() {
  // Make your GOpenAI client
  gopen := gopenai.New()

  // Your openai setting
  gs := gopenai.GOpenAISetting{
    Headers: map[string]string{
		  gopenai.HEADER_CONTENT_TYPE: "application/json",
		  gopenai.API_KEY:             "your api key",
	  },
    Endpoint: "<your endpoint>",
    Method: "POST",
  }

  // model Setting
  ps := gopenai.DefaultPayloadSetting
  /*
  ps := gopenai.PayloadSetting{
    Temperature: 0.3,
		TopP:        0.95,
		MaxTokens:   2048,
  }
  */


  // Make Message
  msg := gopenai.NewMessage("user", "text", "Who are you?")

  // Create Request Body
  payload := gopenai.NewPayload(ps)
  payload.SetMessage(msg)
  reqBody, err := payload.ToPayloadBytes()
  if err != nil {
    panic(err)
  }

  // Call Azure OpenAI
  respBody, err := gopenai.Post(reqBody)
  if err != nil {
    panic(err)
  }

  fmt.Println(respBody)
}
```
