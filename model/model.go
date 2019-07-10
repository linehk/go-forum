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

// mysql 返回 MySQL 的 DSL 连接字符串。
func mysql() string {
	return fmt.Sprintf("%s:%s@%s(%s)/%s?charset=%s&parseTime=%s&loc=%s",
		dbc.User, dbc.Password, dbc.Protocol, dbc.Host, dbc.Name, dbc.Charset, dbc.ParseTime, dbc.Loc)
}

var db *sql.DB

// init 初始化 db 全局变量。
func init() {
	var err error
	db, err = sql.Open(dbc.Dialect, mysql())
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(dbc.MaxIdleConns)
	db.SetMaxOpenConns(dbc.MaxOpenConns)
}

// hash 用于对要插入到数据库的密码进行加密。
func hash(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hash)
}
