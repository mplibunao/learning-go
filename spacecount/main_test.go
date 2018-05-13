package main

import "testing"

type testWebRequest struct {
}

// Fake the Http Request here
func (testWebRequest) FetchBytes(url string) []byte {
	return []byte(`{"number": 2}`)
}

func TestGetAstronauts(t *testing.T) {
	amount := GetAstronauts(testWebRequest{})
	if amount != 2 {
		t.Errorf("People in space, got: %d, want: %d.", amount, 2)
	}
}
