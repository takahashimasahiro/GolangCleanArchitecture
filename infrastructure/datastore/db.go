package datastore

import (
	"database/sql"
	"fmt"
	"log"
	"../../config"
	_ "github.com/go-sql-driver/mysql"
)

type ConnectedSql struct {
	DB *sql.DB
}

func BootMysqlDB() *ConnectedSql{

	// configからDBの読み取り
	connectionCmd := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		config.Conf.Database.User,
		config.Conf.Database.Password,
		config.Conf.Database.Host,
		config.Conf.Database.Port,
		config.Conf.Database.Db,
	)
	fmt.Println(connectionCmd)

	var err error
	DB, err := sql.Open("mysql", connectionCmd)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// 外部のDBをConnectedSQLとして公開
	conn := ConnectedSql{DB: DB}

	return &conn
}