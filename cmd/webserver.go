package cmd

import (
  "net/http"

  "github.com/youngjinpark20/gosu/modules/app"
  "github.com/youngjinpark20/gosu/modules/sessions"
  "github.com/youngjinpark20/gosu/routes"

  "github.com/urfave/cli"
)

// WebServer exports the cli.Command that starts the gosu webserver.
func WebServer(ctx *app.Context) cli.Command {
  return cli.Command{
    Name: "server",
    Usage: "Start gosu server",
    Action: runServer(ctx),
    Flags: []cli.Flag{
      cli.StringFlag{
        Name: "port, p",
        Value: ":3000",
        Usage: "Override port number",
      },
    },
  }
}

func runServer(ctx *app.Context) cli.ActionFunc {
  return func(c *cli.Context) error {
    if c.IsSet("port") {
      ctx.Settings.HTTPPort = c.String("port")
    }
    sessions.Init(ctx)
    r := routes.GetRoutes(ctx)
    ctx.Info("Web Server listening on: " + ctx.Settings.HTTPPort)
    srv := &http.Server{
      Addr: ctx.Settings.HTTPPort,
      ErrorLog: ctx.ErrorLog,
      Handler: r,
    }
    err := srv.ListenAndServe()
    return err
  }
}

