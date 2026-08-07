package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Suiren91/GoExampleWebApp/config"
)

func TestNewMux(t *testing.T) {
	ctx := context.Background()
	cfg, err := config.New()
	if err != nil {
		t.Fatal(err)
	}
	sut, cleanup, err := NewMux(ctx, cfg)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(cleanup)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)

	sut.ServeHTTP(w, r)
	resp := w.Result()

	t.Cleanup(func() { _ = resp.Body.Close() })

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want status code: %d, got: %d", http.StatusOK, resp.StatusCode)
	}
	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	want := `{"status":"ok"}`
	if string(got) != want {
		t.Errorf("want: %q, got: %q", want, got)
	}
}
