// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is PoC of Cloud Driver Info Manager.
//
// by powerkim@etri.re.kr, 2019.07.


//package drivermanager
package main

import (
	cbs "github.com/cloud-barista/poc-cb-store"

	"fmt"
)

type CloudDriverInfo struct {
	providerName string
	driverName string
	driverPath string	
}

func RegisterCloudDriver(providerName string, driverName string, driverPath string) CloudDriverInfo {
	cldDrvInfo := CloudDriverInfo{providerName, driverName, driverPath}
	
	// @todo save into storage
	writer := cbs.GetWriter()	
	err := writer.PutKV("key1", "value1")
	if err != nil {
		panic(err)
	}

	return cldDrvInfo
}

func ListCloudDriver() []*CloudDriverInfo {
	var cldDrvInfoList []*CloudDriverInfo

	// @todo get list from storage

	return cldDrvInfoList
}

func GetCloudDriver(driverName string) CloudDriverInfo {
	var cldDrvInfo CloudDriverInfo

	// @todo

	return cldDrvInfo
}

func UnRegisterCloudDriver(driverName string) bool {
	var result bool

	// @todo

	return result
}

func main() {
	cldDrvInfo := RegisterCloudDriver("csp1", "csp1_driver", "csp1_driver_path")
	fmt.Printf(">>> %#v\n", cldDrvInfo);
}


/* only to refer by powerkim
func main() {

	var plug *plugin.Plugin
	var err error
        if *driverPath == "none" {
		fmt.Println("Usage: CloudDriverManager -driver=/tmp/TestADriver.so")
		return
        }

	//fmt.Println("######### driver path:" + *driverPath)
	plug, err = plugin.Open(*driverPath)

	// fmt.Printf("plug: %#v\n\n", plug)	
	if err != nil {
		log.Fatalf("plugin.Open: %v\n", err)	
		return
	}


	testDriver, err := plug.Lookup("TestDriver")	
	if err != nil {
		log.Fatalf("plug.Lookup: %v\n", err)	
		return
	}

	cloudDriver, ok := testDriver.(idrv.CloudDriver)
	if !ok {
		log.Fatalf("Not CloudDriver interface!!")
		return
	}

	fmt.Printf("%s: %s\n", *driverPath, cloudDriver.GetDriverVersion())

///* in CloudDriver.go
//	type CredentialInfo struct {
		// @todo TBD
		// key-value pairs
//	}

//	type RegionInfo struct {
//		Region string
//		Zone string
//	}

//	type ConnectionInfo struct {
//		CredentialInfo CredentialInfo
//		RegionInfo RegionInfo
//	}

//	credentialInfo := idrv.CredentialInfo{}
//	regionInfo := idrv.RegionInfo{"testRegion", "TestZone"}
//	connectionInfo := idrv.ConnectionInfo{credentialInfo, regionInfo}
	
	
//	cloudConnection, _ := cloudDriver.ConnectCloud(connectionInfo)
//	cloudConnection.CreateVNetworkHandler()

}
*/
