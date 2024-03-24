package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rcarvalho-pb/todo-app-go/internal/infra/configs"
	"github.com/rcarvalho-pb/todo-app-go/internal/router"
	"github.com/rcarvalho-pb/todo-app-go/pkg/db/repository/sqlite3db"
)

func main() {
	configs.InitEnvConfigs()
	db := sqlite3db.NewSqlite3db(configs.EnvConfigs.DBConnString)
	router := router.New(db)
	log.Printf("Server's up on Port: %d", configs.EnvConfigs.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configs.EnvConfigs.Port), router))
}
