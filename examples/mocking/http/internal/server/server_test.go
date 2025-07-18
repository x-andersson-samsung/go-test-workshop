package server

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		srv := &Server{}
		requestBody := []byte("123")

		responseRecorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/", bytes.NewReader(requestBody))

		srv.ServeHTTP(responseRecorder, request)

		if responseRecorder.Code != http.StatusOK {
			t.Fatalf("expected 200, got `%d`", responseRecorder.Code)
		}
		if responseRecorder.Body.String() != "3" {
			t.Fatalf("expected `3`, got `%s`", responseRecorder.Body.String())
		}
	})
}
