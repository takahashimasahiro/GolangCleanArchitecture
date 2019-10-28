package main

import (
	"../config"
	"../infrastructure/datastore"
	"fmt"
)

func main() {
	// DB情報サーバー情報読み込み
	config.LoadConfig()
	fmt.Println(config.Conf.Db)
	// DB起動
	connectedDB := datastore.BootMysqlDB()
	fmt.Println(connectedDB)
	
	fmt.Println("success")
}