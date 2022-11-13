package main

import (
	"fmt"
	"log"

	"github.com/fikriyusrihan/golang-clean-arch/config"
	"github.com/fikriyusrihan/golang-clean-arch/infrastructure/datastore"
	"github.com/fikriyusrihan/golang-clean-arch/infrastructure/router"
	"github.com/fikriyusrihan/golang-clean-arch/registry"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/logger"
)

func main() {
	config.ReadConfig()

	db := datastore.NewDB()
	db.Logger.LogMode(logger.Info)
	
	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost:" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
