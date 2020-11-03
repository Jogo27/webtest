package main

import (
  "crypto/rand"
  "encoding/base64"
  "errors"
  "fmt"
  "log"
  "net/http"
  "path"
  "strings"
  "time"

  gs "github.com/gorilla/sessions"
)

type Session struct {
  session *gs.Session
}

func (self Session) stringValue(key string) string {
  if self.session == nil {
    return ""
  }

  unconverted, ok := self.session.Values[key]
  if !ok {
    return ""
  }

  ret, ok := unconverted.(string)
  if !ok {
    return ""
  }

  return ret
}

func (self Session) Id() string {
  return self.stringValue(sessionKeyId)
}

func (self Session) User() string {
  return self.stringValue(sessionKeyUser)
}

func (self Session) SetUser(user string) {
  self.session.Values[sessionKeyUser] = user
}

func (self Session) Save(wr http.ResponseWriter, req *http.Request) (err error) {
  err = self.session.Save(req, wr)
  if err != nil {
    log.Printf("Error saving session: %v", err)
  }
  return
}

const (
  sessionName = "s"
  sessionKeyId = "id"
  sessionKeyUser = "usr"
  sessionKeyDeadline = "dl"

  sessionMaxAge = 20
  sessionGraceTime = 2
)

var store = gs.NewCookieStore([]byte("flubluland_start_padding********"))

func init() {
  store.Options.MaxAge = sessionMaxAge
}

func randomShortId() (string, error) {

  b := make([]byte, 3)
  if _, err := rand.Read(b); err != nil {
    return "", err
  }

  return base64.URLEncoding.EncodeToString(b), nil
}

func NewSession(req *http.Request) (ret Session, err error) {
  ret.session, err = store.New(req, sessionName)
  if err != nil {
    return
  }

  id, err := randomShortId()
  if err != nil {
    return
  }
  ret.session.Values[sessionKeyId] = id
  ret.session.Values[sessionKeyDeadline] = time.Now().Unix() + sessionMaxAge + sessionGraceTime
  return
}

func CheckSession(req *http.Request, basePath string) (remainingPath []string, ret Session, err error) {
  ret.session, err = store.Get(req, sessionName)
  if ret.session.IsNew {
    err = NewHttpError(http.StatusForbidden, "Unauthorized", "session not found")
  }
  if err != nil {
    return
  }

  err = checkSessionDeadline(ret)
  if err != nil {
    return
  }

  remainingPath, err = checkSessionPath(basePath, req.URL.Path, ret.Id())
  return
}

func checkSessionDeadline(ret Session) error {
  if ret.session == nil {
    // Should never happend
    return errors.New("No session")
  }

  unconverted, ok := ret.session.Values[sessionKeyDeadline]
  if !ok {
    return NewHttpError(http.StatusForbidden, "Unauthorized", "no deadline in the session cookie")
  }

  deadline, ok := unconverted.(int64)
  if !ok {
    return NewHttpError(http.StatusForbidden, "Unauthorized", "the deadline in the session cookie has wrong type")
  }

  if time.Now().Unix() > deadline {
    return NewHttpError(http.StatusForbidden, "Unauthorized", "the session cookie has expired")
  }

  return nil
}
  
func checkSessionPath(basePath string, reqPath string, id string) (remainingPath []string, err error) {
  actionPath := append(strings.Split(path.Clean(basePath), "/"), id)
  cleanReqPath := path.Clean(reqPath)
  remainingPath = strings.Split(cleanReqPath, "/")
  if cleanReqPath[0] == '/' {
    remainingPath = remainingPath[1:]
  }

  lenRem := len(remainingPath)
  skip := 0
  for ; skip < lenRem && remainingPath[skip] != actionPath[0]; skip++ {
  }
  if skip == lenRem {
    err = NewHttpError(http.StatusForbidden, "Unauthorized",
      fmt.Sprintf("first path element %s not found", actionPath[0]))
    return
  }

  lenAction := len(actionPath)
  i := 1
  for ; i < lenAction && skip + i < lenRem && remainingPath[skip + i] == actionPath[i]; i++ {
  }
  if (i != lenAction) {
    err = NewHttpError(http.StatusForbidden, "Unauthorized",
      fmt.Sprintf("path element %s not found", actionPath[i]))
    return
  }

  remainingPath = remainingPath[skip+i:]
  return
}
