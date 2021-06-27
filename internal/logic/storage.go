package logic

import "github.com/popescu-af/optiopay/services/main-svc/pkg/exports"

type Employee struct {
	Name    string
	Managed map[string]*Employee
}

type Storage interface {
	AddEmployee(employeeName, managerName string) error
	RemoveEmployee(employeeName, managerTakingOverName string) error
	Manager(firstEmployeeName, secondEmployeeName string) (string, error)
	Hierarchy() *exports.Employee
}
