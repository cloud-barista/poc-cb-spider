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
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	irs "github.com/hyokyungk/poc-cb-spider/cloud-driver/interfaces/resources"
)

type OpenStackCloudConnection struct {}

func (OpenStackCloudConnection) TestConnection() error {

	// Test OpenStack Connection

	scop := gophercloud.AuthScope{
		ProjectID: "7b50974d16304341975728fb0571851b",
	}
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: "http://{url}:5000/v3",
		Username:         "{UserName}",
		Password:         "{Password}",
		Scope:            &scop,
	}

	provider, _ := openstack.AuthenticatedClient(opts)
	fmt.Println(provider.GetAuthResult())

	return nil
}

func (OpenStackCloudConnection) CreateVNetworkHandler() (irs.VNetworkHandler, error) {
	fmt.Println("OpenStack Cloud Driver: called CreateVNetworkHandler()!")
	return nil, nil
}
func (OpenStackCloudConnection) CreateImageHandler() (irs.ImageHandler, error) {
	return nil, nil
}
func (OpenStackCloudConnection) CreateSecurityHandler() (irs.SecurityHandler, error) {
	return nil, nil
}
func (OpenStackCloudConnection) CreateKeyPairHandler() (irs.KeyPairHandler, error) {
	return nil, nil
}
func (OpenStackCloudConnection) CreateVNicHandler() (irs.VNicHandler, error) {
	return nil, nil
}
func (OpenStackCloudConnection) CreatePublicIPHandler() (irs.PublicIPHandler, error) {
	return nil, nil
}

func (OpenStackCloudConnection) CreateVMHandler() (irs.VMHandler, error) {
	return nil, nil
}

func (OpenStackCloudConnection) IsConnected() (bool, error) {
	return true, nil
}
func (OpenStackCloudConnection) Close() error {
	return nil
}
