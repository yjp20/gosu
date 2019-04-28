package models

import (
  "github.com/youngjinpark20/gosu/modules/app"
)

// Init intializes all the models by auto-migrating them through gorm
func Init(ctx *app.Context) {
  ctx.Info("Ayto-migrating gorm models")
  ctx.DB.AutoMigrate(&User{})
  ctx.DB.AutoMigrate(&Map{})
  ctx.DB.AutoMigrate(&OsuMap{})
  ctx.DB.AutoMigrate(&MapSet{})
  ctx.DB.AutoMigrate(&MapData{})
}
