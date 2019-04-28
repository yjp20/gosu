package renderdata

import (
  "net/http"

  "github.com/youngjinpark20/gosu/modules/app"
  "github.com/youngjinpark20/gosu/modules/forms"
  "github.com/youngjinpark20/gosu/models"
)

// RenderData contains data used within templates
type RenderData struct {
  Settings    settingsData
  Title       string
  Description string
  MapSet      models.MapSet
  MapSets     []models.MapSet
  Map         models.Map
  Maps        []models.Map
  User        *models.User
  Users       []models.User
  Form        *forms.Form
  Flashes     []interface{}
}

type FlashData struct {
  Type string
  Message string
}

type settingsData struct {
  AppName string
  BaseURL string
}

// AddDefaultValues adds some common values used in rendering
func (d *RenderData) AddDefaultValues(r *http.Request, w http.ResponseWriter, ctx *app.Context) {
  session, _ := ctx.Store.Get(r, "session")
  d.Settings = settingsData{
    AppName: ctx.Settings.AppName,
    BaseURL: ctx.Settings.BaseURL,
  }
  d.Flashes = session.Flashes()
  if session.Values["user"] != nil {
    user := session.Values["user"].(models.User)
    d.User = &user
  }
  session.Save(r, w)
}
