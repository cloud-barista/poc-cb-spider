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
	Client *gophercloud.ServiceClient
}

func (OpenStackCloudConnection) CreateVNetworkHandler() (irs.VNetworkHandler, error) {
	fmt.Println("OpenStack Cloud Driver: called CreateVNetworkHandler()!")
	return nil, nil
}

func (cloudConn *OpenStackCloudConnection) CreateImageHandler() (irs.ImageHandler, error) {
	imageHandler := osrs.OpenStackImageHandler{cloudConn.Client}
	return &imageHandler, nil
}

func (OpenStackCloudConnection) CreateSecurityHandler() (irs.SecurityHandler, error) {
	return nil, nil
}
func (cloudConn *OpenStackCloudConnection) CreateKeyPairHandler() (irs.KeyPairHandler, error) {
	keypairHandler := osrs.OpenStackKeyPairHandler{cloudConn.Client}
	return &keypairHandler, nil
}
func (OpenStackCloudConnection) CreateVNicHandler() (irs.VNicHandler, error) {
	return nil, nil
}
func (OpenStackCloudConnection) CreatePublicIPHandler() (irs.PublicIPHandler, error) {
	return nil, nil
}

/* org.
func (OpenStackCloudConnection) CreateVMHandler() (irs.VMHandler, error) {
	var vmHandler irs.VMHandler
	vmHandler = osrs.OpenStackVMHandler{}
	return vmHandler, nil
}
*/

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

	//var vmHandler irs.VMHandler
	vmHandler := osrs.OpenStackVMHandler{cloudConn.Client}
	return &vmHandler, nil
}

func (OpenStackCloudConnection) IsConnected() (bool, error) {
	return true, nil
}
func (OpenStackCloudConnection) Close() error {
	return nil
}
