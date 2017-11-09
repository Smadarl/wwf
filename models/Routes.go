package models

import (
	"encoding/json"
	"fmt"
	"github.com/Smadarl/wwf/classes/DB"
	"github.com/Smadarl/wwf/classes/Response"
	"strconv"
	"time"

	Supernova "github.com/MordFustang21/supernova"
	jwt "github.com/dgrijalva/jwt-go"
)

//JWTSigningKey - key to use for signing tokens
var JWTSigningKey = []byte("SayTheSameThing")
var curAuthToken *jwt.Token

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
	nova.Get("/player/:id", getPlayer)
	nova.Post("/login", playerLogin)
	nova.Get("/friends", playerFriends)
	nova.Get("/game/:id", getGame)
	nova.Get("/games", getGames)
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
	path := string(req.Ctx.Path())
	if path != "/login" {
		var tokenStr string
		var header []byte
		header = req.Ctx.Request.Header.Peek("Authorization")
		tokenBytes := getTokenFromHeader(header)
		if len(tokenBytes) > 0 {
			tokenStr = string(tokenBytes)
			token, err := jwt.Parse(tokenStr, getJwtKey)
			if err != nil {
				req.SendJson(Response.Error(Response.ErrInvalidToken))
			} else {
				claims, _ := token.Claims.(jwt.MapClaims)
				pid, _ := strconv.Atoi(claims["jti"].(string))
				req.Ctx.SetUserValue("token", token)
				req.Ctx.SetUserValue("UserID", pid)
				curAuthToken = token
				next()
			}
		} else {
			req.SendJson(Response.Error(Response.ErrUnauthorized))
		}
	} else {
		next()
	}
}

func getPlayer(req *Supernova.Request) {
	req.Ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	req.Ctx.Response.Header.Set("Access-Control-Allow-Headers", "authorization")
	id, err := strconv.Atoi(req.RouteParams["id"])
	if err != nil {
		req.SendJson(Response.Error(Response.ErrDataInput))
	} else {
		conn := DB.GetConnection()
		defer conn.Close()
		player, err := GetPlayerByID(conn, id)
		if err != nil {
			req.SendJson(Response.Error(Response.ErrDataNotFound))
		} else {
			req.SendJson(player)
		}
	}
}

func playerLogin(req *Supernova.Request) {
	req.Ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	conn := DB.GetConnection()
	defer conn.Close()
	var loginForm loginForm
	err := json.Unmarshal(req.Body, &loginForm)
	if err != nil {
		fmt.Println(err)
		req.SendJson(Response.Error(Response.ErrInternal))
		return
	}
	player, err := GetPlayerByName(conn, loginForm.Username)
	if err != nil {
		req.SendJson(Response.Error(Response.ErrUserNotFound))
		return
	}
	if player.Password != loginForm.Password {
		req.SendJson(Response.Error(Response.ErrUserNotFound))
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
	req.SendJson(rt)
}

func playerFriends(req *Supernova.Request) {
	req.Response.Header.Set("Access-Control-Allow-Origin", "*")
	var token *jwt.Token
	token = req.UserValue("token").(*jwt.Token)
	claims, _ := token.Claims.(jwt.MapClaims)
	conn, _ := DB.GetConnection()
	defer conn.Close()
	pid, _ := strconv.Atoi(claims["jti"].(string))
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
