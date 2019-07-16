package main

import (
	"fmt"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	//irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	azdrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/azure"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/config"
)

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

func main() {
	HandleVM()
}