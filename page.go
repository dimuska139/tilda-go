package tilda_go

import (
	"context"
	"fmt"
)

// GetPage returns detailed page information
func (c *Client) GetPage(ctx context.Context, pageID string) (Page, error) {
	type Response struct {
		Status string `json:"status"`
		Result Page   `json:"result"`
	}

	var response Response
	if err := c.doRequest(ctx, "/v1/getpage/", map[string]any{
		"pageid": pageID,
	}, &response); err != nil {
		return Page{}, fmt.Errorf("do request: %w", err)
	}

	return response.Result, nil
}

// GetPageFull returns detailed page information
func (c *Client) GetPageFull(ctx context.Context, pageID string) (PageFull, error) {
	type Response struct {
		Status string   `json:"status"`
		Result PageFull `json:"result"`
	}

	var response Response
	if err := c.doRequest(ctx, "/v1/getpagefull/", map[string]any{
		"pageid": pageID,
	}, &response); err != nil {
		return PageFull{}, fmt.Errorf("do request: %w", err)
	}

	return response.Result, nil
}

// GetPageExport returns detailed page information
func (c *Client) GetPageExport(ctx context.Context, pageID string) (PageExport, error) {
	type Response struct {
		Status string     `json:"status"`
		Result PageExport `json:"result"`
	}

	var response Response
	if err := c.doRequest(ctx, "/v1/getpageexport/", map[string]any{
		"pageid": pageID,
	}, &response); err != nil {
		return PageExport{}, fmt.Errorf("do request: %w", err)
	}

	return response.Result, nil
}

// GetPageFullExport returns detailed page information
func (c *Client) GetPageFullExport(ctx context.Context, pageID string) (PageExport, error) {
	type Response struct {
		Status string     `json:"status"`
		Result PageExport `json:"result"`
	}

	var response Response
	if err := c.doRequest(ctx, "/v1/getpagefullexport/", map[string]any{
		"pageid": pageID,
	}, &response); err != nil {
		return PageExport{}, fmt.Errorf("do request: %w", err)
	}

	return response.Result, nil
}
