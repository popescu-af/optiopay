package concretes

import (
	"testing"

	"github.com/popescu-af/optiopay/services/main-svc/internal/logic"
	"github.com/popescu-af/optiopay/services/main-svc/pkg/exports"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	s := NewInMemoryStorage()

	tests := []struct {
		employeeName string
		managerName  string
		expectError  bool
	}{
		{"Claire", "Claire", true},     // employee already exists
		{"Alex", "Claire", false},      // OK
		{"Alex", "Alex", true},         // employee already exists
		{"Richard", "Christian", true}, // no such manager
		{"Alex", "Claire", true},       // employee already exists
		{"Stephanie", "Alex", false},   // OK
		{"Richard", "Claire", false},   // OK
	}

	for _, tc := range tests {
		err := s.AddEmployee(tc.employeeName, tc.managerName)
		if tc.expectError {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}

func createStorageAndAddSomeEmployees(t *testing.T) logic.Storage {
	s := NewInMemoryStorage()
	require.NoError(t, s.AddEmployee("Alex", "Claire"))
	require.NoError(t, s.AddEmployee("Richard", "Claire"))
	require.NoError(t, s.AddEmployee("Stephanie", "Alex"))
	require.NoError(t, s.AddEmployee("Tom", "Alex"))
	require.NoError(t, s.AddEmployee("Hank", "Richard"))
	require.NoError(t, s.AddEmployee("Violetta", "Richard"))
	require.NoError(t, s.AddEmployee("Omar", "Tom"))
	require.NoError(t, s.AddEmployee("Saul", "Stephanie"))
	return s
}

func hierarchiesEqual(h0, h1 *exports.Employee) bool {
	for k, v0 := range h0.Managed {
		v1, ok := h1.Managed[k]
		if !ok {
			return false
		}
		if !hierarchiesEqual(v0, v1) {
			return false
		}
	}
	return true
}

func TestRemove(t *testing.T) {
	s := createStorageAndAddSomeEmployees(t)

	tests := []struct {
		employeeName          string
		managerTakingOverName string
		expectError           bool
	}{
		{"Alex", "Alex", true},        // manager taking over is being removed
		{"Jimmy", "Alex", true},       // no such employee
		{"Alex", "Jimmy", true},       // no such employee to take over
		{"Alex", "Stephanie", true},   // employee to take over is managed by employee to be removed
		{"Claire", "Stephanie", true}, // employee to take over is managed by employee to be removed (CEO)
		{"Hank", "Richard", false},    // OK
		{"Alex", "Richard", false},    // OK
	}

	for _, tc := range tests {
		err := s.RemoveEmployee(tc.employeeName, tc.managerTakingOverName)
		if tc.expectError {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}

	expectedHierarchy := &exports.Employee{
		Managed: map[string]*exports.Employee{
			"Richard": {
				Managed: map[string]*exports.Employee{
					"Stephanie": {
						Managed: map[string]*exports.Employee{
							"Saul": {Managed: map[string]*exports.Employee{}},
						},
					},
					"Tom": {
						Managed: map[string]*exports.Employee{
							"Omar": {Managed: map[string]*exports.Employee{}},
						},
					},
					"Violetta": {Managed: map[string]*exports.Employee{}},
				},
			},
		},
	}

	require.True(t, hierarchiesEqual(s.Hierarchy(), expectedHierarchy))
}

func TestManager(t *testing.T) {
	s := createStorageAndAddSomeEmployees(t)

	tests := []struct {
		firstEmployeeName  string
		secondEmployeeName string
		expectError        bool
		expectedManager    string
	}{
		{"John", "Alex", true, ""},                // first employee doesn't exist
		{"Alex", "John", true, ""},                // second employee doesn't exist
		{"Stephanie", "Richard", false, "Claire"}, // OK
		{"Stephanie", "Claire", false, "Claire"},  // OK (CEO)
		{"Saul", "Omar", false, "Alex"},           // OK (manager lower than CEO)
	}

	for _, tc := range tests {
		actualManager, err := s.Manager(tc.firstEmployeeName, tc.secondEmployeeName)
		if tc.expectError {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
			require.Equal(t, actualManager, tc.expectedManager)
		}
	}
}
