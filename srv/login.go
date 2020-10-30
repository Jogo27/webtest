package main

import (
  "io"
  "net/http"
)

func LoginHandler(wr http.ResponseWriter, req *http.Request) {
  userName := req.FormValue("login")
  if userName == "" {
    http.Error(wr, "No login", http.StatusBadRequest)
    return
  }

  session, err := NewSession(req)
  if err != nil {
    http.Error(wr, "NewSession", http.StatusInternalServerError)
    return
  }

  session.SetUser(userName)
  if session.Save(wr, req) != nil {
    http.Error(wr, "session.Save", http.StatusInternalServerError)
    return
  }

  id := session.Id()
  if _, err := io.WriteString(wr, id + "\n"); err != nil {
    http.Error(wr, "WriteString", http.StatusInternalServerError)
    return
  }
}
