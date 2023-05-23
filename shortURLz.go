package shortURLz

import "fmt"

type ShortenUrl struct {
	Shorten   string `json:"link"`
	Original  string `json:"long_url"`
	IsDeleted bool   `json:"is_deleted"`
	Group     string
}

func (curl *ShortenUrl) String() string {
	return fmt.Sprintf("%s (%s): (%s)", curl.Shorten, curl.Group, curl.Original)
}

type URLShortener interface {
	List(config *Config) ([]*ShortenUrl, error)
	Shorten(config *Config, url string) (*ShortenUrl, error)
	Delete(config *Config, shortenURL string) error
}
