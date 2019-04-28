package user

import (
  "net/http"

  "github.com/youngjinpark20/gosu/models"
  "github.com/youngjinpark20/gosu/modules/app"
  "github.com/youngjinpark20/gosu/modules/forms"
  "github.com/youngjinpark20/gosu/modules/renderdata"
  "github.com/youngjinpark20/gosu/modules/tmpl"
)

var (
  userHomeTemplate   = tmpl.New("user/home.tmpl")
  userLoginTemplate  = tmpl.New("user/login.tmpl")
  userSignupTemplate = tmpl.New("user/signup.tmpl")
)

// GetHome returns the http.Handler for GET requests for user/
func GetHome(ctx *app.Context) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    userHomeTemplate.Render(ctx, r, w, &renderdata.RenderData{})
  })
}

// GetLogin returns the http.Handler for GET requests for user/login
func GetLogin(ctx *app.Context) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    userLoginTemplate.Render(ctx, r, w, &renderdata.RenderData{
      Form: forms.New(nil),
    })
  })
}

// PostLogin returns the http.Handler for POST requests for user/login
func PostLogin(ctx *app.Context) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    var err error
    err = r.ParseForm()
    session, _ := ctx.Store.Get(r, "session")

    if session.Values["user"] != nil {
      http.Redirect(w, r, ctx.Settings.BaseURL, 302)
    }

    ctx.WriteError(w, err)
    f := forms.New(r.PostForm)
    f.Required("email", "pass")

    user := &models.User{}
    notExists := ctx.DB.Where(&models.User{Email: f.Get("email")}).Find(&user).RecordNotFound()

    if notExists {
      f.Errors["email"] = "email not found"
    } else {
      if !user.Authenticate(f.Get("pass")) {
        f.Errors["pass"] = "wrong password"
      }
    }

    if f.Valid() {
      session.Values["user"] = user
      session.AddFlash(renderdata.FlashData{
        Type: "success",
        Message: "User logged in succesfully",
      })
      err := session.Save(r, w)
      ctx.WriteError(w, err)

      http.Redirect(w, r, ctx.Settings.BaseURL, 302)
    } else {
      session.AddFlash(renderdata.FlashData{
        Type: "error",
        Message: "Error logging in",
      })
      session.Save(r, w)
      userLoginTemplate.Render(ctx, r, w, &renderdata.RenderData{
        Form: f,
      })
    }
  })
}

// GetSignup returns the http.Handler for GET requests for user/signup
func GetSignup(ctx *app.Context) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    userSignupTemplate.Render(ctx, r, w, &renderdata.RenderData{
      Form: forms.New(nil),
    })
  })
}

// PostSignup returns the http.Handler for POST requests for user/signup
func PostSignup(ctx *app.Context) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    ctx.WriteError(w, err)
    f := forms.New(r.PostForm)

    f.Required("username", "pass", "pass_confirm", "email", "email_confirm")
    f.MustLength("username", 1, 20)
    f.MustLength("pass", 8, 32)
    f.MustConfirm("pass", "pass_confirm")
    f.MustConfirm("email", "email_confirm")
    f.MustEmail("email")

    conflict := models.User{}

    ctx.DB.Where(&models.User{Name: f.Get("username")}).Find(&conflict)
    if len(conflict.ShortID) > 0 {
      f.Errors["username"] = "duplicate username"
    }

    ctx.DB.Where(&models.User{Email: f.Get("email")}).Find(&conflict)
    if len(conflict.ShortID) > 0 {
      f.Errors["email"] = "duplicate email"
    }

    if f.Valid() {
      user := models.User{
        Name: f.Get("username"),
        Email: f.Get("email"),
      }
      user.Init(ctx)
      user.SetPassword(f.Get("pass"))
      ctx.DB.Save(&user)
      http.Redirect(w, r, ctx.Settings.BaseURL, 302)
    } else {
      userSignupTemplate.Render(ctx, r, w, &renderdata.RenderData{
        Form: f,
      })
    }
  })
}

// GetLogout returns the http.Handler for GET requests for user/logout
func GetLogout(ctx *app.Context) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    session, _ := ctx.Store.Get(r, "session")
    session.Values["user"] = nil
    err := session.Save(r, w)
    ctx.WriteError(w, err)

    http.Redirect(w, r, ctx.Settings.BaseURL, 302)
  })
}

