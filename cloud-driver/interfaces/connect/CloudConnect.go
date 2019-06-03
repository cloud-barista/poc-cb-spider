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

