// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is a Cloud Driver Example for PoC Test.
//
// by hyokyung.kim@innogrid.co.kr, 2019.07.

package main

import (
	"fmt"
	azdrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/azure"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/config"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

// Test VM Handler Functions (Get VM Info, VM Status)
func getVMInfo() {
	vmHandler, err := setVMHandler()
	if err != nil {
		panic(err)
	}
	config := readConfigFile()

	// Get VM List
	vmList := vmHandler.ListVM()
	for i, vm := range vmList {
		fmt.Println("[", i, "] ")
		spew.Dump(vm)
	}

	vmId := config.Azure.GroupName + ":" + config.Azure.VMName

	// Get VM Info
	vmInfo := vmHandler.GetVM(vmId)
	spew.Dump(vmInfo)

	// Get VM Status List
	vmStatusList := vmHandler.ListVMStatus()
	for i, vmStatus := range vmStatusList {
		fmt.Println("[", i, "] ", *vmStatus)
	}

	// Get VM Status
	vmStatus := vmHandler.GetVMStatus(vmId)
	fmt.Println(vmStatus)
}

// Test VM Lifecycle Management (Suspend/Resume/Reboot/Terminate)
func handleVM() {
	vmHandler, err := setVMHandler()
	if err != nil {
		panic(err)
	}
	config := readConfigFile()

	fmt.Println("VM LifeCycle Management")
	fmt.Println("1. Suspend VM")
	fmt.Println("2. Resume VM")
	fmt.Println("3. Reboot VM")
	fmt.Println("4. Terminate VM")

	for {
		var commandNum int
		inputCnt, err := fmt.Scan(&commandNum)
		if err != nil {
			panic(err)
		}

		vmId := config.Azure.GroupName + ":" + config.Azure.VMName

		if inputCnt == 1 {
			switch commandNum {
			case 1:
				fmt.Println("Start Suspend VM ...")
				vmHandler.SuspendVM(vmId)
			case 2:
				fmt.Println("Start Resume  VM ...")
				vmHandler.ResumeVM(vmId)
			case 3:
				fmt.Println("Start Reboot  VM ...")
				vmHandler.RebootVM(vmId)
			case 4:
				fmt.Println("Start Terminate  VM ...")
				vmHandler.TerminateVM(vmId)
			}
		}
	}
}

// Test VM Deployment
func createVM() {
	vmHandler, err := setVMHandler()
	if err != nil {
		panic(err)
	}
	config := config.ReadConfigFile()

	vmName := config.Azure.GroupName + ":" + config.Azure.VMName
	imageId := config.Azure.Image.Publisher + ":" + config.Azure.Image.Offer + ":" + config.Azure.Image.Sku + ":" + config.Azure.Image.Version
	vmReqInfo := irs.VMReqInfo{
		Name: vmName,
		ImageInfo: irs.ImageInfo{
			Id: imageId,
		},
		SpecID: config.Azure.VMSize,
		VNetworkInfo: irs.VNetworkInfo{
			Id: config.Azure.Network.ID,
		},
		LoginInfo: irs.LoginInfo{
			AdminUsername: config.Azure.Os.AdminUsername,
			AdminPassword: config.Azure.Os.AdminPassword,
		},
	}

	vmHandler.StartVM(vmReqInfo)
}

func setVMHandler() (irs.VMHandler, error) {
	var cloudDriver idrv.CloudDriver
	cloudDriver = new(azdrv.AzureDriver)

	config := readConfigFile()
	connectionInfo := idrv.ConnectionInfo{
		CredentialInfo: idrv.CredentialInfo{
			SubscriptionId: config.Azure.SubscriptionID,
		},
		RegionInfo: idrv.RegionInfo{
			Region: config.Azure.Location,
		},
	}

	cloudConnection, _ := cloudDriver.ConnectCloud(connectionInfo)
	vmHandler, err := cloudConnection.CreateVMHandler()
	if err != nil {
		return nil, err
	}
	return vmHandler, nil
}

func main() {
	getVMInfo()
	//handleVM()
	//createVM()
}

type Config struct {
	Azure struct {
		SubscriptionID string `yaml:"subscription_id"`
		GroupName      string `yaml:"group_name"`
		VMName         string `yaml:"vm_name"`

		Location string `yaml:"location"`
		VMSize   string `yaml:"vm_size"`
		Image    struct {
			Publisher string `yaml:"publisher"`
			Offer     string `yaml:"offer"`
			Sku       string `yaml:"sku"`
			Version   string `yaml:"version"`
		} `yaml:"image"`
		Os struct {
			ComputeName   string `yaml:"compute_name"`
			AdminUsername string `yaml:"admin_username"`
			AdminPassword string `yaml:"admin_password"`
		} `yaml:"os"`
		Network struct {
			ID      string `yaml:"id"`
			Primary bool   `yaml:"primary"`
		} `yaml:"network"`
		ServerId string `yaml:"server_id"`
	} `yaml:"azure"`
}

func readConfigFile() Config {
	// Set Environment Value of Project Root Path
	rootPath := os.Getenv("CBSPIDER_PATH")
	data, err := ioutil.ReadFile(rootPath + "/config/config.yaml")
	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	return config
}
