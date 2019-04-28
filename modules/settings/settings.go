package settings

import (
  "os"
  "path"

  ini "gopkg.in/ini.v1"
)

// Settings contains configuration information that is used within the gosu app
type Settings struct {
  // Meta - not defined by config file
  AppName     string
  AppVersion  string
  AppPath     string  // path to app dir
  DataPath    string  // path to data dir
  ConfigPath  string  // path to config file
  Secret      string

  // Config - defined by config file
  CacheTmpl   bool
  BaseURL     string  // example.com/{baseurl}
  HTTPPort    string  // HTTP port like ":3000"

  DB dbSettings
}

type dbSettings struct {
  Name        string  // Database name
  User        string  // Database username
  Password    string  // Database password
  SSLMode     string  // Postgres sql SSL mode
}

func readConfig(configPath string) *ini.File {
  cfg := ini.Empty()
  cfg.Append(configPath)
  return cfg
}

// Init intializes a Settings struct by reading from a configuration file and
// asserting default values in the absence of values
func (s *Settings) Init() {
  s.AppPath = os.Getenv("GOSU_APP_PATH")
  if len(s.ConfigPath) == 0 {
    s.ConfigPath = path.Join(s.AppPath, "conf.ini")
  }

  cfg := readConfig(s.ConfigPath)

  s.DataPath = cfg.Section("").Key("DataPath").MustString("data")
  s.Secret = cfg.Section("").Key("Secret").MustString("I'll show you mine if you show me yours")
  s.BaseURL = cfg.Section("").Key("BaseURL").MustString("/")
  s.CacheTmpl = cfg.Section("").Key("CacheTmpl").MustBool(false)
  s.HTTPPort = cfg.Section("").Key("HTTPPort").MustString(":3000")
  s.DB.Name = cfg.Section("DB").Key("Name").MustString("gosu")
  s.DB.User = cfg.Section("DB").Key("User").MustString("gosu")
  s.DB.Password = cfg.Section("DB").Key("Password").MustString("password")
  s.DB.SSLMode = cfg.Section("DB").Key("SSLMode").MustString("disable")
}
