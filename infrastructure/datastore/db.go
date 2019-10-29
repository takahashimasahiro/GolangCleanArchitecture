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
	fmt.Println(DB)
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("success")

	// 外部のDBをConnectedSQLとして公開
	conn := ConnectedSql{DB: DB}

	return &conn
}

func (conn *ConnectedSql) Exec(cmd string, args ...interface{})(database.Result, error){
	result, err := conn.DB.Exec(cmd, args...)
	if err != nil {
		return nil, err
	}
	return &SqlResult{Result: result}, nil
}

type SqlResult struct {
	Result sql.Result
}