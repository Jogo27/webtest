package main

import (
  "errors"
  "net/http"
  "strings"
  "testing"
)

type MockResponseWriter struct {
  msg string
  code int
}

func (self MockResponseWriter) Header() http.Header {
  return http.Header{}
}

func (self *MockResponseWriter) Write(buf []byte) (int, error) {
  self.msg = self.msg + string(buf)
  return len(buf), nil
}

func (self *MockResponseWriter) WriteHeader(code int) {
  self.code = code
}

func TestSendError(t *testing.T) {
  
  t.Run("HttpError", func (t *testing.T) {

    mock := &MockResponseWriter{}
    err := NewHttpError(404, "Message", "Detail")

    SendError(mock, err)
    
    if mock.code != 404 {
      t.Errorf("Wrong code: got %d, want %d.", mock.code, 404)
    }
    if strings.TrimSpace(mock.msg) != "Message" {
      t.Errorf("Wrong message: got %s, want %s.", mock.msg, "Message")
    }
  })

  t.Run("Error", func (t *testing.T) {

    mock := &MockResponseWriter{}
    err := errors.New("Message")

    SendError(mock, err)
    
    if mock.code < 400 {
      t.Errorf("Wrong code: got %d, want < 400.", mock.code)
    }
    if mock.msg == "" {
      t.Errorf("Empty message")
    }
  })
}
