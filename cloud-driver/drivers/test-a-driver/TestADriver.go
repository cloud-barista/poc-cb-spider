package driver

import (
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
        icon "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/connect"
)


type TADCloudDriver struct{}

func (TADCloudDriver) GetDriverVersion() string {
	return "TEST A DRIVER Version 1.0"
}

func (TADCloudDriver) GetDriverCapability() idrv.DriverCapabilityInfo {
	var drvCapabilityInfo idrv.DriverCapabilityInfo
	drvCapabilityInfo.VirtualNetwork = false
	return drvCapabilityInfo
}

func (TADCloudDriver) ConnectCloud(credentialInfo idrv.CredentialInfo) (icon.CloudConnection, error){
	// 1. get info of credential for Test A Cloud.
	// 2. create a client object(or service  object) of Test A Cloud with credential info.
	// 3. create CloudConnection Instance of "connect/TDA_CloudConnection".
	// 4. return CloudConnection Interface of TDA_CloudConnection.

	return nil, nil // return type: (connect.CloudConnection, error)
}
