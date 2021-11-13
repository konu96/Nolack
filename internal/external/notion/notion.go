package notion

import (
	"encoding/json"
	"github.com/konu96/Nolack/internal/domain/entity"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var NOTION_PAGE = "https://api.notion.com/v1/pages/"
var NOTION_VERSION = "2021-08-16"

type Notion struct {
	client *http.Client
}

type Cover struct {
	Type     string `json:"type"`
	External struct {
		Url string `json:"url"`
	} `json:"external"`
}

type PostResponse struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Cover  Cover  `json:"cover"`
}

func NewNotion() Notion {
	return Notion{
		client: http.DefaultClient,
	}
}

func (n *Notion) POST() entity.Page {
	request, err := http.NewRequest("POST", NOTION_PAGE, nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Authorization", "Bearer"+os.Getenv("NOTION_TOKEN"))
	request.Header.Add("Notion-Version", NOTION_VERSION)

	resp, err := n.client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return entity.Page{}
	}

	var page entity.Page
	if err := json.Unmarshal(body, &page); err != nil {
		return entity.Page{}
	}

	return page
}
