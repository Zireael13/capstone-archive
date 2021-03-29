package main

import (
	"github.com/Zireael13/capstone-archive/server/internal/auth"
	"github.com/Zireael13/capstone-archive/server/internal/db"
	"github.com/Zireael13/capstone-archive/server/internal/envs"
	"github.com/Zireael13/capstone-archive/server/internal/serve"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env.local")
	_ = godotenv.Load(".env")

	port := envs.GetPort()

	orm := db.CreateDefaultDatabaseClient()
	argon := auth.CreateArgon()

	db.LoadSampleData(orm)

	g := serve.CreateServer(orm, argon)

	serve.RunServer(g, port)

}
