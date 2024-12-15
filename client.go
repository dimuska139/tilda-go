package tilda_go

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const apiBaseUrl string = "https://api.tildacdn.info"

type Config struct {
	PublicKey string
	SecretKey string
}

// TildaError represents information about errors
type TildaError struct {
	HttpCode int    // HTTP status code (https://en.wikipedia.org/wiki/List_of_HTTP_status_codes)
	Url      string // URL of Tilda endpoint associated with callable function
	Body     string // Raw body of response
	Err      error  // Error
}

// Error() converts error to string
func (e *TildaError) Error() string {
	return fmt.Sprintf("Http code: %d, url: %s, body: %s, message: %s", e.HttpCode, e.Url, e.Body, e.Err.Error())
}

type Client struct {
	config     *Config
	httpClient *http.Client
	baseURL    string
}

// NewClient creates new Tilda client
func NewClient(config *Config, options ...func(*Client)) *Client {
	client := &Client{
		config:     config,
		httpClient: http.DefaultClient,
		baseURL:    apiBaseUrl,
	}

	for _, o := range options {
		o(client)
	}

	return client
}

// WithBaseURL option allows to set custom base url (if you use proxy, for example)
func WithBaseURL(baseURL string) func(*Client) {
	return func(s *Client) {
		s.baseURL = baseURL
	}
}

// WithCustomHttpClient option allows to set custom http client
func WithCustomHttpClient(httpClient *http.Client) func(*Client) {
	return func(s *Client) {
		s.httpClient = httpClient
	}
}

func (c *Client) doRequest(ctx context.Context, path string, params map[string]any, result any) error {
	url := c.baseURL + path

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("content-type", "application/json;charset=utf-8")
	q := req.URL.Query()

	q.Add("publickey", c.config.PublicKey)
	q.Add("secretkey", c.config.SecretKey)

	for param, value := range params {
		q.Add(param, fmt.Sprintf("%v", value))
	}
	req.URL.RawQuery = q.Encode()
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return &TildaError{http.StatusServiceUnavailable, url, "", fmt.Errorf("do request: %w", err)}
	}

	if resp != nil {
		defer resp.Body.Close()
	} else {
		return &TildaError{http.StatusServiceUnavailable, url, "", errors.New("response is nil")}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &TildaError{resp.StatusCode, url, string(body), fmt.Errorf("read response body: %w", err)}
	}

	if resp.StatusCode != http.StatusOK {
		return &TildaError{resp.StatusCode, url, string(body), errors.New("response status code is not 200")}
	}

	type ResponseCheck struct {
		Status string `json:"status"`
	}

	var responseCheck ResponseCheck
	if err := json.Unmarshal(body, &responseCheck); err != nil {
		return &TildaError{resp.StatusCode, url, string(body), fmt.Errorf("unmarshal response: %w", err)}
	}

	if responseCheck.Status != "FOUND" {
		return &TildaError{resp.StatusCode, url, string(body), errors.New("invalid status in response, expected FOUND")}
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return &TildaError{resp.StatusCode, url, string(body), fmt.Errorf("unmarshal response: %w", err)}
	}

	return nil
}
