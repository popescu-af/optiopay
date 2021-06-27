package concretes

import (
	"fmt"
	"sync"

	"github.com/popescu-af/optiopay/services/main-svc/internal/logic"
	"github.com/popescu-af/optiopay/services/main-svc/pkg/exports"
)

type InMemoryStorage struct {
	mtx    sync.RWMutex
	ceo    *logic.Employee
	lookup map[string]*logic.Employee
}

func NewInMemoryStorage() logic.Storage {
	ceo := &logic.Employee{
		Name:    "Claire",
		Managed: make(map[string]*logic.Employee),
	}
	result := &InMemoryStorage{
		ceo:    ceo,
		lookup: make(map[string]*logic.Employee),
	}
	result.lookup["Claire"] = ceo
	return result
}

func (i *InMemoryStorage) AddEmployee(employeeName, managerName string) error {
	i.mtx.Lock()
	defer i.mtx.Unlock()

	if _, ok := i.lookup[employeeName]; ok {
		return logic.NewAlreadyFoundError(fmt.Sprintf("employee named '%s' already exists", employeeName))
	}

	manager, ok := i.lookup[managerName]
	if !ok {
		return logic.NewNotFoundError(fmt.Sprintf("employee named '%s' does not exist", managerName))
	}

	if _, ok := manager.Managed[employeeName]; ok {
		return logic.NewAlreadyFoundError(fmt.Sprintf("employee named '%s' is already managed by '%s'", employeeName, managerName))
	}

	i.lookup[employeeName] = &logic.Employee{
		Name:    employeeName,
		Managed: make(map[string]*logic.Employee),
	}

	manager.Managed[employeeName] = i.lookup[employeeName]
	return nil
}

func (i *InMemoryStorage) RemoveEmployee(employeeName, managerTakingOverName string) error {
	i.mtx.Lock()
	defer i.mtx.Unlock()

	if employeeName == managerTakingOverName {
		return logic.NewArgumentError(fmt.Sprintf("employee and manager taking over have the same name '%s'", employeeName))
	}

	employee, ok := i.lookup[employeeName]
	if !ok {
		return logic.NewNotFoundError(fmt.Sprintf("employee named '%s' does not exist", employeeName))
	}

	managerTakingOver, ok := i.lookup[managerTakingOverName]
	if !ok {
		return logic.NewNotFoundError(fmt.Sprintf("employee named '%s' to take over employees does not exist", managerTakingOverName))
	}

	// TODO: check that manager taking over is not managed by the employee being removed

	managersOfEmployee, err := i.path(employeeName)
	if err != nil {
		return err
	}

	countManagers := len(managersOfEmployee)
	if countManagers == 0 {
		return logic.NewArgumentError("cannot remove ceo")
	}

	for k, v := range employee.Managed {
		managerTakingOver.Managed[k] = v
	}

	delete(i.lookup, employeeName)
	delete(managersOfEmployee[countManagers-1].Managed, employeeName)
	return nil
}

func (i *InMemoryStorage) Manager(firstEmployeeName, secondEmployeeName string) (string, error) {
	i.mtx.RLock()
	defer i.mtx.RUnlock()

	managersFirstEmployee, err := i.path(firstEmployeeName)
	if err != nil {
		return "", err
	}

	managersSecondEmployee, err := i.path(secondEmployeeName)
	if err != nil {
		return "", err
	}

	l := len(managersFirstEmployee)
	if len(managersSecondEmployee) < l {
		l = len(managersSecondEmployee)
	}

	if l == 0 {
		return i.ceo.Name, nil
	}

	for i := 0; i < l; i++ {
		if managersFirstEmployee[i] != managersSecondEmployee[i] {
			return managersFirstEmployee[i-1].Name, nil
		}
	}

	return managersFirstEmployee[l-1].Name, nil
}

func (i *InMemoryStorage) Hierarchy() *exports.Employee {
	i.mtx.RLock()
	defer i.mtx.RUnlock()

	return i.copy(i.ceo)
}

func (i *InMemoryStorage) copy(employee *logic.Employee) *exports.Employee {
	result := &exports.Employee{
		Managed: make(map[string]*exports.Employee),
	}
	for k, v := range employee.Managed {
		result.Managed[k] = i.copy(v)
	}
	return result
}

func (i *InMemoryStorage) path(employeeName string) ([]*logic.Employee, error) {

	return nil, nil
}
