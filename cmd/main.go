package main

import (
	"../config"
	"fmt"
)

func main() {
	// DB情報サーバー情報読み込み
	config.LoadConfig()
	fmt.Println(config.conf.Db)
	// DB起動
	
}