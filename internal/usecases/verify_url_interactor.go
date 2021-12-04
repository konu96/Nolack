package usecases

import (
	"encoding/json"
	"fmt"
	"github.com/slack-go/slack/slackevents"
	"net/http"
)

type VerifyURLUseCase interface {
	Exec(w http.ResponseWriter, body []byte) error
}

type VerifyURLInteractor struct {
}

func NewVerifyURLInteractor() *VerifyURLInteractor {
	return &VerifyURLInteractor{}
}

func (i *VerifyURLInteractor) Exec(w http.ResponseWriter, body []byte) error {
	var res *slackevents.ChallengeResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return fmt.Errorf("unmarshal verify url response: %w", err)
	}

	w.Header().Set("Content-Type", "text/plain")
	if _, err := w.Write([]byte(res.Challenge)); err != nil {
		return fmt.Errorf("write verify url response: %w", err)
	}

	return nil
}
