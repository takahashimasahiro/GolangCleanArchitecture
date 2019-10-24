package config

import (
)

type Config struct {
	Database
	Server
}
type Database struct {
	Host string
	Post string
	User string
	Password string
	Db string
}

type Server struct {
	Address string
}

var Conf Config

func LoadConfig(){
// MySQLで接続
}