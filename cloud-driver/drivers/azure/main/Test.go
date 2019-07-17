package main

import (
	"fmt"
	azdrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/azure"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/config"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
)

// Test VM Handler Functions (Get VM Info, VM Status)
func TestVMHandler() {
	var cloudDriver idrv.CloudDriver
	cloudDriver = new(azdrv.AzureDriver)

	connectionInfo := idrv.ConnectionInfo{}
	cloudConnection, _ := cloudDriver.ConnectCloud(connectionInfo)
	vmHandler, err := cloudConnection.CreateVMHandler()
	if err != nil {
		panic(err)
	}

	config := config.ReadConfigFile()

	// Get VM List
	vmList := vmHandler.ListVM()
	for i, vm := range vmList {
		fmt.Println("[", i, "] ", *vm)
	}

	// Get VM Info
	vmInfo := vmHandler.GetVM(config.Azure.ServerId)
	fmt.Println(vmInfo)

	// Get VM Status List
	vmStatusList := vmHandler.ListVMStatus()
	for i, vmStatus := range vmStatusList {
		fmt.Println("[",i,"] ",*vmStatus)
	}

	// Get VM Status
	vmStatus := vmHandler.GetVMStatus(config.Azure.VMName)
	fmt.Println(vmStatus)
}

// Test VM Lifecycle Management (Suspend/Resume/Reboot/Terminate)
func HandleVM() {

	fmt.Println("VM LifeCycle Management")
	fmt.Println("1. Suspend VM")
	fmt.Println("2. Resume VM")
	fmt.Println("3. Reboot VM")
	fmt.Println("4. Terminate VM")

	config := config.ReadConfigFile()

	var cloudDriver idrv.CloudDriver
	cloudDriver = new(azdrv.AzureDriver)

	connectionInfo := idrv.ConnectionInfo{}
	cloudConnection, _ := cloudDriver.ConnectCloud(connectionInfo)
	vmHandler, err := cloudConnection.CreateVMHandler()
	if err != nil {
		panic(err)
	}

	for {
		var commandNum int
		inputCnt, err := fmt.Scan(&commandNum)
		if err != nil {
			panic(err)
		}

		vmId := config.Azure.VMName

		if inputCnt == 1 {
			switch commandNum {
			case 1:
				vmHandler.SuspendVM(vmId)
			case 2:
				vmHandler.ResumeVM(vmId)
			case 3:
				vmHandler.RebootVM(vmId)
			case 4:
				vmHandler.TerminateVM(vmId)
			}
		}
	}
}

func CreateVM() {
	var cloudDriver idrv.CloudDriver
	cloudDriver = new(azdrv.AzureDriver)

	connectionInfo := idrv.ConnectionInfo{}
	cloudConnection, _ := cloudDriver.ConnectCloud(connectionInfo)
	vmHandler, err := cloudConnection.CreateVMHandler()
	if err != nil {
		panic(err)
	}

	config := config.ReadConfigFile()

	imageId := config.Azure.Image.Publisher + ":" + config.Azure.Image.Offer + ":" + config.Azure.Image.Sku + ":" + config.Azure.Image.Version
	//specId := config.Azure.VMSize
	//fmt.Println(imageId)

	vmReqInfo := irs.VMReqInfo{
		Name: config.Openstack.VMName,
		ImageInfo: irs.ImageInfo{
			Id: imageId,
		},
		//SpecID: config.Azure.VMSize,
		VNetworkInfo: irs.VNetworkInfo{
			Id: config.Azure.Network.ID,
		},
	}

	vmHandler.StartVM(vmReqInfo)
}

func main() {
	//TestVMHandler()
	//HandleVM()
	//CreateVM()
}
