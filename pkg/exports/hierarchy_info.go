package exports

type Employee struct {
	Managed map[string]*Employee `json:"manages"`
}

// HierarchyInfo - generated API structure
type HierarchyInfo struct {
	Data *Employee `json:"Claire"`
}
