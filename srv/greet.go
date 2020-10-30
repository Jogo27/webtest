package main

import (
  "io"
  "net/http"
)

func GreetHandler(wr http.ResponseWriter, req *http.Request) {
  _, session, err := CheckSession(req, "greet")
  if err != nil {
    http.Error(wr, "check session failed", http.StatusInternalServerError)
    return
  }

  if _, err := io.WriteString(wr, "Hi " + session.User() + "! Nice to see you!\n"); err != nil {
    http.Error(wr, "WriteString", http.StatusInternalServerError)
    return
  }

}
