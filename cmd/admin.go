package cmd

import (
  "github.com/youngjinpark20/gosu/modules/app"
  "github.com/youngjinpark20/gosu/models"

  "github.com/urfave/cli"
)

// Admin exports the cli.Command pertaining to administration duties within gosu.
func Admin(ctx *app.Context) cli.Command {
  return cli.Command{
    Name: "admin",
    Usage: "Do administrative tasks",
    Subcommands: []cli.Command{
      {
        Name: "user",
        Usage: "User related tasks",
        Subcommands: []cli.Command{
          {
            Name: "deleteall",
            Usage: "Delete all users and related data",
            Action: deleteAllUsers(ctx),
          },
        },
      },
      {
        Name: "map",
        Usage: "Map related tasks",
        Subcommands: []cli.Command{
          {
            Name: "add",
            Usage: "Add a map by file name",
            Action: addMap(ctx),
            Flags: []cli.Flag{},
          },
          {
            Name: "deleteall",
            Usage: "Delete all maps and related data",
            Action: deleteAllMaps(ctx),
            Flags: []cli.Flag{},
          },
        },
      },
    },
  }
}

func deleteAllUsers(ctx *app.Context) cli.ActionFunc {
  return func(c *cli.Context) error {
    ctx.Info("Deleteing All User-related data: " + c.Args().Get(0))
    ctx.DB.Delete(&models.User{})
    return nil
  }
}

func deleteAllMaps(ctx *app.Context) cli.ActionFunc {
  return func(c *cli.Context) error {
    ctx.Info("Deleteing All Map-related data: " + c.Args().Get(0))
    ctx.DB.Delete(&models.MapSet{})
    ctx.DB.Delete(&models.Map{})
    ctx.DB.Delete(&models.OsuMap{})
    ctx.DB.Delete(&models.MapData{})
    return nil
  }
}

func addMap(ctx *app.Context) cli.ActionFunc {
  return func(c *cli.Context) error {
    err := models.AddMapSetByZipFileName(ctx, c.Args().Get(0))
    return err
  }
}
