// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is a Cloud Driver Example for PoC Test.
//
// by hyokyung.kim@innogrid.co.kr, 2019.07.

package azure

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-06-01/compute"
	_ "github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	azcon "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/azure/connect"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	icon "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/connect"
	"time"
)

type AzureDriver struct{}

func (AzureDriver) GetDriverVersion() string {
	return "AZURE DRIVER Version 1.0"
}

func (AzureDriver) GetDriverCapability() idrv.DriverCapabilityInfo {
	var drvCapabilityInfo idrv.DriverCapabilityInfo

	drvCapabilityInfo.ImageHandler = false
	drvCapabilityInfo.VNetworkHandler = false
	drvCapabilityInfo.SecurityHandler = false
	drvCapabilityInfo.KeyPairHandler = false
	drvCapabilityInfo.VNicHandler = false
	drvCapabilityInfo.PublicIPHandler = false
	drvCapabilityInfo.VMHandler = true

	return drvCapabilityInfo
}

func (driver *AzureDriver) ConnectCloud(connectionInfo idrv.ConnectionInfo) (icon.CloudConnection, error) {
	// 1. get info of credential and region for Test A Cloud from connectionInfo.
	// 2. create a client object(or service  object) of Test A Cloud with credential info.
	// 3. create CloudConnection Instance of "connect/TDA_CloudConnection".
	// 4. return CloudConnection Interface of TDA_CloudConnection.

	Ctx, Client, err := getVMClient(connectionInfo.CredentialInfo)
	if err != nil {
		return nil, err
	}

	iConn := azcon.AzureCloudConnection{connectionInfo.RegionInfo, Ctx, Client}
	return &iConn, nil // return type: (icon.CloudConnection, error)
}

func getVMClient(credential idrv.CredentialInfo) (context.Context, *compute.VirtualMachinesClient, error) {
	/*auth.NewClientCredentialsConfig()
	  authorizer, err := auth.NewAuthorizerFromFile(azure.PublicCloud.ResourceManagerEndpoint)
	  if err != nil {
	      return nil, nil, err
	  }*/
	config := auth.NewClientCredentialsConfig(credential.ClientId, credential.ClientSecret, credential.TenantId)
	authorizer, err := config.Authorizer()
	if err != nil {
		return nil, nil, err
	}

	vmClient := compute.NewVirtualMachinesClient(credential.SubscriptionId)
	vmClient.Authorizer = authorizer
	ctx, _ := context.WithTimeout(context.Background(), 600*time.Second)

	return ctx, &vmClient, nil
}

var TestDriver AzureDriver
