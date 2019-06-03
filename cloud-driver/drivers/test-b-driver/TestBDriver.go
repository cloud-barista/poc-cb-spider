package driver

import (
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
        icon "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/connect"
)


type TBDCloudDriver struct{}

func (TBDCloudDriver) GetDriverVersion() string {
	return "TEST B DRIVER Version 1.0"
}

func (TBDCloudDriver) GetDriverCapability() idrv.DriverCapabilityInfo {
	var drvCapabilityInfo idrv.DriverCapabilityInfo
	drvCapabilityInfo.VirtualNetwork = false
	return drvCapabilityInfo
}

func (TBDCloudDriver) ConnectCloud(credentialInfo idrv.CredentialInfo) (icon.CloudConnection, error){
	// 1. get info of credential for Test B Cloud.
	// 2. create a client object(or service  object) of Test B Cloud with credential info.
	// 3. create CloudConnection Instance of "connect/TDB_CloudConnection".
	// 4. return CloudConnection Interface of TDB_CloudConnection.

	return nil, nil // return type: (connect.CloudConnection, error)
}
