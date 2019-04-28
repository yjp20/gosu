package maps

import (
  "net/http"

  "github.com/youngjinpark20/gosu/models"
  "github.com/youngjinpark20/gosu/modules/app"
  "github.com/youngjinpark20/gosu/modules/tmpl"
  "github.com/youngjinpark20/gosu/modules/renderdata"

  "github.com/gorilla/mux"
)

var (
  mapsHomeTemplate = tmpl.New("maps/home.tmpl")
  mapsMapTemplate  = tmpl.New("maps/map.tmpl")
)

// GetHome returns the http.Handler for GET requests to /maps
func GetHome(ctx *app.Context) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    data := renderdata.RenderData{}
    ctx.DB.Find(&data.MapSets)
    mapsHomeTemplate.Render(ctx, r, w, &data)
  })
}

// GetMapSet returns the http.Handler for GET requests to /maps/{map_set_id}
// and /map/{map_set_id}/{map_collection_id}
func GetMapSet(ctx *app.Context) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    data := renderdata.RenderData{}

    ctx.DB.Where("short_id = ?", vars["map_set_id"]).Find(&data.MapSet)

    for _, mapID := range ([]string)(data.MapSet.MapIDs) {
      m := models.Map{}
      ctx.DB.Where("short_id = ?", mapID).Find(&m)
      if mapID == vars["map_id"] {
        m.Active = true
        data.Map = m
      }
      data.Maps = append(data.Maps, m)
    }

    mapsMapTemplate.Render(ctx, r, w, &data)
  })
}

