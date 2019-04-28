package db

import (
  "github.com/youngjinpark20/gosu/modules/app"

  "github.com/jinzhu/gorm"
  // The pq driver loads some postgressql-specific features into the standard
  // go sql library.
  _ "github.com/lib/pq"
)

// Init appends a *gorm.DB connection to the *app.Context
func Init(ctx *app.Context) {
  var err error
	connStr := "user=" +      ctx.Settings.DB.User +
            " dbname=" +    ctx.Settings.DB.Name +
            " password=" +  ctx.Settings.DB.Password +
            " sslmode=" +   ctx.Settings.DB.SSLMode
  ctx.Info("Intializing postgres database connection")
  ctx.DB, err = gorm.Open("postgres", connStr)
  ctx.Error(err)
}
