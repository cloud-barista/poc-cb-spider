// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is PoC of Cloud Driver Manager.
//
// by powerkim@etri.re.kr, 2019.06.


//package drivermanager
package main

import (
	"flag"
	"fmt"
	osdrv "poc-cb-spider2/cloud-driver/drivers/openstack"
	//osrsc "poc-cb-spider2/cloud-driver/drivers/openstack/resources"
	//idrv "github.com/hyokyungk/poc-cb-spider/cloud-driver/interfaces"
	idrv "poc-cb-spider2/cloud-driver/interfaces"
)


var driverPath *string
func init() {
	driverPath = flag.String("driver", "none", "select driver: -driver=/tmp/TestADriver.so")
	flag.Parse()
}

func main() {

	// Define credential info
	credentialInfo := idrv.CredentialInfo{}
	regionInfo := idrv.RegionInfo{"testRegion", "TestZone"}
	connectionInfo := idrv.ConnectionInfo{credentialInfo, regionInfo}
	fmt.Println(connectionInfo)

	// Test openstack driver
	var cloudDriver idrv.CloudDriver
	cloudDriver = new(osdrv.OpenStackDriver)

	cloudConnection, _ := cloudDriver.ConnectCloud(connectionInfo)

	imageHandler, _ := cloudConnection.CreateImageHandler()
	imgArr, err := imageHandler.ListImage()

	if err != nil {
		fmt.Println(err.Error())
	}
	for _, img := range imgArr {
		fmt.Println(*img)
	}
}

