package models

import (
  "archive/zip"
  "bytes"
  "io/ioutil"
  "os"
  "path"
  "strings"
  "time"

  "github.com/youngjinpark20/gosu/modules/app"

  "github.com/lib/pq"
  "github.com/teris-io/shortid"
  "github.com/jinzhu/gorm"
)

// MapSet contains information for a set of maps, like in osu where there are
// often "families" of individual maps and difficulties.
type MapSet struct {
  gorm.Model
  ShortID string
  Title string
  TitleOther string
  Artist string
  ArtistOther string
  Creator string
  SourceURL string
  ImageURL string
  Tags pq.StringArray `gorm:"type:varchar(64)[]"`
  Date time.Time
  MapIDs pq.StringArray `gorm:"type:varchar(64)[]"`
}

// Map contains information that pertains to maps in general, meant to
// generalize to osu and other rhythm games
type Map struct {
  gorm.Model
  ShortID string
  Title string
  TitleOther string
  Artist string
  ArtistOther string
  Creator string
  Version string
  Tags pq.StringArray `gorm:"type:varchar(64)[]"`
  Type string
  PlayCount int
  Date time.Time
  Active bool `gorm:"-"`

  ParentID string
  DataID string
}

// MapData is a meta struct meant for other types of maps to extend off of.
type MapData struct {
  gorm.Model
  ShortID string
}

// Init initializes the MapSet with a ShortID
func (ms *MapSet) Init(ctx *app.Context) error {
  var err error
  ms.ShortID, err = shortid.Generate()
  os.Mkdir(path.Join(ctx.Settings.DataPath, ms.ShortID), 0755)
  return err
}

// HandleFile handles a *zip.File and performs actoins depending on the
// file type
func (ms *MapSet) HandleFile(ctx *app.Context, file *zip.File) error {
  buf := new(bytes.Buffer)
  ext := path.Ext(file.Name)

  rc, err := file.Open()
  if err != nil {
    return err
  }
  defer rc.Close()

  buf.ReadFrom(rc)

  f, err := os.Create(path.Join(ctx.Settings.DataPath, ms.ShortID, file.Name))
  if err != nil {
    return err
  }
  defer f.Close()

  f.Write(buf.Bytes())

  switch strings.ToLower(ext) {
  case ".osu":
    err = ms.ParseOsuMap(ctx, buf)
    if err != nil {
      return err
    }
  case ".png",".jpg":
    ms.ImageURL = file.Name
  }
  return nil
}

// Init initializes the Map with a ShortID
func (m *Map) Init() error {
  var err error
  m.ShortID, err = shortid.Generate()
  return err
}

func AddMapSetByZipFileName(ctx *app.Context, name string) error {
  var ms MapSet
  var zipReader *zip.ReadCloser
  var err error
  ms.Init(ctx)

  zipReader, err = zip.OpenReader(name)
  if err != nil {
    return err
  }
  defer zipReader.Close()

  // save a copy of the source
  var input []byte
  input, err = ioutil.ReadFile(name)
  if err != nil {
    return err
  }
  err = ioutil.WriteFile(path.Join(ctx.Settings.DataPath, ms.ShortID, path.Base(name)), input, 0644)
  if err != nil {
    return err
  }

  for _, file := range zipReader.File {
    err := ms.HandleFile(ctx, file)
    if err != nil {
      return err
    }
  }
  ms.SourceURL = path.Base(name)
  ctx.DB.Create(&ms)
  return nil
}
