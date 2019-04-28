package routes

import (
  "net/http"

  "github.com/youngjinpark20/gosu/modules/app"
  "github.com/youngjinpark20/gosu/routes/index"
  "github.com/youngjinpark20/gosu/routes/maps"
  "github.com/youngjinpark20/gosu/routes/user"

  "github.com/gorilla/mux"
  "github.com/gobuffalo/packr"
)

// GetRoutes returns a mux router that contains all the routes for the gosu app
func GetRoutes(ctx *app.Context) *mux.Router {
  public := packr.NewBox("../public")
  data := packr.NewBox("../data")

  ctx.Debug(ctx.Settings.BaseURL)

  r := mux.NewRouter()
  m := r.PathPrefix("/maps").Subrouter()
  u := r.PathPrefix("/user").Subrouter()

  r.NotFoundHandler = http.Handler(index.GetError404(ctx))
  r.PathPrefix("/public/").Handler(
    http.StripPrefix(
      ctx.Settings.BaseURL,
      http.StripPrefix(
        "public/",
        http.FileServer(public),
      ),
    ),
  )
  r.PathPrefix("/data/").  Handler(http.StripPrefix("/data/", http.FileServer(data)))

  // routes start here
  r.Handle("/", index.GetHome(ctx)).                           Methods("GET")
  r.Handle("/ping", index.GetPing(ctx)).                       Methods("GET")
  r.Handle("/version", index.GetVersion(ctx)).                 Methods("GET")

  // maps
  m.Handle("", maps.GetHome(ctx)).                             Methods("GET")
  m.Handle("/{map_set_id}", maps.GetMapSet(ctx)).              Methods("GET")
  m.Handle("/{map_set_id}/{map_id}", maps.GetMapSet(ctx)).     Methods("GET")

  // user
  u.Handle("", user.GetHome(ctx)).                             Methods("GET")
  u.Handle("/login", user.GetLogin(ctx)).                      Methods("GET")
  u.Handle("/login", user.PostLogin(ctx)).                     Methods("POST")
  u.Handle("/signup", user.GetSignup(ctx)).                    Methods("GET")
  u.Handle("/signup", user.PostSignup(ctx)).                   Methods("POST")
  u.Handle("/logout", user.GetLogout(ctx)).                  Methods("GET")

  return r
}
