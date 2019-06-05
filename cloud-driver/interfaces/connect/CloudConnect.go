// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is Connection interfaces of Cloud Driver.
//
// by powerkim@etri.re.kr, 2019.06.


package connect

import (
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
)

type CloudConnection interface {
	CreateVirtualNetworkHandler() (irs.VirtualNetworkHandler, error)
	CreateImageHandler() (irs.ImageHandler, error)	
	CreateSecurityHandler() (irs.SecurityHandler, error)
	CreateConnectionKeyHandler() (irs.ConnectionKeyHandler, error)

	IsConnected() (bool, error)
	Close()	error
}

