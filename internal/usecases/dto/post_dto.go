package dto

type Parent struct {
	Type   string `json:"type"`
	PageID string `json:"page_id"`
}

type External struct {
	URL string `json:"url"`
}

type Cover struct {
	Type     string   `json:"type"`
	External External `json:"external"`
}

type PostRequest struct {
	Parent Parent `json:"parent"`
	Cover  Cover  `json:"cover"`
}

type PostResponse struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Cover  Cover  `json:"cover"`
}
