package notion

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

var Page = "https://api.notion.com/v1/pages"
var Version = "2021-08-16"

type Notion struct {
	Client *http.Client
}

func NewNotion() Notion {
	return Notion{
		Client: http.DefaultClient,
	}
}

func (n Notion) POST(data []byte) (*http.Response, error) {
	request, err := http.NewRequest("POST", Page, bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("create new request: %w", err)
	}

	setHeaders(request)

	response, err := n.Client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("post: %w", err)
	}
	return response, err
}

func setHeaders(request *http.Request) {
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("NOTION_TOKEN")))
	request.Header.Set("Notion-Version", Version)
	request.Header.Set("Content-Type", "application/json")
}
