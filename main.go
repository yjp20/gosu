// Copyright 2019 Young Jin Park. All rights reserved

package main

import (
  "os"
  "runtime"

  "github.com/youngjinpark20/gosu/cmd"
  "github.com/youngjinpark20/gosu/models"
  "github.com/youngjinpark20/gosu/modules/app"
  "github.com/youngjinpark20/gosu/modules/db"

  "github.com/urfave/cli"
)

func main() {
  ctx := &app.Context{}
  ctx.Init()
  ctx.Settings.AppName = "gosu"
  ctx.Settings.AppVersion = "0.0.1"
  ctx.Settings.Init()
  db.Init(ctx)
  models.Init(ctx)

  cliApp := cli.NewApp()
  cliApp.Name = "gosu"
  cliApp.Usage = "Go-powered Osu on the web"
  cliApp.Version = ctx.Settings.AppVersion + builtWith()
  cliApp.Commands = []cli.Command{
    cmd.WebServer(ctx),
    cmd.Admin(ctx),
  }

  err := cliApp.Run(os.Args)
  ctx.Error(err)
}

func builtWith() string {
  return " built with " + runtime.Version()
}
