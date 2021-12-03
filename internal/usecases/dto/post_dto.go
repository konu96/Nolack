package dto

type Parent struct {
	PageID string `json:"page_id"`
}

type External struct {
	URL string `json:"url"`
}

type Cover struct {
	Type     string   `json:"type"`
	External External `json:"external"`
}

type TitleContent struct {
	Content string `json:"content"`
}

type TitleInTitle struct {
	Type string       `json:"type"`
	Text TitleContent `json:"text"`
}

type Title struct {
	Title []TitleInTitle `json:"title"`
}

type Properties struct {
	Title Title `json:"title"`
}

type PostRequest struct {
	Parent     Parent     `json:"parent"`
	Cover      Cover      `json:"cover"`
	Properties Properties `json:"properties"`
}

type PostResponse struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Cover  Cover  `json:"cover"`
}
