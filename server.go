package main

import (
	"fmt"
	"github.com/MordFustang21/supernova"
	"github.com/Smadarl/wwf/models"
)

func main() {
	nova := supernova.New()

	nova.All("", optionHeaders)

	models.RegisterRoutes(nova)

	err := nova.ListenAndServe(":8080")
	if err != nil {
		fmt.Println(err)
	}

}

func optionHeaders(req *supernova.Request) {
	fmt.Println(req.GetMethod() + string(req.Request.URI().Path()))
	req.Response.Header.Set("Access-Control-Allow-Origin", "*")
	req.Response.Header.Set("Access-Control-Allow-Headers", "Authorization")
}
