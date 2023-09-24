package repository

type VmClientsUser struct {
	Username        string
	Password        string
	VirtualMachines []VirtualMachine
	Balance         float64
	PayDay          int64
}

var UsersTest []VmClientsUser

func CreateUser(username string, password string) VmClientsUser {
	NewUser := VmClientsUser{username, password, []VirtualMachine{}, 0, 0}
	UsersTest = append(UsersTest, NewUser)
	return NewUser
}

func (user *VmClientsUser) UpdateUserData(username string, password string, balance float64, payday int64) *VmClientsUser {
	user.Username = username
	user.Password = password
	user.Balance = balance
	user.PayDay = payday

	return user
}

func (user *VmClientsUser) AddUserVMs(machine VirtualMachine) {
	user.VirtualMachines = append(user.VirtualMachines, machine)
}

func test() {

}
