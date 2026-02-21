package geocoder

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const apiBaseURL = "https://api.opencagedata.com/geocode/v1/json"

type Client struct {
	apiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}

// Geocode виконує запит і повертає перший результат або помилку
func (c *Client) Geocode(query string) (*Result, error) {
	// Формування URL з екрануванням спецсимволів
	requestURL := fmt.Sprintf("%s?q=%s&key=%s",
		apiBaseURL, url.QueryEscape(query), c.apiKey)

	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned non-200 status: %d", resp.StatusCode)
	}

	var apiResponse Response
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	// Якщо результатів немає, повертаємо nil
	if len(apiResponse.Results) == 0 {
		return nil, nil
	}

	return apiResponse.Results[0], nil
}
