package network

type Node struct {
	Id               int
	MaxQSize         int
	MaxInPorts       int
	MaxOutPorts      int
	TransmissionRate float32
	IpAddress        string
	MacAddress       string
}
