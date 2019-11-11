package main

import (
	"fmt"
	"log"

	"../config"
	"../infrastructure/datastore"
	"../infrastructure/router"
	"../infrastructure/server"
	"../interface/controllers"
)

func main() {
	// DB情報サーバー情報読み込み
	config.LoadConfig()
	fmt.Println(config.Conf.Db)
	// DB起動
	connectedDB := datastore.BootMysqlDB()
	fmt.Println(connectedDB)
	// interactorを作成
	interactor := controllers.NetInteractor(connectedDB)
	// AppHandlerの取得
	appController := interactor.NewAppController()
	// Routerの起動
	serv := server.New()
	router.BootRouter(serv, appController)
	// DBのClose
	defer func() {
		if err := connectedDB.DB.close(); err != nil {
			log.Fatal(fmt.Sprintf("Failed to close: %v", err))
		}
	}()
	fmt.Println("success")
	// Server Start
	serv.Start(config.Conf.Server.Address)
}
