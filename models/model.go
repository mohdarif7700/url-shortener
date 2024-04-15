package models

type ShortenURLRequest struct {
	OriginalURL string `json:"originalURL"`
}

type ShortenURLResponse struct {
	OriginalURL  string `json:"originalURL"`
	ShortenedURL string `json:"shortenedURL"`
}
