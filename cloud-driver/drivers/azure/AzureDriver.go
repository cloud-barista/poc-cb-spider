package azure

import (
	azcon "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/azure/connect"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	icon "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/connect"
)

type AzureDriver struct{}

func (AzureDriver) GetDriverVersion() string {
	return "AZURE DRIVER Version 1.0"
}

func (AzureDriver) GetDriverCapability() idrv.DriverCapabilityInfo {
	var drvCapabilityInfo idrv.DriverCapabilityInfo
	drvCapabilityInfo.VNetworkHandler = false

	return drvCapabilityInfo
}

func (AzureDriver) ConnectCloud(connectionInfo idrv.ConnectionInfo) (icon.CloudConnection, error) {
	// 1. get info of credential and region for Test A Cloud from connectionInfo.
	// 2. create a client object(or service  object) of Test A Cloud with credential info.
	// 3. create CloudConnection Instance of "connect/TDA_CloudConnection".
	// 4. return CloudConnection Interface of TDA_CloudConnection.

	// sample code, do not user like this^^
	var iConn icon.CloudConnection
	iConn = azcon.AzureCloudConnection{}

	return iConn, nil // return type: (icon.CloudConnection, error)
}

var TestDriver AzureDriver
