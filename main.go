package main

import (
	"api/initialisers"
	"api/internal/database"
	"api/internal/routes"
	"api/internal/utils"
	"flag"
)

func init() {
	initialisers.LoadEnvVariables()
}

func main() {
	database.DBCon = database.ConnectDB()

	// Parse flags
	seed := flag.Bool("seed", false, "seed the database")
	flag.Parse()

	if *seed {
		utils.SeedDB(database.DBCon)
		return
	}

	err := routes.NewServer()
	if err != nil {
		panic(err)
	}
}