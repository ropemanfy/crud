package db

import (
	"database/sql"
	"log"
)

type StoreConnect interface {
	GetClient() (con *sql.DB)
	OpenDB() error
	CloseDB()
}

type storeCfg struct {
	dsn     string
	driver  string
	connect *sql.DB
}

func NewDB(driver string, dsn string) StoreConnect {
	return &storeCfg{driver: driver, dsn: dsn}
}

func (s *storeCfg) GetClient() (con *sql.DB) {
	con = s.connect
	return
}

func (s *storeCfg) OpenDB() error {
	db, err := sql.Open(s.driver, s.dsn)
	if err != nil {
		log.Println(err)
		return (err)
	}
	s.connect = db
	return err
}

func (s *storeCfg) CloseDB() {
	s.connect.Close()
}
