package main

import (
	"fmt"

	"github.com/codegangsta/negroni"
	"github.com/serviceComputing1/server/model"
	"github.com/serviceComputing1/server/routes"
)

func main() {

	// open db should not be here, it should be moved into model modules
	err := model.Open()
	if err != nil {
		fmt.Println(err)
		return
	}

	server := NewServer()
	server.Run(":8080")

	model.Close()

}

func NewServer() *negroni.Negroni {

	n := negroni.Classic()

	mx := routes.NewRouter()

	//initRoutes(mx)

	n.UseHandler(mx)

	return n
}
