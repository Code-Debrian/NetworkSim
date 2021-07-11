package tools

type NetworkJSON struct {
	Links []LinkJSON `json:"links"`
	Nodes []NodeJSON `json:"nodes"`
}
