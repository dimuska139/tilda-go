package tilda_go

import (
	"context"
	"fmt"
)

// GetProjectsList returns the list of projects
func (c *Client) GetProjectsList(ctx context.Context) ([]Project, error) {
	type Response struct {
		Status string    `json:"status"`
		Result []Project `json:"result"`
	}

	var response Response
	if err := c.doRequest(ctx, "/v1/getprojectslist/", nil, &response); err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	return response.Result, nil
}

// GetProjectInfo returns detailed project information
func (c *Client) GetProjectInfo(ctx context.Context, projectID string) (Project, error) {
	type Response struct {
		Status string  `json:"status"`
		Result Project `json:"result"`
	}

	var response Response
	if err := c.doRequest(ctx, "/v1/getprojectinfo/", map[string]any{
		"projectid": projectID,
	}, &response); err != nil {
		return Project{}, fmt.Errorf("do request: %w", err)
	}

	return response.Result, nil
}

// GetProjectPages returns the list of pages for the project
func (c *Client) GetProjectPages(ctx context.Context, projectID string) ([]Page, error) {
	type Response struct {
		Status string `json:"status"`
		Result []Page `json:"result"`
	}

	var response Response
	if err := c.doRequest(ctx, "/v1/getpageslist/", map[string]any{
		"projectid": projectID,
	}, &response); err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	return response.Result, nil
}
