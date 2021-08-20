package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestIndex(t *testing.T) {
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		t.Fatalf("Failed to access the server: %s", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	t.Logf("Raw response: %s", string(bodyBytes))
	var respBodyData response
	if err := json.Unmarshal(bodyBytes, &respBodyData); err != nil {
		t.Fatalf("Failed to unmarshal response: %s", err)
	}

	want := "How are you?"
	if respBodyData.Greeting != want {
		t.Errorf("Got: %s, Want: %s", respBodyData, want)
	}
}
