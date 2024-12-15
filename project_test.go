package tilda_go

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestClient_GetProjectsList(t *testing.T) {
	type fields struct {
		config     *Config
		httpClient *http.Client
		baseURL    string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		stubFilename      string
		registerResponder func(responseBody []byte)
		want              []Project
		wantErr           bool
	}{
		{
			name: "success",
			fields: fields{
				config: &Config{
					PublicKey: "public",
					SecretKey: "secret",
				},
				httpClient: http.DefaultClient,
				baseURL:    apiBaseUrl,
			},
			args: args{
				ctx: context.Background(),
			},
			stubFilename: "projects_list.json",
			registerResponder: func(responseBody []byte) {
				url := fmt.Sprintf("%s/v1/getprojectslist/?publickey=public&secretkey=secret", apiBaseUrl)

				httpmock.RegisterResponder(http.MethodGet, url,
					func(req *http.Request) (*http.Response, error) {
						resp := httpmock.NewBytesResponse(http.StatusOK, responseBody)

						return resp, nil
					},
				)
			},
			want: []Project{
				{
					ID:          "123",
					Title:       "My project 1",
					Description: "Description of my project 1",
				}, {
					ID:          "124",
					Title:       "My project 2",
					Description: "Description of my project 2",
				}, {
					ID:          "125",
					Title:       "My project 3",
					Description: "Description of my project 3",
				},
			},
			wantErr: false,
		}, {
			name: "failed",
			fields: fields{
				config: &Config{
					PublicKey: "public",
					SecretKey: "secret",
				},
				httpClient: http.DefaultClient,
				baseURL:    apiBaseUrl,
			},
			args: args{
				ctx: context.Background(),
			},
			stubFilename: "projects_list.json",
			registerResponder: func(_ []byte) {
				url := fmt.Sprintf("%s/v1/getprojectslist/?publickey=public&secretkey=secret", apiBaseUrl)

				httpmock.RegisterResponder(http.MethodGet, url,
					func(req *http.Request) (*http.Response, error) {
						resp := httpmock.NewBytesResponse(http.StatusOK, []byte("{"))

						return resp, nil
					},
				)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			var html []byte
			if tt.stubFilename != "" {
				bts, err := os.ReadFile(fmt.Sprintf("stub/%s", tt.stubFilename))
				assert.NoError(t, err)

				html = bts
			}

			tt.registerResponder(html)

			c := &Client{
				config:     tt.fields.config,
				httpClient: tt.fields.httpClient,
				baseURL:    tt.fields.baseURL,
			}
			got, err := c.GetProjectsList(tt.args.ctx)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClient_GetProjectInfo(t *testing.T) {
	type fields struct {
		config     *Config
		httpClient *http.Client
		baseURL    string
	}
	type args struct {
		ctx       context.Context
		projectID string
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		stubFilename      string
		registerResponder func(responseBody []byte)
		want              Project
		wantErr           bool
	}{
		{
			name: "success",
			fields: fields{
				config: &Config{
					PublicKey: "public",
					SecretKey: "secret",
				},
				httpClient: http.DefaultClient,
				baseURL:    apiBaseUrl,
			},
			args: args{
				ctx:       context.Background(),
				projectID: "123",
			},
			stubFilename: "project.json",
			registerResponder: func(responseBody []byte) {
				url := fmt.Sprintf("%s/v1/getprojectinfo/?projectid=123&publickey=public&secretkey=secret", apiBaseUrl)

				httpmock.RegisterResponder(http.MethodGet, url,
					func(req *http.Request) (*http.Response, error) {
						resp := httpmock.NewBytesResponse(http.StatusOK, responseBody)

						return resp, nil
					},
				)
			},
			want: Project{
				ID:          "12345",
				Title:       "My Site",
				Description: "Description of My Site",
			},
		}, {
			name: "failed",
			fields: fields{
				config: &Config{
					PublicKey: "public",
					SecretKey: "secret",
				},
				httpClient: http.DefaultClient,
				baseURL:    apiBaseUrl,
			},
			args: args{
				ctx:       context.Background(),
				projectID: "123",
			},
			stubFilename: "project.json",
			registerResponder: func(_ []byte) {
				url := fmt.Sprintf("%s/v1/getprojectinfo/?projectid=123&publickey=public&secretkey=secret", apiBaseUrl)

				httpmock.RegisterResponder(http.MethodGet, url,
					func(req *http.Request) (*http.Response, error) {
						resp := httpmock.NewBytesResponse(http.StatusOK, []byte("{"))

						return resp, nil
					},
				)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			var html []byte
			if tt.stubFilename != "" {
				bts, err := os.ReadFile(fmt.Sprintf("stub/%s", tt.stubFilename))
				assert.NoError(t, err)

				html = bts
			}

			tt.registerResponder(html)

			c := &Client{
				config:     tt.fields.config,
				httpClient: tt.fields.httpClient,
				baseURL:    tt.fields.baseURL,
			}
			got, err := c.GetProjectInfo(tt.args.ctx, tt.args.projectID)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClient_GetProjectPages(t *testing.T) {
	type fields struct {
		config     *Config
		httpClient *http.Client
		baseURL    string
	}
	type args struct {
		ctx       context.Context
		projectID string
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		stubFilename      string
		registerResponder func(responseBody []byte)
		want              []Page
		wantErr           bool
	}{
		{
			name: "success",
			fields: fields{
				config: &Config{
					PublicKey: "public",
					SecretKey: "secret",
				},
				httpClient: http.DefaultClient,
				baseURL:    apiBaseUrl,
			},
			args: args{
				ctx:       context.Background(),
				projectID: "123",
			},
			stubFilename: "pages_list.json",
			registerResponder: func(responseBody []byte) {
				url := fmt.Sprintf("%s/v1/getpageslist/?projectid=123&publickey=public&secretkey=secret", apiBaseUrl)

				httpmock.RegisterResponder(http.MethodGet, url,
					func(req *http.Request) (*http.Response, error) {
						resp := httpmock.NewBytesResponse(http.StatusOK, responseBody)

						return resp, nil
					},
				)
			},
			want: []Page{
				{
					ID:          "12345",
					ProjectID:   "54321",
					Date:        DateTime(time.Date(2024, 12, 14, 19, 7, 0, 0, time.UTC)),
					Title:       "Main page",
					Description: "Description of main page",
					Img:         "https://static.tildacdn.com/tild3039-6533-4362-b934-12345/___2024-06-13_13-35-.png",
					Sort:        10,
					Published:   1734258070,
					FeatureImg:  "",
					Alias:       "",
					Filename:    "page12345.html",
				}, {
					ID:          "123456",
					ProjectID:   "54321",
					Date:        DateTime(time.Date(2024, 12, 15, 13, 20, 30, 0, time.UTC)),
					Title:       "Photography blog",
					Description: "Photographer's blog with tiled galleries and email subscription.",
					Img:         "",
					Sort:        20,
					Published:   1734259400,
					FeatureImg:  "",
					Alias:       "blog",
					Filename:    "page123456.html",
				}},
		}, {
			name: "failed",
			fields: fields{
				config: &Config{
					PublicKey: "public",
					SecretKey: "secret",
				},
				httpClient: http.DefaultClient,
				baseURL:    apiBaseUrl,
			},
			args: args{
				ctx:       context.Background(),
				projectID: "123",
			},
			stubFilename: "pages_list.json",
			registerResponder: func(_ []byte) {
				url := fmt.Sprintf("%s/v1/getpageslist/?projectid=123&publickey=public&secretkey=secret", apiBaseUrl)

				httpmock.RegisterResponder(http.MethodGet, url,
					func(req *http.Request) (*http.Response, error) {
						resp := httpmock.NewBytesResponse(http.StatusOK, []byte("{"))

						return resp, nil
					},
				)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			var html []byte
			if tt.stubFilename != "" {
				bts, err := os.ReadFile(fmt.Sprintf("stub/%s", tt.stubFilename))
				assert.NoError(t, err)

				html = bts
			}

			tt.registerResponder(html)

			c := &Client{
				config:     tt.fields.config,
				httpClient: tt.fields.httpClient,
				baseURL:    tt.fields.baseURL,
			}
			got, err := c.GetProjectPages(tt.args.ctx, tt.args.projectID)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
