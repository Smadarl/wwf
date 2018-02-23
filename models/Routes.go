package models

import (
	//	"encoding/json"
	"fmt"
	"github.com/Smadarl/wwf/classes/DB"
	"github.com/Smadarl/wwf/classes/Response"
	"strconv"
	"time"

	Supernova "github.com/MordFustang21/supernova"
	"github.com/dgrijalva/jwt-go"
)

//JWTSigningKey - key to use for signing tokens
var JWTSigningKey = []byte("SayTheSameThing") // You bought me a dog?

//RequestError - Error message to send back
type RequestError struct {
	message string
}

func (r *RequestError) Error() string {
	return r.message
}

type loginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//TokenResponse - struct to hold token json
type TokenResponse struct {
	Token   string
	Expires int64
}

// RegisterRoutes - setup routes in supernova
func RegisterRoutes(nova *Supernova.Server) {
	nova.Get("/api/player/:id", getPlayer)
	nova.Post("/api/login", playerLogin)
	nova.Get("/api/friends", playerFriends)
	nova.Get("/api/game/:id", getGame)
	nova.Get("/api/games", getGames)
	//	nova.Post(VERSION + "/player/:id", postListing)
	//	nova.Put(VERSION + "/player/:id", putListing)
	//	nova.Delete(VERSION + "/player/:id", deleteListing)
	nova.Use(authenticate)
}

// Authenticate - Middleware method to verify request auth
func authenticate(req *Supernova.Request, next func()) {
	if req.GetMethod() == "OPTIONS" {
		next()
		return
	}
	path := string(req.Path())
	if path != "/api/login" {
		var tokenStr string
		var header []byte
		header = req.Request.Header.Peek("Authorization")
		tokenBytes := getTokenFromHeader(header)
		if len(tokenBytes) > 0 {
			tokenStr = string(tokenBytes)
			token, err := jwt.Parse(tokenStr, getJwtKey)
			if err != nil {
				req.JSON(402, Response.Error(Response.ErrInvalidToken))
			} else {
				claims, _ := token.Claims.(jwt.MapClaims)
				pid, _ := strconv.Atoi(claims["jti"].(string))
				req.SetUserValue("token", token)
				req.SetUserValue("UserID", pid)
				next()
			}
		} else {
			req.JSON(401, Response.Error(Response.ErrUnauthorized))
		}
	} else {
		next()
	}
}

func getPlayer(req *Supernova.Request) {
	req.Response.Header.Set("Access-Control-Allow-Origin", "*")
	req.Response.Header.Set("Access-Control-Allow-Headers", "authorization")
	id, err := strconv.Atoi(req.RouteParam("id"))
	if err != nil {
		req.JSON(200, Response.Error(Response.ErrDataInput))
	} else {
		conn, _ := DB.GetConnection()
		defer conn.Close()
		player, err := GetPlayerByID(conn, id)
		if err != nil {
			req.JSON(200, Response.Error(Response.ErrDataNotFound))
		} else {
			req.JSON(200, player)
		}
	}
}

func playerLogin(req *Supernova.Request) {
	req.Response.Header.Set("Access-Control-Allow-Origin", "*")
	conn, _ := DB.GetConnection()
	defer conn.Close()
	var loginForm loginForm
	err := req.ReadJSON(&loginForm)
	//	err := json.Unmarshal(req.Body, &loginForm)
	if err != nil {
		fmt.Println(err)
		req.JSON(500, Response.Error(Response.ErrInternal))
		return
	}
	player, err := GetPlayerByName(conn, loginForm.Username)
	if err != nil {
		req.JSON(401, Response.Error(Response.ErrUserNotFound))
		return
	}
	if player.Password != loginForm.Password {
		req.JSON(401, Response.Error(Response.ErrUserNotFound))
		return
	}
	t := time.Now()
	t = t.Add(time.Hour * 24 * 5)
	claims := &jwt.StandardClaims{
		ExpiresAt: t.Unix(),
		Issuer:    "smada.com",
		Id:        strconv.Itoa(player.ID),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(JWTSigningKey)
	rt := TokenResponse{
		Token:   ss,
		Expires: t.Unix(),
	}
	req.JSON(200, rt)
}

func playerFriends(req *Supernova.Request) {
	req.Response.Header.Set("Access-Control-Allow-Origin", "*")
	var pid = req.UserValue("UserID").(int)
	conn, _ := DB.GetConnection()
	defer conn.Close()
	friends, err := GetFriends(conn, pid)
	if err != nil {
		fmt.Println(err)
	}
	req.JSON(200, friends)
}

func getJwtKey(token *jwt.Token) (interface{}, error) {
	return []byte(JWTSigningKey), nil
}

func getTokenFromHeader(header []byte) []byte {
	if len(header) > 8 {
		//Grab token based on position
		return header[7:]
	}
	return make([]byte, 0)
}
