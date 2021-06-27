package exports

// API defines the operations supported by the main-svc service.
type API interface {
	// /add
	AddEmployee(*AddInfo) error

	// /remove
	RemoveEmployee(*RemoveInfo) error

	// /manager
	Manager(string, string) (*ManagerInfo, error)

	// /hierarchy
	Hierarchy() (*HierarchyInfo, error)
}

// APIClient defines the operations supported by the main-svc service client.
type APIClient interface {
	// /add
	AddEmployee(*AddInfo) error

	// /remove
	RemoveEmployee(*RemoveInfo) error

	// /manager
	Manager(string, string) (*ManagerInfo, error)

	// /hierarchy
	Hierarchy() (*HierarchyInfo, error)
}
