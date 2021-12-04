package dto

type CreatePageResponse struct {
	ID     string `json:"id"`
	Object string `json:"object"`
}

type CreatePageErrorResponse struct {
	Object  string `json:"object"`
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}
