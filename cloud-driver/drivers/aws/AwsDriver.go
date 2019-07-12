// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is a Cloud Driver Example for PoC Test.
//
// by powerkim@etri.re.kr, 2019.06.

package main

import (
	"C"

	acon "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/aws/connect"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	icon "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/connect"
)

type AwsDriver struct{}

func (AwsDriver) GetDriverVersion() string {
	return "TEST AWS DRIVER Version 0.1"
}

func (AwsDriver) GetDriverCapability() idrv.DriverCapabilityInfo {
	var drvCapabilityInfo idrv.DriverCapabilityInfo
	drvCapabilityInfo.VNetworkHandler = false

	return drvCapabilityInfo
}

func (AwsDriver) ConnectCloud(connectionInfo idrv.ConnectionInfo) (icon.CloudConnection, error) {
	// 1. get info of credential and region for Test A Cloud from connectionInfo.
	// 2. create a client object(or service  object) of Test A Cloud with credential info.
	// 3. create CloudConnection Instance of "connect/TDA_CloudConnection".
	// 4. return CloudConnection Interface of TDA_CloudConnection.

	// sample code, do not user like this^^
	var iConn icon.CloudConnection
	iConn = acon.AwsCloudConnection{}
	return iConn, nil // return type: (icon.CloudConnection, error)
}

var TestDriver AwsDriver
