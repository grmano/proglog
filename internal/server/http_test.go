package server

import (
  "testing"
  "net/http/httptest"
  "strings"
  // "fmt"
)


func TestHandleProduce200(t *testing.T) {
  s := newHTTPServer()
  req := httptest.NewRequest("POST", "http://test.com", strings.NewReader(`{"test":"oi"}`))
  w := httptest.NewRecorder()
  s.handleProduce(w, req)

  resp := w.Result()
  if resp.StatusCode != 200 {
    t.Fatalf("Expected status code to be 200 got %v", resp.StatusCode)
  }
}

func TestHandleProduceBadJson(t *testing.T) {
  s := newHTTPServer()
  req := httptest.NewRequest("POST", "http://test.com", strings.NewReader(`{"test":"oi"`))
  w := httptest.NewRecorder()
  s.handleProduce(w, req)

  resp := w.Result()
  if resp.StatusCode != 400 {
    t.Fatalf("Expected status code to be 400 got %v", resp.StatusCode)
  }
}
