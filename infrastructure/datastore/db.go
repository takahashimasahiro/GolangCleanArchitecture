package datastore

import (
	"database/sql"
	"fmt"
	"log"

	"../../config"
	"../../interface/database"

	_ "github.com/go-sql-driver/mysql"
)

type ConnectedSql struct {
	DB *sql.DB
}

func BootMysqlDB() *ConnectedSql {

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

	// 接続情報
	var err error
	DB, err := sql.Open("mysql", connectionCmd)
	if err != nil {
		log.Fatal(err)
	}
	// 接続確認
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

func (conn *ConnectedSql) Exec(cmd string, args ...interface{}) (database.Result, error) {
	result, err := conn.DB.Exec(cmd, args...)
	if err != nil {
		return nil, err
	}
	return &SqlResult{Result: result}, nil
}

func (conn *ConnectedSql) Query(cmd string, args ...interface{}) (database.Rows, error) {
	rows, err := conn.DB.Query(cmd, args...)
	if err != nil {
		return nil, err
	}
	return &SqlRows{Rows: rows}, nil
}

type SqlResult struct {
	Result sql.Result
}

func (r *SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r *SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SqlRows struct {
	Rows *sql.Rows
}

func (r SqlRows) Scans(ctr ...interface{}) error {
	return r.Rows.Scan(ctr...)
}

func (r SqlRows) Next() bool {
	return r.Rows.Next()
}

func (r SqlRows) Close() error {
	return r.Rows.Close()
}

type SqlRow struct {
	Row *sql.Row
}

func (r SqlRow) Scan(ctr ...interface{}) error {
	return r.Row.Scan(ctr...)
}
