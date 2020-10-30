package main

import (
  "crypto/rand"
  "encoding/base64"
  "net/http"

  gs "github.com/gorilla/sessions"
)

const sessionName = "s";

var store = gs.NewCookieStore([]byte("flubluland_start_padding********"))

func randomShortId() (string, error) {

  b := make([]byte, 3)
  if _, err := rand.Read(b); err != nil {
    return "", err
  }

  return base64.URLEncoding.EncodeToString(b), nil
}

/*
func reduceEqual(a, b []string) {
  i := 0
  for la, lb := len(a), len(b); i < la && i < lb && a[i] == b[i]; i++ {
  }
  b = b[i:]
}
*/


func NewSession(req *http.Request) (session *gs.Session, err error) {
  session, err = store.New(req, sessionName)
  if err != nil {
    return
  }

  id, err := randomShortId()
  if err != nil {
    return
  }
  session.Values["id"] = id
  return
}

/*
func CheckSession(req *http.Request, action []string) (path []string, session gs.Session, err error) {
  session, err = store.Get(req, sessionName)
  if err != nil {
    return
  }


}
*/
