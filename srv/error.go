package main

import (
  "errors"
  "log"
  "net/http"
)

type HttpError struct {
  Code int
  msg string
  Detail string
}

func NewHttpError(code int, msg string, detail string) HttpError {
  return HttpError{code, msg, detail}
}

func (self HttpError) Error() string {
  return self.msg
}

// Send error to wr.
// If err is an HttpError, its code and msg are used in the HTPP response.
// Also log the error.
func SendError(wr http.ResponseWriter, err error) {
  var pError HttpError
  if errors.As(err, &pError) {
    http.Error(wr, pError.msg, pError.Code)
    log.Printf("%d %s: %s", pError.Code, pError.msg, pError.Detail)
  } else {
    http.Error(wr, err.Error(), http.StatusInternalServerError)
    log.Printf("%d %v", http.StatusInternalServerError, err)
  }
}
