package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vm_stats_api/pkg/repository"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func VirtualMachineList(c *gin.Context) {
	vms := [5]repository.VirtualMachine{
		repository.VirtualMachine{Id: 0, Cpu: "RYZEN 5 5600X", Ram: 32, Rom: 1000, Lan: 100, OSystem: "Ubuntu 18 LTS"},
		repository.VirtualMachine{Id: 1, Cpu: "RYZEN 7 5800X3D", Ram: 64, Rom: 6000, Lan: 1000, OSystem: "Ubuntu 18 LTS"},
		repository.VirtualMachine{Id: 2, Cpu: "INTEL CORE i5 12600k", Ram: 16, Rom: 1000, Lan: 100, OSystem: "Fedora 19"},
		repository.VirtualMachine{Id: 3, Cpu: "RYZEN 3 3200u", Ram: 32, Rom: 1000, Lan: 100, OSystem: "CentOS 19"},
		repository.VirtualMachine{Id: 4, Cpu: "INTEL XEON E5", Ram: 8, Rom: 10000, Lan: 1000, OSystem: "WindowsServer 2016"},
	}
	c.JSON(http.StatusOK, gin.H{"VMs": vms})
}

func GetVirtualMachineStats(c *gin.Context) {
	vms := [5]repository.VirtualMachine{
		repository.VirtualMachine{Id: 0, Cpu: "RYZEN 5 5600X", Ram: 32, Rom: 1000, Lan: 100, OSystem: "Ubuntu 18 LTS"},
		repository.VirtualMachine{Id: 1, Cpu: "RYZEN 7 5800X3D", Ram: 64, Rom: 6000, Lan: 1000, OSystem: "Ubuntu 18 LTS"},
		repository.VirtualMachine{Id: 2, Cpu: "INTEL CORE i5 12600k", Ram: 16, Rom: 1000, Lan: 100, OSystem: "Fedora 19"},
		repository.VirtualMachine{Id: 3, Cpu: "RYZEN 3 3200u", Ram: 32, Rom: 1000, Lan: 100, OSystem: "CentOS 19"},
		repository.VirtualMachine{Id: 4, Cpu: "INTEL XEON E5", Ram: 8, Rom: 10000, Lan: 1000, OSystem: "WindowsServer 2016"},
	}
	paramVmId := c.Param("vm_id")
	vmId, err := strconv.Atoi(paramVmId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Vm id must be int, not string"})
	} else {
		if len(vms)-1 < vmId {
			c.JSON(http.StatusNotFound, gin.H{"error": "Vm " + strconv.Itoa(vmId) + " not found"})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": vms[vmId]})
		}

	}

}
