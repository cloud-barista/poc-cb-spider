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
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"log"
)

type OpenStackCloudConnection struct {
}

func (OpenStackCloudConnection) connectionTest() {

	scop := gophercloud.AuthScope{
		ProjectID: "7b50974d16304341975728fb0571851b",
	}
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: "http://{url}:5000/v3",
		Username:         "{UserName}",
		Password:         "{Password}",
		DomainName:       "default",
		Scope:            &scop,
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(provider.GetAuthResult())
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
