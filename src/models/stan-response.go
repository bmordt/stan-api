package models

type StanResponse struct {
	Response []Response `json:"response"`
}

type Response struct {
	Image string `json:"image"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
}
