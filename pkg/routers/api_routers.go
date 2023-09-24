package routers

import (
	"github.com/gin-gonic/gin"
	"vm_stats_api/pkg/handlers"
)

func IncludeAPIRouters(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/health", handlers.HealthCheck)
	}
	apiVms := router.Group("/api/vms")
	{
		apiVms.GET("/vm_list", handlers.VirtualMachineList)
		apiVms.GET("/vm/:vm_id", handlers.GetVirtualMachineStats)
	}
	apiUsers := router.Group("/api/users")
	{
		apiUsers.GET("/")
	}

}
