package main

import (
	"log"

	"github.com/behzadayubifar/ecom/cmd/api"
	"github.com/behzadayubifar/ecom/configs"
	"github.com/behzadayubifar/ecom/database"
	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:                 configs.Envs.DBUser,
		Passwd:               configs.Envs.DBPassword,
		Addr:                 configs.Envs.DBAddress,
		DBName:               configs.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := database.NewMySQLStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	database.InitStorage(db)

	server := api.NewAPIServer(":8080", nil)
	if err := server.RUN(); err != nil {
		log.Fatal(err)
	}

}
