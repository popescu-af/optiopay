package logic

import (
	"errors"

	"github.com/popescu-af/saas-y/pkg/log"

	"github.com/popescu-af/optiopay/services/main-svc/pkg/exports"
)

// Implementation is the main implementation of the API interface.
type Implementation struct {
}

// NewImpl creates an instance of the main implementation.
func NewImpl() exports.API {
	return &Implementation{}
}

// /add

// AddEmployee implementation.
func (i *Implementation) AddEmployee(input *exports.AddInfo) error {
	log.Info("called add_employee")
	return errors.New("method 'add_employee' not implemented")
}

// /remove

// RemoveEmployee implementation.
func (i *Implementation) RemoveEmployee(input *exports.RemoveInfo) error {
	log.Info("called remove_employee")
	return errors.New("method 'remove_employee' not implemented")
}

// /manager

// Manager implementation.
func (i *Implementation) Manager(firstEmployee string, secondEmployee string) (*exports.ManagerInfo, error) {
	log.Info("called manager")
	return nil, errors.New("method 'manager' not implemented")
}

// /hierarchy

// Hierarchy implementation.
func (i *Implementation) Hierarchy() (*exports.HierarchyInfo, error) {
	log.Info("called hierarchy")
	return nil, errors.New("method 'hierarchy' not implemented")
}
