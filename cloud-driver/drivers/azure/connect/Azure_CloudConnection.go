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
	//irs "github.com/hyokyungk/poc-cb-spider/cloud-driver/interfaces/resources"
	irs "poc-cb-spider2/cloud-driver/interfaces/resources"
)

type AzureCloudConnection struct {}

func (AzureCloudConnection) TestConnection() error {

	// Test Azure Connection

	return nil
}

func (AzureCloudConnection) CreateVNetworkHandler() (irs.VNetworkHandler, error) {
	fmt.Println("Azure Cloud Driver: called CreateVNetworkHandler()!")
	return nil, nil
}
func (AzureCloudConnection) CreateImageHandler() (irs.ImageHandler, error) {
	return nil, nil
}
func (AzureCloudConnection) CreateSecurityHandler() (irs.SecurityHandler, error) {
	return nil, nil
}
func (AzureCloudConnection) CreateKeyPairHandler() (irs.KeyPairHandler, error) {
	return nil, nil
}
func (AzureCloudConnection) CreateVNicHandler() (irs.VNicHandler, error) {
	return nil, nil
}
func (AzureCloudConnection) CreatePublicIPHandler() (irs.PublicIPHandler, error) {
	return nil, nil
}

func (AzureCloudConnection) CreateVMHandler() (irs.VMHandler, error) {
	return nil, nil
}

func (AzureCloudConnection) IsConnected() (bool, error) {
	return true, nil
}
func (AzureCloudConnection) Close() error {
	return nil
}
