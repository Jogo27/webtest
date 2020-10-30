package main

import (
  "io"
  "net/http"
)

func LoginHandler(wr http.ResponseWriter, req *http.Request) {
  userName := req.FormValue("login")
  if userName == "" {
    http.Error(wr, "No login", http.StatusBadRequest)
  }

  session, err := NewSession(req)
  if err != nil {
    http.Error(wr, "NewSession", http.StatusInternalServerError)
  }

  session.Values["user"] = userName
  if session.Save(req, wr) != nil {
    http.Error(wr, "session.Save", http.StatusInternalServerError)
  }

  id, ok := session.Values["id"].(string)
  if !ok {
    http.Error(wr, "conversion", http.StatusInternalServerError)
  }
  if _, err := io.WriteString(wr, id + "\n"); err != nil {
    http.Error(wr, "WriteString", http.StatusInternalServerError)
  }
}
