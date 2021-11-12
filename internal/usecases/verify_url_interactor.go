package usecases

import (
	"encoding/json"
	"github.com/slack-go/slack/slackevents"
	"log"
	"net/http"
)

type VerifyURLInteractor struct {
}

func NewVerifyURLInteractor() VerifyURLInteractor {
	return VerifyURLInteractor{}
}

func (i *VerifyURLInteractor) Exec(w http.ResponseWriter, body []byte) {
	var res *slackevents.ChallengeResponse
	if err := json.Unmarshal(body, &res); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	if _, err := w.Write([]byte(res.Challenge)); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
