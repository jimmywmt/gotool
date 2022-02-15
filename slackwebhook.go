package gotool

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
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
func (slackwebhook *Slackwebhook) SentMessage(message string) {
	slackBody, _ := json.Marshal(SlackRequestBody{Text: message})
	req, err := http.NewRequest("POST", slackwebhook.url, bytes.NewBuffer(slackBody))
	if err != nil {
		log.WithError(err).Warnln("connection failed")
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.WithError(err).Warnln("connection failed")
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		log.WithError(err).Warnln("non-ok response returned from Slack")
	}
	defer resp.Body.Close()
}
