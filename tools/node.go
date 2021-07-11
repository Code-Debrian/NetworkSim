package tools

type NodeJSON struct {
	Id               int     `json:"id"`
	MaxQSize         int     `json:"max_qsize"`
	MaxInPorts       int     `json:"max_in_ports"`
	MaxOutPorts      int     `json:"max_out_ports"`
	TransmissionRate float32 `json:"transmission_rate"`
	IpAddress        string  `json:"ip"`
	MacAddress       string  `json:"mac"`
}
