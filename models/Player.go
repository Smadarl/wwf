package models

import (
	"database/sql"
	"time"
)

// Player - struct to hold player data
type Player struct {
	ID          int
	Name        string
	Password    string
	Created     int64
	LastConnect int64
}

//SetName - Sets the player's username
func (p Player) SetName(name string) {
	p.Name = name
}

func (p Player) setCreated() {
	p.Created = time.Now().Unix()
}

func (p Player) setLastConnect() {
	p.LastConnect = time.Now().Unix()
}

// GetPlayerByName - Search db by player name, return Player
func GetPlayerByName(db *sql.DB, username string) (*Player, error) {
	var player Player
	err := db.QueryRow("select * from Player where Username = ?", username).Scan(&player.ID, &player.Name, &player.Created, &player.LastConnect, &player.Password)
	return &player, err
}

// GetPlayerByID - Search db by player ID, return Player
func GetPlayerByID(db *sql.DB, id int) (*Player, error) {
	var player Player
	err := db.QueryRow("select * from Player where PlayerID = ?", id).Scan(&player.ID, &player.Name, &player.Created, &player.LastConnect, &player.Password)
	return &player, err
}

//GetFriends - gets list of friends for player id
func GetFriends(db *sql.DB, id int) ([]*Player, error) {
	var fid int
	objects := make([]*Player, 0)
	rows, err := db.Query("select FriendID FROM Friend WHERE PlayerID = ?", id)
	for rows.Next() {
		err = rows.Scan(&fid)
		if err != nil {
			return objects, err
		}
		friend, _ := GetPlayerByID(db, fid)
		objects = append(objects, friend)
	}
	err = rows.Err()
	if len(objects) == 0 {
		return objects, sql.ErrNoRows
	} else if err != nil && err != sql.ErrNoRows {
		return objects, err
	}
	rows.Close()
	return objects, nil
}
