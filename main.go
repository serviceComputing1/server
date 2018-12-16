package main

import(
	//swapi "github.com/873421427/server/swapi"
	"fmt"
	"github.com/873421427/server/service"
	"github.com/873421427/server/model"
)

func main(){
	//var p swapi.People;
	
	err:=model.Open()
	if err !=nil{
		fmt.Println(err)
		return
	}

	//db_names :=[]string{"people","films","planets","vehicles","starships","species"}

	//swapi.Get_db_info("films")
	server := service.NewServer()
	server.Run(":8080")

	model.Close()

}

	/*
	swapi.Get_People();
	swapi.Get_Films();
	swapi.Get_Planets();
	swapi.Get_Vehicles();
	swapi.Get_StarShips();
	swapi.Get_Species();
	*/

	//swapi.GetPeople()