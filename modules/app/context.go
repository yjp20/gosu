package app

import (
  "fmt"
  "log"
  "net/http"
  "os"
  "runtime/debug"

  "github.com/youngjinpark20/gosu/modules/settings"

  "github.com/jinzhu/gorm"
  "github.com/gorilla/sessions"
)

// Context provides the shared context that is used throughout the gosu
// application
type Context struct {
  InfoLog   *log.Logger
  DebugLog  *log.Logger
  ErrorLog  *log.Logger
  DB        *gorm.DB
  Settings  settings.Settings
  Store     *sessions.CookieStore
}

// Init intializes the Context with various logs
func (ctx *Context) Init() {
  ctx.InfoLog =  log.New(os.Stdout, "INFO\t",  log.Ldate | log.Ltime)
  ctx.DebugLog = log.New(os.Stdout, "DEBUG\t", log.Ldate | log.Ltime)
  ctx.ErrorLog = log.New(os.Stderr, "ERROR\t", log.Ldate | log.Ltime | log.Lshortfile)
}

// Info prints information to InfoLog
func (ctx *Context) Info(text string) {
  ctx.InfoLog.Println(text)
}

// Debug prints information to DebugLog
func (ctx *Context) Debug(text string) {
  ctx.DebugLog.Println(text)
}

// Error prints information to ErrorLog if the argument error is not nil
func (ctx *Context) Error(err error) {
  if err != nil {
    trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
    ctx.ErrorLog.Println(trace)
  }
}

// WriteError prints information to ErrorLog and the http.ResponseWriter if the
// argument error is not nil
func (ctx *Context) WriteError(w http.ResponseWriter, err error) {
  if err != nil {
    trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
    ctx.ErrorLog.Println(trace)
    http.Error(w, trace, http.StatusInternalServerError);
  }
}
