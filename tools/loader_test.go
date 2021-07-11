package tools

import (
	"fmt"
	"testing"
)

func TestOpenFile(t *testing.T) {
	net := CreateNetworkFromFile("network_setup.json")
	fmt.Println(*net)
}
