package main

import (
	db "crud/internal/DB"
	dbsvc "crud/internal/DB-svc"
	"crud/internal/service"

	"github.com/go-sql-driver/mysql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	cfg := mysql.Config{User: "mysql", Passwd: "mysql", Net: "tcp", Addr: "svc-mysql:3306", DBName: "Store"}
	App := fiber.New()
	usersdb := db.NewDB("mysql", cfg.FormatDSN())
	err := usersdb.OpenDB()
	if err != nil {
		panic(err)
	}
	defer usersdb.CloseDB()
	dbservice := dbsvc.NewDB(usersdb)
	appservice := service.NewService(App, dbservice)
	log.Info("ok")
	appservice.Start()
}
