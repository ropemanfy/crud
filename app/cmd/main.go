package main

import (
	db "crud/internal/DB"
	dbsvc "crud/internal/DB-svc"
	"crud/internal/service"
	"crud/internal/transport"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
)

func main() {
	newLog := zerolog.New(os.Stdout)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT)

	newLog.Warn().Str("service", "crudNew").Msg("start service")
	defer newLog.Warn().Msg("shutdown server")

	cfg := mysql.Config{User: "mysql", Passwd: "mysql", Net: "tcp", Addr: "svc-mysql:3306", DBName: "Store"}
	// App := fiber.New()
	usersdb := db.NewDB("mysql", cfg.FormatDSN())
	err := usersdb.OpenDB()
	if err != nil {
		panic(err)
	}
	defer usersdb.CloseDB()
	dbservice := dbsvc.NewDB(usersdb)
	// appservice := service.NewService(App, dbservice)
	// log.Info("ok")
	// appservice.Start()

	svc := service.NewSvc(dbservice)
	options := []transport.Option{
		transport.Crud(transport.NewCrud(svc)),
	}
	trans := transport.New(newLog, options...).WithLog().WithMetrics()
	trans.ServeMetrics(newLog, "/", ":9003")

	go func() {
		newLog.Warn().Str("", ":9000").Msg("listen on")
		if err := trans.Fiber().Listen(":9000"); err != nil {
			newLog.Panic().Err(err).Msg("server error")
		}
	}()
	<-shutdown
}
