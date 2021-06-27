package logic

import (
	"github.com/popescu-af/saas-y/pkg/log"

	"github.com/popescu-af/optiopay/services/main-svc/pkg/exports"
)

// Implementation is the main implementation of the API interface.
type Implementation struct {
	storage Storage
}

// NewImpl creates an instance of the main implementation.
func NewImpl(storage Storage) exports.API {
	return &Implementation{
		storage: storage,
	}
}

// /add

// AddEmployee implementation.
func (i *Implementation) AddEmployee(input *exports.AddInfo) error {
	log.Info("called add_employee")
	return i.storage.AddEmployee(input.EmployeeName, input.ManagerName)
}

// /remove

// RemoveEmployee implementation.
func (i *Implementation) RemoveEmployee(input *exports.RemoveInfo) error {
	log.Info("called remove_employee")
	return i.storage.RemoveEmployee(input.EmployeeName, input.ManagerTakingOver)
}

// /manager

// Manager implementation.
func (i *Implementation) Manager(firstEmployee string, secondEmployee string) (*exports.ManagerInfo, error) {
	log.Info("called manager")
	managerName, err := i.storage.Manager(firstEmployee, secondEmployee)
	if err != nil {
		return nil, err
	}
	return &exports.ManagerInfo{Name: managerName}, nil
}

// /hierarchy

// Hierarchy implementation.
func (i *Implementation) Hierarchy() (*exports.HierarchyInfo, error) {
	log.Info("called hierarchy")
	return &exports.HierarchyInfo{
		Data: i.storage.Hierarchy(),
	}, nil
}
