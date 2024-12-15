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

func TestClient_doRequest(t *testing.T) {
	type fields struct {
		config     *Config
		httpClient *http.Client
		baseURL    string
	}
	type args struct {
		ctx    context.Context
		path   string
		params map[string]any
		result any
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		stubFilename      string
		registerResponder func(responseBody []byte)
		wantErr           bool
	}{
		{
			name: "default",
			fields: fields{
				config: &Config{
					PublicKey: "public",
					SecretKey: "secret",
				},
				httpClient: http.DefaultClient,
				baseURL:    apiBaseUrl,
			},
			args: args{
				ctx:  context.Background(),
				path: "/v1/getprojectslist/",
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
			wantErr: false,
		}, {
			name: "with extra params",
			fields: fields{
				config: &Config{
					PublicKey: "public",
					SecretKey: "secret",
				},
				httpClient: http.DefaultClient,
				baseURL:    apiBaseUrl,
			},
			args: args{
				ctx:  context.Background(),
				path: "/v1/getprojectslist/",
				params: map[string]any{
					"foo":  "bar",
					"foo1": "bar1",
				},
			},
			stubFilename: "projects_list.json",
			registerResponder: func(responseBody []byte) {
				url := fmt.Sprintf("%s/v1/getprojectslist/?foo=bar&foo1=bar1&publickey=public&secretkey=secret", apiBaseUrl)

				httpmock.RegisterResponder(http.MethodGet, url,
					func(req *http.Request) (*http.Response, error) {
						resp := httpmock.NewBytesResponse(http.StatusOK, responseBody)

						return resp, nil
					},
				)
			},
			wantErr: false,
		},
		{
			name: "not 200",
			fields: fields{
				config: &Config{
					PublicKey: "public",
					SecretKey: "secret",
				},
				httpClient: http.DefaultClient,
				baseURL:    apiBaseUrl,
			},
			args: args{
				ctx:  context.Background(),
				path: "/v1/getprojectslist/",
			},
			stubFilename: "projects_list.json",
			registerResponder: func(responseBody []byte) {
				url := fmt.Sprintf("%s/v1/getprojectslist/?publickey=public&secretkey=secret", apiBaseUrl)

				httpmock.RegisterResponder(http.MethodGet, url,
					func(req *http.Request) (*http.Response, error) {
						resp := httpmock.NewBytesResponse(http.StatusInternalServerError, responseBody)

						return resp, nil
					},
				)
			},
			wantErr: true,
		}, {
			name: "bad response body",
			fields: fields{
				config: &Config{
					PublicKey: "public",
					SecretKey: "secret",
				},
				httpClient: http.DefaultClient,
				baseURL:    apiBaseUrl,
			},
			args: args{
				ctx:  context.Background(),
				path: "/v1/getprojectslist/",
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
		}, {
			name: "invalid status response body",
			fields: fields{
				config: &Config{
					PublicKey: "public",
					SecretKey: "secret",
				},
				httpClient: http.DefaultClient,
				baseURL:    apiBaseUrl,
			},
			args: args{
				ctx:  context.Background(),
				path: "/v1/getprojectslist/",
			},
			stubFilename: "projects_list.json",
			registerResponder: func(_ []byte) {
				url := fmt.Sprintf("%s/v1/getprojectslist/?publickey=public&secretkey=secret", apiBaseUrl)

				httpmock.RegisterResponder(http.MethodGet, url,
					func(req *http.Request) (*http.Response, error) {
						resp := httpmock.NewBytesResponse(http.StatusOK, []byte("{\"status\": \"error\"}"))

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
			err := c.doRequest(tt.args.ctx, tt.args.path, tt.args.params, tt.args.result)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		config  *Config
		options []func(*Client)
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{
			name: "default",
			args: args{
				config: &Config{
					PublicKey: "public",
					SecretKey: "secret",
				},
			},
			want: &Client{
				config: &Config{
					PublicKey: "public",
					SecretKey: "secret",
				},
				httpClient: http.DefaultClient,
				baseURL:    apiBaseUrl,
			},
		}, {
			name: "with custom base url",
			args: args{
				config: &Config{
					PublicKey: "public",
					SecretKey: "secret",
				},
				options: []func(*Client){
					WithBaseURL("http://example.com"),
				},
			},
			want: &Client{
				config: &Config{
					PublicKey: "public",
					SecretKey: "secret",
				},
				httpClient: http.DefaultClient,
				baseURL:    "http://example.com",
			},
		}, {
			name: "with custom http client",
			args: args{
				config: &Config{
					PublicKey: "public",
					SecretKey: "secret",
				},
				options: []func(*Client){
					WithCustomHttpClient(&http.Client{
						Timeout: time.Second * 300,
					}),
				},
			},
			want: &Client{
				config: &Config{
					PublicKey: "public",
					SecretKey: "secret",
				},
				httpClient: &http.Client{
					Timeout: time.Second * 300,
				},
				baseURL: apiBaseUrl,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(tt.args.config, tt.args.options...)
			assert.Equal(t, tt.want, client)
		})
	}
}
