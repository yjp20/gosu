package tmpl

import (
  "html/template"
  "net/http"
  "fmt"

  "github.com/youngjinpark20/gosu/modules/app"
  "github.com/youngjinpark20/gosu/modules/renderdata"

  "github.com/gobuffalo/packr"
)

type Tmpl struct {
  Name string
  cache *template.Template
  cached bool
}

func (t *Tmpl) Render(ctx *app.Context, r *http.Request, w http.ResponseWriter, data *renderdata.RenderData) {
  ctx.Debug(fmt.Sprintf("Rendering from: %s", r.URL.Path))
  data.AddDefaultValues(r, w, ctx)

  if !t.cached {
    box := packr.NewBox("../../templates")
    templates := template.Must(template.New("").Parse(box.String(t.Name)))
    box.WalkPrefix("base/", func(p string, f packr.File) error {
      templates.New(p).Parse(box.String(p))
      return nil
    })
    t.cache = templates
    if ctx.Settings.CacheTmpl {
      t.cached = true
    }
  }

  err := t.cache.ExecuteTemplate(w, "base", data)
  ctx.Error(err)
}

func New(name string) *Tmpl {
  return &Tmpl{
    Name: name,
    cached: false,
  }
}
