package sessions

import (
  "encoding/gob"

  "github.com/youngjinpark20/gosu/models"
  "github.com/youngjinpark20/gosu/modules/app"
  "github.com/youngjinpark20/gosu/modules/renderdata"

  "github.com/gorilla/sessions"
)

// Init appends session store to the *app.Context
func Init(ctx *app.Context) {
  ctx.Info("Initializing sessions store")
  ctx.Store = sessions.NewCookieStore([]byte(ctx.Settings.Secret))
  gob.Register(renderdata.FlashData{})
  gob.Register(models.User{})
}
