package connect

import (
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"fmt"
)

type TBDCloudConnection struct{}

func (TBDCloudConnection) CreateVirtualNetworkHandler() (irs.VirtualNetworkHandler, error) {
        fmt.Println("TEST B Cloud Driver: called CreateVirtualNetworkHandler()!")
        return nil, nil
}

