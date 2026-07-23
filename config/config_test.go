package config

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	wantPort := 3333
	wantEnv := "dev"
	t.Setenv("PORT", fmt.Sprint(wantPort))

	got, err := New()
	if err != nil {
		t.Fatalf("cannot create config: %v", err)
	}
	if got.Port != wantPort {
		t.Errorf("want: %d, got: %d", wantPort, got.Port)
	}
	if got.Env != wantEnv {
		t.Errorf("want: %s, got: %s", wantEnv, got.Env)
	}
}
