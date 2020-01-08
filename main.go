package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/letanthang/nc_crm/config"
	"github.com/letanthang/nc_crm/db"
	mw "github.com/letanthang/nc_crm/middleware"
	"github.com/letanthang/nc_crm/route"
)

func main() {
	fmt.Printf("config app: %+v", config.Config)
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(mw.SimpleLogger())
	route.All(e)

	levels, err := db.GetLevels()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", levels)
	log.Println(e.Start(":9090"))
}
