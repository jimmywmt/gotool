package gotool

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type Slackwebhook struct {
	url string
}

type SlackRequestBody struct {
	Text string `json:"text"`
}

func NewSlackWebhook(url string) *Slackwebhook {
	return &Slackwebhook{url: url}
}
func (slackwebhook *Slackwebhook) SentMessage(message string) error {
	slackBody, _ := json.Marshal(SlackRequestBody{Text: message})
	req, err := http.NewRequest("POST", slackwebhook.url, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("non-ok response returned from Slack")
	}
	return nil
}
