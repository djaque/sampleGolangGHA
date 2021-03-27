package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(handler)
	hf.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)

	expected := `Hello Mundo!`
	actual := recorder.Body.String()
	assert.Equal(t, expected, actual, "Unexpected body: actual: %v expected: %v", actual, expected)
}
