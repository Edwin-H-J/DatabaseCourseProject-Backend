package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pelletier/go-toml"
)

var Db *sql.DB

func init(){
	config, _ := toml.LoadFile("./config.toml")
	user := config.Get("database.user").(string)
	password := config.Get("database.password").(string)
	server := config.Get("database.server").(string)
	db_name := config.Get("database.db_name").(string)
	connectString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",user,password,server,db_name)
	var err error
	Db,err = sql.Open("mysql", connectString)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	Db.SetConnMaxLifetime(time.Minute * 3)
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
}