package models

type VmClientsUser struct {
	Username        string
	Password        string
	VirtualMachines []VirtualMachine
	Balance         float64
	PayDay          int64
}
