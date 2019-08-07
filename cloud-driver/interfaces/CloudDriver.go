// Cloud Driver Interface of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is interfaces of Cloud Driver.
//
// by powerkim@etri.re.kr, 2019.06.

package interfaces

import (
	icon "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/connect"
)

type DriverCapabilityInfo struct {
	ImageHandler    bool // support: true, do not support: false
	VNetworkHandler bool // support: true, do not support: false
	SecurityHandler bool // support: true, do not support: false
	KeyPairHandler  bool // support: true, do not support: false
	VNicHandler     bool // support: true, do not support: false
	PublicIPHandler bool // support: true, do not support: false
	VMHandler       bool // support: true, do not support: false
}

type CredentialInfo struct {
	// @todo TBD
	// key-value pairs
	SubscriptionId string // Azure Credential
}

type RegionInfo struct {
	Region string
	Zone   string
}

type ConnectionInfo struct {
	CredentialInfo CredentialInfo
	RegionInfo     RegionInfo
}

type CloudDriver interface {
	GetDriverVersion() string
	GetDriverCapability() DriverCapabilityInfo

	ConnectCloud(connectionInfo ConnectionInfo) (icon.CloudConnection, error)
	ConnectNetworkCloud(connectionInfo ConnectionInfo) (icon.CloudConnection, error)
}
