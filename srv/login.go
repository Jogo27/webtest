package main

import (
  "io"
  "log"
  "net/http"
)

func LoginHandler(wr http.ResponseWriter, req *http.Request) {
  userName := req.FormValue("login")
  if userName == "" {
    log.Println("No login")
    http.Error(wr, "No login", http.StatusBadRequest)
    return
  }

  session, err := NewSession(req)
  if err != nil {
    SendError(wr, err)
    return
  }

  session.SetUser(userName)
  if err = session.Save(wr, req); err != nil {
    SendError(wr, err)
    return
  }

  id := session.Id()
  if _, err := io.WriteString(wr, id + "\n"); err != nil {
    SendError(wr, err)
    return
  }

  log.Printf("Logged as " + userName)
}
