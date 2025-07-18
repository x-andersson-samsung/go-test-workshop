package client

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setupTestServer(resp []byte) *httptest.Server {
	handlerFn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	})

	srv := httptest.NewServer(handlerFn)

	return srv
}

func TestClient_GetSize(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		srv := setupTestServer([]byte("9"))
		defer srv.Close()

		client := srv.Client()
		resp, err := client.Get(srv.URL)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("got status %d, want %d", resp.StatusCode, http.StatusOK)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}

		if string(body) != "Hello World!" {
			t.Errorf("got body %q, want %q", string(body), "Hello World!")
		}

	})
	t.Run("ok", func(t *testing.T) {
		input := []byte("123456789")
		srv := setupTestServer([]byte("9"))
		defer srv.Close()

		client := NewClient(srv.URL, srv.Client())

		got, err := client.GetSize(input)
		if err != nil {
			t.Fatalf("expected no error, got `%s`", err.Error())
		}
		if got != len(input) {
			t.Fatalf("expected %d, got %d", len(input), got)
		}
	})
	t.Run("error", func(t *testing.T) {
		input := []byte("123456789")
		srv := setupTestServer([]byte("NaN"))
		defer srv.Close()

		client := NewClient(srv.URL, srv.Client())

		_, err := client.GetSize(input)

		// Note: This is a wrong way of checking errors. You should not compare error messages.
		if !strings.Contains(err.Error(), "invalid response body") {
			t.Fatalf("expected invalid response error, got `%s`", err.Error())
		}

		// Note: Golang supports error wrapping which means you should not compare errors directly.
		// Use `errors.Is` instead
		if err != ErrInvalidResponse {
			t.Fatalf("expected invalid response error, got `%s`", err.Error())
		}

		if !errors.Is(err, ErrInvalidResponse) {
			t.Fatalf("expected invalid response error, got `%s`", err.Error())
		}
	})
}
