package http_test

import (
	"context"
	"strings"
	"testing"

	http "github.com/github.com/raulb/conduit-connector-http"
)

func TestConfigureSource_FailsWhenConfigEmpty(t *testing.T) {
	con := http.Source{}
	err := con.Configure(context.Background(), make(map[string]string))
	if err == nil {
		t.Error("expected error for missing config params")
	}

	if strings.HasPrefix(err.Error(), "config is invalid:") {
		t.Errorf("expected error to be about missing config, got %v", err)
	}
}

func TestTeardownSource_NoOpen(t *testing.T) {
	con := http.NewSource()
	err := con.Teardown(context.Background())
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
