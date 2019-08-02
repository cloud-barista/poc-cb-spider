// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is a Cloud Driver Example for PoC Test.
//
// by hyokyung.kim@innogrid.co.kr, 2019.07.

package connect

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-06-01/compute"
	azrs "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/azure/resources"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
)

type AzureCloudConnection struct {
	Region idrv.RegionInfo
	Ctx    context.Context
	Client *compute.VirtualMachinesClient
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

func (cloudConn *AzureCloudConnection) CreateVMHandler() (irs.VMHandler, error) {
	vmHandler := azrs.AzureVMHandler{cloudConn.Region, cloudConn.Ctx, cloudConn.Client}
	return &vmHandler, nil
}

func (AzureCloudConnection) IsConnected() (bool, error) {
	return true, nil
}
func (AzureCloudConnection) Close() error {
	return nil
}
