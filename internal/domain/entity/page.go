package entity

type Page struct {
	Parent     parent     `json:"parent"`
	Cover      cover      `json:"cover"`
	Properties properties `json:"properties"`
}

type parent struct {
	PageID string `json:"page_id"`
}

type external struct {
	URL string `json:"url"`
}

type cover struct {
	Type     string   `json:"type"`
	External external `json:"external"`
}

type titleContent struct {
	Content string `json:"content"`
}

type titleInTitle struct {
	Type string       `json:"type"`
	Text titleContent `json:"text"`
}

type title struct {
	Title []titleInTitle `json:"title"`
}

type properties struct {
	Title title `json:"title"`
}

func NewPage(pageID string, url string) Page {
	return Page{
		Parent: parent{PageID: pageID},
		Cover: cover{
			Type: "external",
			External: external{
				URL: url,
			},
		},
		Properties: properties{Title: title{Title: []titleInTitle{
			{
				Type: "text",
				Text: titleContent{Content: "Sample page created by nolack"},
			},
		}}},
	}
}
