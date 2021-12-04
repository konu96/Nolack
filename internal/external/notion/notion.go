package notion

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

var NOTION_PAGE = "https://api.notion.com/v1/pages"
var NOTION_VERSION = "2021-08-16"

type Notion struct {
	Client *http.Client
}

func NewNotion() Notion {
	return Notion{
		Client: http.DefaultClient,
	}
}

func (n Notion) POST(data []byte) (*http.Response, error) {
	request, err := http.NewRequest("POST", NOTION_PAGE, bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("NOTION_TOKEN")))
	request.Header.Set("Notion-Version", NOTION_VERSION)
	request.Header.Set("Content-Type", "application/json")

	resp, err := n.Client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	return resp, err
}
