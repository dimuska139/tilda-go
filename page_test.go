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

func TestClient_GetPage(t *testing.T) {
	type fields struct {
		config     *Config
		httpClient *http.Client
		baseURL    string
	}
	type args struct {
		ctx    context.Context
		pageID string
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		stubFilename      string
		registerResponder func(responseBody []byte)
		want              Page
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
				ctx:    context.Background(),
				pageID: "123",
			},
			stubFilename: "page.json",
			registerResponder: func(responseBody []byte) {
				url := fmt.Sprintf("%s/v1/getpage/?pageid=123&publickey=public&secretkey=secret", apiBaseUrl)

				httpmock.RegisterResponder(http.MethodGet, url,
					func(req *http.Request) (*http.Response, error) {
						resp := httpmock.NewBytesResponse(http.StatusOK, responseBody)

						return resp, nil
					},
				)
			},
			want: Page{
				ID:          "12345",
				ProjectID:   "54321",
				Title:       "Photography blog",
				Description: "Photographer's blog with tiled galleries and email subscription.",
				Date:        DateTime(time.Date(2024, 12, 15, 13, 20, 30, 0, time.UTC)),
				Sort:        20,
				Published:   1734259400,
				Alias:       "blog",
				Filename:    "page12345.html",
				HTML:        "<!--allrecords--> <div id=\"allrecords\" class=\"t-records\" data-hook=\"blocks-collection-content-node\" data-tilda-project-id=\"54321\" data-tilda-page-id=\"12345\" data-tilda-page-alias=\"blog\" data-tilda-formskey=\"qwerty\" data-tilda-cookie=\"no\" data-tilda-lazy=\"yes\" data-tilda-root-zone=\"one\"></div> <!--/allrecords-->",
				JS: []string{
					"https://static.tildacdn.com/js/tilda-polyfill-1.0.min.js",
					"https://static.tildacdn.com/js/tilda-scripts-3.0.min.js",
					"https://static.tildacdn.com/ws/project54321/tilda-blocks-page12345.min.js?t=1734259633",
					"https://static.tildacdn.com/js/tilda-lazyload-1.0.min.js",
					"https://static.tildacdn.com/js/tilda-forms-1.0.min.js",
					"https://static.tildacdn.com/js/tilda-zero-1.1.min.js",
					"https://static.tildacdn.com/js/tilda-popup-1.0.min.js",
					"https://static.tildacdn.com/js/tilda-zero-scale-1.0.min.js",
					"https://static.tildacdn.com/js/tilda-events-1.0.min.js",
					"https://static.tildacdn.com/js/tilda-stat-1.0.min.js",
				},
				CSS: []string{
					"https://static.tildacdn.com/css/tilda-grid-3.0.min.css",
					"https://static.tildacdn.com/ws/project54321/tilda-blocks-page12345.min.css?t=1734259633",
					"https://static.tildacdn.com/css/tilda-forms-1.0.min.css",
					"https://static.tildacdn.com/css/tilda-popup-1.1.min.css",
					"https://static.tildacdn.com/css/fonts-tildasans.css",
				},
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
				ctx:    context.Background(),
				pageID: "123",
			},
			stubFilename: "page.json",
			registerResponder: func(_ []byte) {
				url := fmt.Sprintf("%s/v1/getpage/?pageid=123&publickey=public&secretkey=secret", apiBaseUrl)

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
			got, err := c.GetPage(tt.args.ctx, tt.args.pageID)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClient_GetPageFull(t *testing.T) {
	type fields struct {
		config     *Config
		httpClient *http.Client
		baseURL    string
	}
	type args struct {
		ctx    context.Context
		pageID string
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		stubFilename      string
		registerResponder func(responseBody []byte)
		want              PageFull
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
				ctx:    context.Background(),
				pageID: "123",
			},
			stubFilename: "page_full.json",
			registerResponder: func(responseBody []byte) {
				url := fmt.Sprintf("%s/v1/getpagefull/?pageid=123&publickey=public&secretkey=secret", apiBaseUrl)

				httpmock.RegisterResponder(http.MethodGet, url,
					func(req *http.Request) (*http.Response, error) {
						resp := httpmock.NewBytesResponse(http.StatusOK, responseBody)

						return resp, nil
					},
				)
			},
			want: PageFull{
				ID:          "12345",
				ProjectID:   "54321",
				Title:       "Photography blog",
				Description: "Photographer's blog with tiled galleries and email subscription.",
				Date:        DateTime(time.Date(2024, 12, 15, 13, 20, 30, 0, time.UTC)),
				Sort:        20,
				Published:   1734259400,
				Alias:       "blog",
				Filename:    "page12345.html",
				HTML:        "<!DOCTYPE html> <html>...</html>",
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
				ctx:    context.Background(),
				pageID: "123",
			},
			stubFilename: "page_full.json",
			registerResponder: func(_ []byte) {
				url := fmt.Sprintf("%s/v1/getpagefull/?pageid=123&publickey=public&secretkey=secret", apiBaseUrl)

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
			got, err := c.GetPageFull(tt.args.ctx, tt.args.pageID)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClient_GetPageExport(t *testing.T) {
	type fields struct {
		config     *Config
		httpClient *http.Client
		baseURL    string
	}
	type args struct {
		ctx    context.Context
		pageID string
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		stubFilename      string
		registerResponder func(responseBody []byte)
		want              PageExport
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
				ctx:    context.Background(),
				pageID: "123",
			},
			stubFilename: "page_export.json",
			registerResponder: func(responseBody []byte) {
				url := fmt.Sprintf("%s/v1/getpageexport/?pageid=123&publickey=public&secretkey=secret", apiBaseUrl)

				httpmock.RegisterResponder(http.MethodGet, url,
					func(req *http.Request) (*http.Response, error) {
						resp := httpmock.NewBytesResponse(http.StatusOK, responseBody)

						return resp, nil
					},
				)
			},
			want: PageExport{
				ID:          "12345",
				ProjectID:   "54321",
				Title:       "Photography blog",
				Description: "Photographer's blog with tiled galleries and email subscription.",
				Date:        DateTime(time.Date(2024, 12, 15, 13, 20, 30, 0, time.UTC)),
				Sort:        20,
				Published:   1734259400,
				Alias:       "blog",
				Filename:    "page12345.html",
				HTML:        "<!--allrecords--> <div id=\"allrecords\" class=\"t-records\" data-hook=\"blocks-collection-content-node\" data-tilda-project-id=\"54321\" data-tilda-page-id=\"12345\" data-tilda-page-alias=\"blog\" data-tilda-formskey=\"qwerty\" data-tilda-cookie=\"no\" data-tilda-lazy=\"yes\" data-tilda-root-zone=\"one\"></div> <!--/allrecords-->",
				Images: []Image{
					{
						From: "https://static.tildacdn.com/img/tildacopy.png",
						To:   "tildacopy.png",
					}, {
						From: "https://static.tildacdn.com/img/tildacopy_black.png",
						To:   "tildacopy_black.png",
					},
				},
				JS: []JS{
					{
						From:  "https://static.tildacdn.com/js/tilda-polyfill-1.0.min.js",
						To:    "tilda-polyfill-1.0.min.js",
						Attrs: []string{"nomodule"},
					}, {
						From:  "https://static.tildacdn.com/js/tilda-scripts-3.0.min.js",
						To:    "tilda-scripts-3.0.min.js",
						Attrs: []string{"defer"},
					}, {
						From:  "https://static.tildacdn.com/js/lazyload-1.3.min.export.js",
						To:    "lazyload-1.3.min.export.js",
						Attrs: []string{"async"},
					}, {
						From: "https://static.tildacdn.com/js/tilda-stat-1.0.min.js",
						To:   "tilda-stat-1.0.min.js",
					},
				},
				CSS: []CSS{
					{
						From: "https://static.tildacdn.com/css/tilda-popup-1.1.min.css",
						To:   "tilda-popup-1.1.min.css",
					}, {
						From: "https://static.tildacdn.com/css/fonts-tildasans.css",
						To:   "fonts-tildasans.css",
					},
				},
				ProjectAlias: "mysiteqwerty",
				PageAlias:    "blog",
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
				ctx:    context.Background(),
				pageID: "123",
			},
			stubFilename: "page_export.json",
			registerResponder: func(_ []byte) {
				url := fmt.Sprintf("%s/v1/getpageexport/?pageid=123&publickey=public&secretkey=secret", apiBaseUrl)

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
			got, err := c.GetPageExport(tt.args.ctx, tt.args.pageID)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClient_GetPageFullExport(t *testing.T) {
	type fields struct {
		config     *Config
		httpClient *http.Client
		baseURL    string
	}
	type args struct {
		ctx    context.Context
		pageID string
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		stubFilename      string
		registerResponder func(responseBody []byte)
		want              PageExport
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
				ctx:    context.Background(),
				pageID: "123",
			},
			stubFilename: "page_export_full.json",
			registerResponder: func(responseBody []byte) {
				url := fmt.Sprintf("%s/v1/getpagefullexport/?pageid=123&publickey=public&secretkey=secret", apiBaseUrl)

				httpmock.RegisterResponder(http.MethodGet, url,
					func(req *http.Request) (*http.Response, error) {
						resp := httpmock.NewBytesResponse(http.StatusOK, responseBody)

						return resp, nil
					},
				)
			},
			want: PageExport{
				ID:          "12345",
				ProjectID:   "54321",
				Title:       "Photography blog",
				Description: "Photographer's blog with tiled galleries and email subscription.",
				Date:        DateTime(time.Date(2024, 12, 15, 13, 20, 30, 0, time.UTC)),
				Sort:        20,
				Published:   1734259400,
				Alias:       "blog",
				Filename:    "page12345.html",
				HTML:        "<!DOCTYPE html> <html>...</html>",
				Images: []Image{
					{
						From: "https://static.tildacdn.com/img/tildacopy.png",
						To:   "tildacopy.png",
					}, {
						From: "https://static.tildacdn.com/img/tildacopy_black.png",
						To:   "tildacopy_black.png",
					},
				},
				JS: []JS{
					{
						From:  "https://static.tildacdn.com/js/tilda-polyfill-1.0.min.js",
						To:    "tilda-polyfill-1.0.min.js",
						Attrs: []string{"nomodule"},
					}, {
						From:  "https://static.tildacdn.com/js/tilda-scripts-3.0.min.js",
						To:    "tilda-scripts-3.0.min.js",
						Attrs: []string{"defer"},
					}, {
						From:  "https://static.tildacdn.com/js/lazyload-1.3.min.export.js",
						To:    "lazyload-1.3.min.export.js",
						Attrs: []string{"async"},
					}, {
						From: "https://static.tildacdn.com/js/tilda-stat-1.0.min.js",
						To:   "tilda-stat-1.0.min.js",
					},
				},
				CSS: []CSS{
					{
						From: "https://static.tildacdn.com/css/tilda-popup-1.1.min.css",
						To:   "tilda-popup-1.1.min.css",
					}, {
						From: "https://static.tildacdn.com/css/fonts-tildasans.css",
						To:   "fonts-tildasans.css",
					},
				},
				ProjectAlias: "mysiteqwerty",
				PageAlias:    "blog",
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
				ctx:    context.Background(),
				pageID: "123",
			},
			stubFilename: "page_export_full.json",
			registerResponder: func(_ []byte) {
				url := fmt.Sprintf("%s/v1/getpagefullexport/?pageid=123&publickey=public&secretkey=secret", apiBaseUrl)

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
			got, err := c.GetPageFullExport(tt.args.ctx, tt.args.pageID)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
