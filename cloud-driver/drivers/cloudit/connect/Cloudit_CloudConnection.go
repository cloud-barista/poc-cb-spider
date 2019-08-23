// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is a Cloud Driver Example for PoC Test.
//
// by hyokyung.kim@innogrid.com, 2019.08.

package connect

import (
	"fmt"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client"
	cirs "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/resources"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
)

type ClouditCloudConnection struct {
	Client client.RestClient
}

func (cloudConn *ClouditCloudConnection) CreateVNetworkHandler() (irs.VNetworkHandler, error) {
	fmt.Println("Cloudit Cloud Driver: called CreateVNetworkHandler()!")
	return nil, nil
}

func (cloudConn *ClouditCloudConnection) CreateImageHandler() (irs.ImageHandler, error) {
	fmt.Println("Cloudit Cloud Driver: called CreateImageHandler()!")
	return nil, nil
}

func (cloudConn ClouditCloudConnection) CreateSecurityHandler() (irs.SecurityHandler, error) {
	fmt.Println("Cloudit Cloud Driver: called CreateSecurityHandler()!")
	return nil, nil
}

func (cloudConn *ClouditCloudConnection) CreateKeyPairHandler() (irs.KeyPairHandler, error) {
	fmt.Println("Cloudit Cloud Driver: called CreateKeyPairHandler()!")
	return nil, nil
}

func (ClouditCloudConnection) CreateVNicHandler() (irs.VNicHandler, error) {
	fmt.Println("Cloudit Cloud Driver: called CreateVNicHandler()!")
	return nil, nil
}

func (cloudConn ClouditCloudConnection) CreatePublicIPHandler() (irs.PublicIPHandler, error) {
	fmt.Println("Cloudit Cloud Driver: called CreatePublicIPHandler()!")
	return nil, nil
}
func (cloudConn *ClouditCloudConnection) CreateVMHandler() (irs.VMHandler, error) {
	fmt.Println("Cloudit Cloud Driver: called CreateVMHandler()!")
	vmHandler := cirs.ClouditVMHandler{&cloudConn.Client}
	return &vmHandler, nil
}

func (ClouditCloudConnection) IsConnected() (bool, error) {
	return true, nil
}
func (ClouditCloudConnection) Close() error {
	return nil
}
