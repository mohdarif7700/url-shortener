package service

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/url-shortener/models"
	"log"
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
	urlMap[req.OriginalURL] = shortURL
	return models.ShortenURLResponse{
		OriginalURL:  req.OriginalURL,
		ShortenedURL: shortURL,
	}, nil
}

func RedirectURL(ctx *fiber.Ctx, shortURL string) error {
	longURL := findLongURL(shortURL)
	if longURL == "" {
		return errors.New("short url not found")
	}

	// Attempt to redirect to the original URL
	if err := ctx.Redirect(longURL, fiber.StatusFound); err != nil {
		// Error occurred during redirection, log the error
		log.Printf("Error redirecting to %s: %v", longURL, err)
		// Return an internal server error response
		return err
	}

	domainMap[longURL]++
	return nil
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

// FindShortURL fetches URL from memory if exists
func findLongURL(shorturl string) string {
	for key, value := range urlMap {
		if value == shorturl {
			return key
		}
	}
	return ""
}
