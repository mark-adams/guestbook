package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func TestHitCounter(t *testing.T) {
	server := httptest.NewServer(getMux())
	defer server.Close()

	startingVal := hitCounter

	_, err := http.Get(server.URL)
	if err != nil {
		t.Errorf("Request failed: %s", err)
	}

	expected := startingVal + 1
	if hitCounter != expected {
		t.Errorf("Hit counter does match expected value: %d != %d", hitCounter, expected)
	}
}

func TestAddSignature(t *testing.T) {
	server := httptest.NewServer(getMux())
	defer server.Close()

	startingSignatures := signatures[:]

	_, err := http.PostForm(server.URL, url.Values{
		"name": {"John Smith"},
	})

	if err != nil {
		t.Errorf("Request failed: %s", err)
	}

	if len(startingSignatures) == len(signatures) {
		t.Error("Expected a new signature but none found")
	}

	newSig := signatures[len(signatures)-1]

	if newSig.Name != "John Smith" {
		t.Errorf("New signature did not match expected name: %s", newSig.Name)
	}

	if newSig.Timestamp.Format("2016-01-02") != time.Now().Format("2016-01-02") {
		t.Errorf("new signature did not match expected timestamp: %s != %s",
			newSig.Timestamp.Format("2016-01-02"), time.Now().Format("2016-01-02"))
	}
}

func TestAddSignatureTrimsSpaces(t *testing.T) {
	server := httptest.NewServer(getMux())
	defer server.Close()

	startingSignatures := signatures[:]

	_, err := http.PostForm(server.URL, url.Values{
		"name": {"    John Smith     "}, // Note the extra whitespace
	})

	if err != nil {
		t.Errorf("Request failed: %s", err)
	}

	if len(startingSignatures) == len(signatures) {
		t.Error("Expected a new signature but none found")
	}

	newSig := signatures[len(signatures)-1]

	if newSig.Name != "John Smith" {
		t.Errorf("New signature did not match expected name: %s", newSig.Name)
	}
}
