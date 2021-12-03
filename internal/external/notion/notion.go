package notion

import (
	"bytes"
	"encoding/json"
	"github.com/konu96/Nolack/internal/usecases/dto"
	"io/ioutil"
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

func (n Notion) POST(data []byte) (*dto.PostResponse, error) {
	request, err := http.NewRequest("POST", NOTION_PAGE, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Authorization", "Bearer "+os.Getenv("NOTION_TOKEN"))
	request.Header.Add("Notion-Version", NOTION_VERSION)
	request.Header.Set("Content-Type", "application/json")

	resp, err := n.Client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var page dto.PostResponse
	if err := json.Unmarshal(body, &page); err != nil {
		return nil, err
	}

	return &page, nil
}
