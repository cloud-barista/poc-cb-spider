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
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-04-01/network"
	azrs "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/azure/resources"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
)

type AzureCloudConnection struct {
	Region              idrv.RegionInfo
	Ctx                 context.Context
	VMClient            *compute.VirtualMachinesClient
	ImageClient         *compute.ImagesClient
	PublicIPClient      *network.PublicIPAddressesClient
	SecurityGroupClient *network.SecurityGroupsClient
	VNetClient          *network.VirtualNetworksClient
	VNicClient          *network.InterfacesClient
	SubnetClient        *network.SubnetsClient
}

func (cloudConn *AzureCloudConnection) CreateVNetworkHandler() (irs.VNetworkHandler, error) {
	fmt.Println("Azure Cloud Driver: called CreateVNetworkHandler()!")
	vNetHandler := azrs.AzureVNetworkHandler{cloudConn.Region, cloudConn.Ctx, cloudConn.VNetClient}
	return &vNetHandler, nil
}

func (cloudConn *AzureCloudConnection) CreateImageHandler() (irs.ImageHandler, error) {
	fmt.Println("Azure Cloud Driver: called CreateImageHandler()!")
	imageHandler := azrs.AzureImageHandler{cloudConn.Region, cloudConn.Ctx, cloudConn.ImageClient}
	return &imageHandler, nil
}

func (cloudConn *AzureCloudConnection) CreateSecurityHandler() (irs.SecurityHandler, error) {
	fmt.Println("Azure Cloud Driver: called CreateSecurityHandler()!")
	sgHandler := azrs.AzureSecurityHandler{cloudConn.Region, cloudConn.Ctx, cloudConn.SecurityGroupClient}
	return &sgHandler, nil
}
func (AzureCloudConnection) CreateKeyPairHandler() (irs.KeyPairHandler, error) {
	return nil, nil
}
func (cloudConn *AzureCloudConnection) CreateVNicHandler() (irs.VNicHandler, error) {
	fmt.Println("Azure Cloud Driver: called CreateVNicHandler()!")
	vNicHandler := azrs.AzureVNicHandler{cloudConn.Region, cloudConn.Ctx, cloudConn.VNicClient, cloudConn.SubnetClient}
	return &vNicHandler, nil
}
func (cloudConn *AzureCloudConnection) CreatePublicIPHandler() (irs.PublicIPHandler, error) {
	fmt.Println("Azure Cloud Driver: called CreatePublicIPHandler()!")
	publicIPHandler := azrs.AzurePublicIPHandler{cloudConn.Region, cloudConn.Ctx, cloudConn.PublicIPClient}
	return &publicIPHandler, nil
}

func (cloudConn *AzureCloudConnection) CreateVMHandler() (irs.VMHandler, error) {
	fmt.Println("Azure Cloud Driver: called CreateVMHandler()!")
	vmHandler := azrs.AzureVMHandler{cloudConn.Region, cloudConn.Ctx, cloudConn.VMClient}
	return &vmHandler, nil
}

func (AzureCloudConnection) IsConnected() (bool, error) {
	return true, nil
}
func (AzureCloudConnection) Close() error {
	return nil
}
