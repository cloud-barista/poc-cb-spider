// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is a Cloud Driver Example for PoC Test.
//
// by powerkim@etri.re.kr, 2019.06.

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

