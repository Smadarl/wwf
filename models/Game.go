package models

import (
	"fmt"
	"github.com/Smadarl/wwf/classes/DB"
	"github.com/Smadarl/wwf/classes/Response"
	"strconv"

	Supernova "github.com/MordFustang21/supernova"
)

//Turn - holds a single turn
type Turn struct {
	PlayerID int
	Guess    string
	AtTime   int
	Score    int
}

//Game - holds game metadata
type Game struct {
	GameID        int
	Opponent      string
	Word          string
	MaxLength     int
	MaxRecurrance int
	Started       int
	StartedBy     int
	WhoseTurn     int
	LastTurn      int
	Turns         []*Turn
	Status        string
}

func getGames(req *Supernova.Request) {
	req.Response.Header.Set("Access-Control-Allow-Origin", "*")
	var game Game
	db, _ := DB.GetConnection()
	objects := make([]*Game, 0)
	rows, err := db.Query(`SELECT g.*, o.Username, gp.Word, gp.LastTime 
						   FROM Game g 
						     JOIN Game_Player gp ON gp.GameID = g.GameID 
							 JOIN Game_Player gp2 ON gp2.GameID = g.GameID AND gp2.PlayerID != gp.PlayerID 
							 JOIN Player o ON o.PlayerID = gp2.PlayerID 
						   WHERE gp.PlayerId = ?`, req.UserValue("UserID"))
	for rows.Next() {
		err = rows.Scan(&game.GameID, &game.MaxLength, &game.MaxRecurrance, &game.Started, &game.StartedBy, &game.WhoseTurn, &game.Status, &game.Opponent, &game.Word, &game.LastTurn)
		if err != nil {
			return
		}
		objects = append(objects, &game)
	}
	req.JSON(200, objects)
}

func getGame(req *Supernova.Request) {
	id, err := strconv.Atoi(req.RouteParam("id"))
	if err != nil {
		req.JSON(500, Response.Error(Response.ErrDataInput))
		return
	}
	var game Game
	db, _ := DB.GetConnection()
	row := db.QueryRow("select g.*, gp.Word, gp.LastTime from Game g JOIN Game_Player gp ON gp.GameID = g.GameID WHERE g.GameID = ?", id)
	err = row.Scan(&game.GameID, &game.MaxLength, &game.MaxRecurrance, &game.Started, &game.StartedBy, &game.WhoseTurn, &game.Status, &game.Word, &game.LastTurn)

	game.Turns = make([]*Turn, 0)
	fmt.Println(fmt.Sprintf("SELECT pm.PlayerID, pm.Guess, pm.AtTime, pm.Result FROM PlayerMove pm WHERE pm.GameID = %d AND pm.PlayerID = %d", id, req.UserValue("UserID")))
	rows, err := db.Query("SELECT pm.PlayerID, pm.Guess, pm.AtTime, pm.Result FROM PlayerMove pm WHERE pm.GameID = ? AND pm.PlayerID = ?", id, req.UserValue("UserID"))
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var turn Turn
		err = rows.Scan(&turn.PlayerID, &turn.Guess, &turn.AtTime, &turn.Score)
		fmt.Println(turn)
		if err != nil {
			return
		}
		game.Turns = append(game.Turns, &turn)
	}
	rows.Close()

	req.JSON(200, game)
}

func takeTurn(req *Supernova.Request) {

}
