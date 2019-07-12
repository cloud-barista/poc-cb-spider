package main

import (
	"fmt"
	osdrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/openstack"
	osrs "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/openstack/resources"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
)

func TestConnection() {
	client, err := osrs.GetServiceClient()
	if err != nil {
		panic(err)
	}
	fmt.Println(client)
}

func TestVMHandler() {
	var cloudDriver idrv.CloudDriver
	cloudDriver = new(osdrv.OpenStackDriver)

	connectionInfo := idrv.ConnectionInfo{}
	cloudConnection, _ := cloudDriver.ConnectCloud(connectionInfo)
	vmHandler, err := cloudConnection.CreateVMHandler()
	if err != nil {
		panic(err)
	}

	config := osrs.ReadConfigFile()

	// Get VM List
	vmList := vmHandler.ListVM()
	for i, vm := range vmList {
		fmt.Println("[",i,"] ",*vm)
	}

	// Get VM Info
	vmInfo := vmHandler.GetVM(config.Openstack.ServerId)
	fmt.Println(vmInfo)
}

func CreateVM() {
	var cloudDriver idrv.CloudDriver
	cloudDriver = new(osdrv.OpenStackDriver)

	connectionInfo := idrv.ConnectionInfo{}
	cloudConnection, _ := cloudDriver.ConnectCloud(connectionInfo)
	vmHandler, err := cloudConnection.CreateVMHandler()
	if err != nil {
		panic(err)
	}

	config := osrs.ReadConfigFile()

	// Create VM Server
	vmReqInfo := irs.VMReqInfo{
		Name: config.Openstack.VMName,
		ImageInfo: irs.ImageInfo{
			Id: config.Openstack.ImageId,
		},
		SpecID: config.Openstack.FlavorId,
		VNetworkInfo: irs.VNetworkInfo{
			Id: config.Openstack.NetworkId,
		},
		SecurityInfo: irs.SecurityInfo{
			Name: config.Openstack.SecurityGroups,
		},
		KeyPairInfo: irs.KeyPairInfo{
			Name: config.Openstack.KeypairName,
		},
	}
	createdVM, err := vmHandler.StartVM(vmReqInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("VM_ID=", createdVM.Id)
}

func HandleVM() {

	fmt.Println("VM LifeCycle Management")
	fmt.Println("1. Suspend VM")
	fmt.Println("2. Resume VM")
	fmt.Println("3. Reboot VM")
	fmt.Println("4. Terminate VM")

	config := osrs.ReadConfigFile()

	var cloudDriver idrv.CloudDriver
	cloudDriver = new(osdrv.OpenStackDriver)

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

		if inputCnt == 1 {
			switch commandNum {
			case 1:
				vmHandler.SuspendVM(config.Openstack.ServerId)
			case 2:
				vmHandler.ResumeVM(config.Openstack.ServerId)
			case 3:
				vmHandler.RebootVM(config.Openstack.ServerId)
			case 4:
				vmHandler.TerminateVM(config.Openstack.ServerId)
			}
		}
	}
}

func main() {
	//TestConnection()
	//TestVMHandler()
	//CreateVM()
	HandleVM()
}
