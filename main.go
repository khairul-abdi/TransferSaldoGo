package main

import (
	"TransferSaldo/packages"
	"TransferSaldo/src/controllers"
	"TransferSaldo/src/repositories"
	"TransferSaldo/src/routers"
	"TransferSaldo/src/usecases"
	"fmt"
	"log"
	"net/http"
	"os"

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
	ctrl := controllers.NewCtrl(usecase)
	route := routers.NewRouter(ctrl).RouterGroup()
	log.Printf("[SERVER] starting at port : %v", os.Getenv("SERVER_PORT"))
	log.Fatalln(
		http.ListenAndServe(
			fmt.Sprintf(":%v", os.Getenv("SERVER_PORT")),
			packages.AllowCORS(route),
		),
	)
}
