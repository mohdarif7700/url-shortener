package service

import (
	"errors"
	"github.com/url-shortener/models"
	"sort"
	"sync"
)

var (
	mutex     sync.Mutex
	urlMap    = make(map[string]string)
	domainMap = make(map[string]int)
)

func CreateShortenURL(req models.ShortenURLRequest) (models.ShortenURLResponse, error) {
	mutex.Lock()
	defer mutex.Unlock()

	// Check if the long URL has already been shortened
	if shortURL, ok := urlMap[req.OriginalURL]; ok {
		return models.ShortenURLResponse{
			OriginalURL:  req.OriginalURL,
			ShortenedURL: shortURL,
		}, nil
	}

	shortURL := CreateShortLink(req.OriginalURL)

	// Store the mapping of long URL to short URL
	urlMap[shortURL] = req.OriginalURL
	return models.ShortenURLResponse{
		OriginalURL:  req.OriginalURL,
		ShortenedURL: shortURL,
	}, nil
}

func RedirectURL(req models.RedirectURLRequest) (string, error) {
	mutex.Lock()
	defer mutex.Unlock()

	// Check if the short URL exists in the map
	longURL, ok := urlMap[req.ShortURL]
	if !ok {
		return "", errors.New("not found")
	}

	// Increment domain count for metrics
	domainMap[longURL]++
	return longURL, nil
}

func GetMetrics() (map[string]int, error) {
	mutex.Lock()
	defer mutex.Unlock()

	// Sort domainMap by count in descending order
	sortedDomains := make([]string, 0, len(domainMap))
	for domain := range domainMap {
		sortedDomains = append(sortedDomains, domain)
	}

	sort.Slice(sortedDomains, func(i, j int) bool {
		return domainMap[sortedDomains[i]] > domainMap[sortedDomains[j]]
	})

	// Get top 3 domains
	topDomains := make(map[string]int)
	for i, domain := range sortedDomains {
		if i >= 3 {
			break
		}
		topDomains[domain] = domainMap[domain]
	}

	return topDomains, nil
}
