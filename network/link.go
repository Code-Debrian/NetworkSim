package network

type Link struct {
	PropogationRate float32
	Target          int
	Length          int
	InTransit       bool
	Exists          bool
}
