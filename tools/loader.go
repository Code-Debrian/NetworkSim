package tools

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/figueroacode/networksim/network"
)

func CreateNetworkFromFile(filename string) *network.Network {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("Failed to open file: %v", filename)
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)
	var jsonNetwork NetworkJSON
	json.Unmarshal(byteValue, &jsonNetwork)

	var net network.Network
	// Copy the node information
	for _, node := range jsonNetwork.Nodes {
		net.Nodes = append(net.Nodes, network.Node(node))
	}
	// Create adjacency list
	net.InitLinks(len(net.Nodes))
	for _, link := range jsonNetwork.Links {
		net.Links[link.Source][link.Target] = network.Link{
			PropogationRate: link.PropogationRate,
			Length:          link.Length,
			InTransit:       false,
			Exists:          true,
		}
	}
	return &net
}
