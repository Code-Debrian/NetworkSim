package network

type Link struct {
	PropogationRate float32
	Length          int
	InTransit       bool
	Exists          bool
}
