package registry

// Service ...
type Service struct {
	Name  string  `json:"name"`
	Nodes []*Node `json:"nodes"`
}

// Node ...
type Node struct {
	ID     string `json:"id"`
	IP     string `json:"ip"`
	Port   int    `json:"port"`
	Weight int    `json:"weight"`
}
