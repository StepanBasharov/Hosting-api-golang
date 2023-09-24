package vm_stats_api

import (
	"github.com/gin-gonic/gin"
	"vm_stats_api/pkg/routers"
)

type Server struct {
	port string
}

func RunServer() {
	server := Server{port: ":5050"}
	route := gin.Default()
	routers.IncludeAPIRouters(route)
	route.Run(server.port)
}
