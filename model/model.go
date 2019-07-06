package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"

	"github.com/linehk/go-forum/config"
)

var dbc = config.Cfg.Database

func mysql() string {
	return fmt.Sprintf("%s:%s@%s(%s)/%s?charset=%s&parseTime=%s&loc=%s",
		dbc.User, dbc.Password, dbc.Protocol, dbc.Host, dbc.Name, dbc.Charset, dbc.ParseTime, dbc.Loc)
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open(dbc.Dialect, mysql())
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(dbc.MaxIdleConns)
	db.SetMaxOpenConns(dbc.MaxOpenConns)
}

func hash(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hash)
}
