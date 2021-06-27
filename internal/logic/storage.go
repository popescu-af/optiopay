package logic

import "github.com/popescu-af/optiopay/services/main-svc/pkg/exports"

type Employee struct {
	Name    string
	Managed []*Employee
}

type Storage interface {
	AddEmployee(employeeName, managerName string) error
	RemoveEmployee(employeeName, managerTakingOver string) error
	Manager(firstEmployee string, secondEmployee string) (string, error)
	Hierarchy() (*exports.HierarchyInfo, error)
}
