// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is the PoC for dynamic plugin of Cloud Driver Manager.
//
// by powerkim@etri.re.kr, 2019.06.


//package drivermanager
package main


import (
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	//icon "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/connect"

	"flag"
	"plugin"
	"fmt"
	"log"
)


var driverPath *string
func init() {
        driverPath = flag.String("driver", "none", "select driver: -driver=/tmp/TestADriver.so")
        flag.Parse()
}


//func GetCloudConnectionInterface(driverPath string) icon.CloudConnection {


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

/* in CloudDriver.go
	type CredentialInfo struct {
		// @todo TBD
		// key-value pairs
	}

	type RegionInfo struct {
		Region string
		Zone string
	}

	type ConnectionInfo struct {
		CredentialInfo CredentialInfo
		RegionInfo RegionInfo
	}
*/
	credentialInfo := idrv.CredentialInfo{}
	regionInfo := idrv.RegionInfo{"testRegion", "TestZone"}
	connectionInfo := idrv.ConnectionInfo{credentialInfo, regionInfo}
	
	
	cloudConnection, _ := cloudDriver.ConnectCloud(connectionInfo)
	cloudConnection.CreateVNetworkHandler()

}

