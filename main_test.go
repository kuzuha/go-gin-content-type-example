package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRequests(t *testing.T) {
	g := gin.Default()
	router(g)

	tests := []struct {
		name        string
		method      string
		contentType string
		url         string
		body        string
		expected    interface{}
	}{
		{
			name:        "from json",
			method:      "POST",
			contentType: "application/json",
			url:         "/json",
			body:        `{"name": "json user"}`,
			expected:    `{"name":"json user"}`,
		},
		{
			name:        "from form",
			method:      "POST",
			contentType: "application/x-www-form-urlencoded",
			url:         "/form",
			body:        `name=form%20user`,
			expected:    `{"name":"form user"}`,
		},
		{
			name:        "from query",
			method:      "GET",
			contentType: "",
			url:         "/query?name=query%20user",
			body:        ``,
			expected:    `{"name":"query user"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, e := http.NewRequest(tt.method, tt.url, strings.NewReader(tt.body))
			if assert.NoError(t, e) {
				req.Header.Set("Content-Type", tt.contentType)
				g.ServeHTTP(w, req)
				if assert.Equal(t, 200, w.Code) {
					b, e := io.ReadAll(w.Body)
					if assert.NoError(t, e) {
						assert.Equal(t, tt.expected, string(b))
					}
				}
			}
		})
	}

}
