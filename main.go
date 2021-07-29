package main

import (
	"TransferSaldo/packages"
	"TransferSaldo/src/repositories"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	packages.LoadMongo()
}

func main() {
	db := packages.LoadDatabase()
	repository := repositories.NewRepo(db)
	usecase := usecases.NewUC(repository)

}
