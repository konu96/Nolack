package usecases

import (
	"encoding/json"
	"github.com/slack-go/slack/slackevents"
	"net/http"
)

type VerifyURLInteractor struct {
}

func NewVerifyURLInteractor() VerifyURLInteractor {
	return VerifyURLInteractor{}
}

func (i *VerifyURLInteractor) Exec(w http.ResponseWriter, body []byte) error {
	var res *slackevents.ChallengeResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return err
	}

	w.Header().Set("Content-Type", "text/plain")
	if _, err := w.Write([]byte(res.Challenge)); err != nil {
		return err
	}

	return nil
}
