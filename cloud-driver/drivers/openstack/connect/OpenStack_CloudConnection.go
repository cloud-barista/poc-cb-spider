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
	osrs "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/openstack/resources"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/rackspace/gophercloud"
)

// modified by powerkim, 2019.07.29
type OpenStackCloudConnection struct {
	Client        *gophercloud.ServiceClient
	ImageClient   *gophercloud.ServiceClient
	NetworkClient *gophercloud.ServiceClient
}

func (cloudConn *OpenStackCloudConnection) CreateVNetworkHandler() (irs.VNetworkHandler, error) {
	fmt.Println("OpenStack Cloud Driver: called CreateVNetworkHandler()!")
	vNetworkHandler := osrs.OpenStackVNetworkHandler{cloudConn.NetworkClient}
	return &vNetworkHandler, nil
}

func (cloudConn *OpenStackCloudConnection) CreateImageHandler() (irs.ImageHandler, error) {
	fmt.Println("OpenStack Cloud Driver: called CreateImageHandler()!")
	imageHandler := osrs.OpenStackImageHandler{cloudConn.Client, cloudConn.ImageClient}
	return &imageHandler, nil
}

func (cloudConn OpenStackCloudConnection) CreateSecurityHandler() (irs.SecurityHandler, error) {
	fmt.Println("OpenStack Cloud Driver: called CreateSecurityHandler()!")
	securityHandler := osrs.OpenStackSecurityHandler{cloudConn.Client}
	return &securityHandler, nil
}
func (cloudConn *OpenStackCloudConnection) CreateKeyPairHandler() (irs.KeyPairHandler, error) {
	fmt.Println("OpenStack Cloud Driver: called CreateKeyPairHandler()!")
	keypairHandler := osrs.OpenStackKeyPairHandler{cloudConn.Client}
	return &keypairHandler, nil
}
func (cloudConn *OpenStackCloudConnection) CreateVNicHandler() (irs.VNicHandler, error) {
	fmt.Println("OpenStack Cloud Driver: called CreateVNicHandler()!")
	vNicHandler := osrs.OpenStackVNicworkHandler{cloudConn.NetworkClient}
	return &vNicHandler, nil
}
func (cloudConn OpenStackCloudConnection) CreatePublicIPHandler() (irs.PublicIPHandler, error) {
	fmt.Println("OpenStack Cloud Driver: called CreatePublicIPHandler()!")
	publicIPHandler := osrs.OpenStackPublicIPHandler{cloudConn.Client}
	return &publicIPHandler, nil
}

// modified by powerkim, 2019.07.29
func (cloudConn *OpenStackCloudConnection) CreateVMHandler() (irs.VMHandler, error) {
	//func (OpenStackCloudConnection) CreateVMHandler() (irs.VMHandler, error) {
	//	isConnected, _ := cloudConn.IsConnected()
	//	if(!isConnected) {
	//		return nil, fmt.Errorf("OpenStack Driver is not connected!!")
	//	}

	//	Client, err := config.GetServiceClient()
	//       if err != nil {
	//              panic(err)
	//     }

	fmt.Println("OpenStack Cloud Driver: called CreateVMHandler()!")
	vmHandler := osrs.OpenStackVMHandler{cloudConn.Client}
	return &vmHandler, nil
}

func (OpenStackCloudConnection) IsConnected() (bool, error) {
	return true, nil
}
func (OpenStackCloudConnection) Close() error {
	return nil
}
