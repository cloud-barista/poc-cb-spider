package driver

import (
        icon "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/connect"
)

type DriverCapabilityInfo struct {
	VirtualNetwork bool // support: true, do not support: false
	// @todo TBD
}

type CredentialInfo struct {
	// @todo TBD
	// key-value pairs
}

type CloudDriver interface {
	GetDriverVersion() string
	GetDriverCapability() DriverCapabilityInfo

	ConnectCloud(credentialInfo CredentialInfo) (icon.CloudConnection, error)
}

