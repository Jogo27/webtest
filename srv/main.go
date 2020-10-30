package main

import (
  "crypto/rand"
  "encoding/base64"
  "io"
  "log"
  "net/http"
)

func RandomShortId() (string, error) {

  b := make([]byte, 3)
  if _, err := rand.Read(b); err != nil {
    return "", err
  }

  return base64.URLEncoding.EncodeToString(b), nil
}

func dumpHandler(resp http.ResponseWriter, req *http.Request) {
  log.Printf("%s %s", req.Method, req.URL.Path)
  
  id, err := RandomShortId()
  if err != nil {
    http.Error(resp, "RandomShortId", http.StatusInternalServerError)
  }
  if _, err := io.WriteString(resp, id + "\n"); err != nil {
    http.Error(resp, "WriteString", http.StatusInternalServerError)
  }
}

func main() {
  http.HandleFunc("/", dumpHandler)
  http.HandleFunc("/login", LoginHandler)
  http.ListenAndServe(":8080", nil)
}
