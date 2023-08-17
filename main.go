package main

import (
	"RCSserver/routes"

	// "fmt"
	// "log"

	// dbs "github.com/conicuznhm/rcsdb/dbservice" //To create database
	// sfn "github.com/conicuznhm/rcsdb/dbservice/servicefn" //To use path and id fn
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main(){
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	routes.UserRoutes(e)
	
	port := ":8888"
	e.Logger.Fatal(e.Start(port))
	
	// to create database
	// if err := dbs.JSONdbCreate("","user"); err != nil{
	// 	fmt.Println(err)
	// 	log.Fatal(err)
	// }

	// to read path
	// if path, err := sfn.GetPath("user"); err != nil{
	// 	fmt.Printf("%v\n",err)
	// 	log.Fatal(err)
	// }else{
	// 	fmt.Println(path.Lastid)
	// }
}
