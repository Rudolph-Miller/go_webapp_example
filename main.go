package main

import (
	"github.com/Rudolph-Miller/go_webapp_example/handlers/v1"
	"github.com/Rudolph-Miller/go_webapp_example/support"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var DBCONF = "./db/dbconf.yml"

type DBConf struct {
	Driver string
	Open   string
}

func main() {
	env := os.Getenv("GOENV")

	if env == "" {
		env = "development"
	}

	dbconfFile, _ := filepath.Abs(DBCONF)
	yamlFile, err := ioutil.ReadFile(dbconfFile)

	if err != nil {
		log.Fatal(err)
	}

	var envDbConf map[string]DBConf

	err = yaml.Unmarshal(yamlFile, &envDbConf)

	dbConf := envDbConf[env]

	db, err := gorm.Open(dbConf.Driver, dbConf.Open)

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &support.CustomContext{c, db}
			return h(cc)
		}
	})

	api := e.Group("/api")
	v1 := api.Group("/v1")

	handlersV1.UserGroup(v1)

	e.Logger.Fatal(e.Start(":1323"))
}
