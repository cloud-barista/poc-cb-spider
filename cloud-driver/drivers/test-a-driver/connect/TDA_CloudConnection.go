package connect

import (
	"fmt"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
)

type TADCloudConnection struct{}

func (TADCloudConnection) CreateVirtualNetworkHandler() (irs.VirtualNetworkHandler, error) {
        fmt.Println("TEST A Cloud Driver: called CreateVirtualNetworkHandler()!")
        return nil, nil
}

