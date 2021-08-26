package server

import (
  "testing"
  "net/http/httptest"
  "strings"
  "fmt"
)


func TestHandleProduceBadJson(t *testing.T) {
  s := newHTTPServer()
  req := httptest.NewRequest("POST", "http://test.com", strings.NewReader(`{"test":"oi"}`))
  w := httptest.NewRecorder()
  s.handleProduce(w, req)

  resp := w.Result()
  fmt.Println(resp.StatusCode)
}
