package tests

import (
	"database/sql"
	"fmt"
	"github.com/Smadarl/wwf/classes/Config"
	"github.com/Smadarl/wwf/models"
	"github.com/mattn/go-sqlite3"
)

func dbTest() {
	fmt.Println(sqlite3.Version())
	cfg := Config.GetConfig()
	fmt.Println(cfg)
	dbCfg := Config.GetDatabase()
	fmt.Println(dbCfg)
	db, err := sql.Open("sqlite3", dbCfg.Path)
	if err != nil {
		fmt.Println("Db failed")
	}
	defer db.Close()

	var player *models.Player
	player, err = models.GetPlayerByName(db, "smada")
	if err != nil {
		fmt.Println("query failed")
		fmt.Println(err)
	}
	fmt.Println(player)
}
