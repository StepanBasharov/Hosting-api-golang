package vm_stats_api

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"vm_stats_api/pkg/routers"
)

type Server struct {
	port string
}

func LoadEnv(server *Server) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("ERROR: can't load .env")
	}
	server.port = os.Getenv("PORT")
}

func RunServer() {
	server := Server{}
	LoadEnv(&server)
	router := routers.SetupRouters()
	router.Run(server.port)
}
