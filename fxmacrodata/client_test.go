package fxmacrodata

import "testing"

func TestBuildURL(t *testing.T) {
	client := NewClient("test-key")
	got := client.buildURL("/forex/aud/usd", nil)
	want := "https://api.fxmacrodata.com/v1/forex/aud/usd?api_key=test-key"
	if got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}
