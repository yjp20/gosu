package index

import (
  "net/http"

  "github.com/youngjinpark20/gosu/modules/app"
  "github.com/youngjinpark20/gosu/modules/renderdata"
  "github.com/youngjinpark20/gosu/modules/tmpl"
)

var (
  homeTemplate     = tmpl.New("home.tmpl")
  error404Template = tmpl.New("404.tmpl")
)

// GetHome returns the http.Handler for GET requests to /
func GetHome(ctx *app.Context) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    homeTemplate.Render(ctx, r, w, &renderdata.RenderData{})
  })
}

// GetError404 returns the http.Handler for NotFound errors
func GetError404(ctx *app.Context) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    error404Template.Render(ctx, r, w, &renderdata.RenderData{})
  })
}

// GetVersion returns the http.Handler for GET requests to /version, which
// prints the application version
func GetVersion(ctx *app.Context) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(ctx.Settings.AppVersion))
  })
}

// GetPing returns the http.Handler for GET requests to /ping, which prints
// "pong" that signifies that the server is alive
func GetPing(ctx *app.Context) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("ping"))
  })
}
