// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is interfaces of Cloud Driver.
//
// by powerkim@etri.re.kr, 2019.06.

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

