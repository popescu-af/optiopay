package main

import (
	"github.com/popescu-af/saas-y/pkg/log"

	"github.com/popescu-af/optiopay/services/main-svc/pkg/client"
	"github.com/popescu-af/optiopay/services/main-svc/pkg/exports"
)

func main() {
	log.Info("main-svc-client started")

	c := client.NewMainSvcClient("bureaucrat")

	employeesToAdd := []*exports.AddInfo{
		{EmployeeName: "Marcellus", ManagerName: "Claire"},
		{EmployeeName: "Mia", ManagerName: "Claire"},
		{EmployeeName: "Vincent", ManagerName: "Marcellus"},
		{EmployeeName: "Jules", ManagerName: "Marcellus"},
	}

	for _, input := range employeesToAdd {
		err := c.AddEmployee(input)
		if err != nil {
			log.FatalCtx("failed to add user", log.Context{"error": err})
		}
	}

	manager, err := c.Manager("Jules", "Vincent")
	if err != nil {
		log.FatalCtx("failed to calculate manager", log.Context{"error": err})
	}

	if manager.Name != "Marcellus" {
		log.FatalCtx("wrong manager found", log.Context{
			"actual":   manager.Name,
			"expected": "Marcellus",
		})
	}

	log.Info("integration tests OK")
}
