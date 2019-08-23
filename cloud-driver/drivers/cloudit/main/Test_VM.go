package main

import (
	"fmt"
	cidrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

func testVMHandler() {
	vmHandler, err := getVMHandler()
	if err != nil {
		panic(err)
	}
	config := readConfigFile()

	fmt.Println("Test VMHandler")
	fmt.Println("1. List VM")
	fmt.Println("2. Get VM")
	fmt.Println("3. List VMStatus")
	fmt.Println("4. Get VMStatus")
	fmt.Println("5. Create VM")
	fmt.Println("6. Suspend VM")
	fmt.Println("7. Resume VM")
	fmt.Println("8. Reboot VM")
	fmt.Println("9. Terminate VM")
	fmt.Println("10. Exit")

	for {
		var commandNum int
		inputCnt, err := fmt.Scan(&commandNum)
		if err != nil {
			panic(err)
		}

		if inputCnt == 1 {
			switch commandNum {
			case 1:
				fmt.Println("Start List VM ...")
				vmHandler.ListVM()
				/*vmList := vmHandler.ListVM()
				for i, vm := range vmList {
					fmt.Println("[", i, "] ")
					spew.Dump(vm)
				}*/
				fmt.Println("Finish List VM")
			case 2:
				fmt.Println("Start Get VM ...")
				/*vmInfo := vmHandler.GetVM(config.Cloudit.ServerId)
				spew.Dump(vmInfo)*/
				vmHandler.GetVM(config.Cloudit.ServerId)
				fmt.Println("Finish Get VM")
			case 3:
				fmt.Println("Start List VMStatus ...")
				vmStatusList := vmHandler.ListVMStatus()
				for i, vmStatus := range vmStatusList {
					fmt.Println("[", i, "] ", *vmStatus)
				}
				fmt.Println("Finish List VMStatus")
			case 4:
				fmt.Println("Start Get VMStatus ...")
				//vmStatus := vmHandler.GetVMStatus(config.Openstack.ServerId)
				//fmt.Println(vmStatus)
				fmt.Println("Finish Get VMStatus")
			case 5:
				fmt.Println("Start Create VM ...")
				//createVM(config, vmHandler)
				fmt.Println("Finish Create VM")
			case 6:
				fmt.Println("Start Suspend VM ...")
				//vmHandler.SuspendVM(config.Openstack.ServerId)
				fmt.Println("Finish Suspend VM")
			case 7:
				fmt.Println("Start Resume  VM ...")
				//vmHandler.ResumeVM(config.Openstack.ServerId)
				fmt.Println("Finish Resume VM")
			case 8:
				fmt.Println("Start Reboot  VM ...")
				//vmHandler.RebootVM(config.Openstack.ServerId)
				fmt.Println("Finish Reboot VM")
			case 9:
				fmt.Println("Start Terminate  VM ...")
				//vmHandler.TerminateVM(config.Openstack.ServerId)
				fmt.Println("Finish Terminate VM")
			}
		}
	}
}

func getVMHandler() (irs.VMHandler, error) {
	var cloudDriver idrv.CloudDriver
	cloudDriver = new(cidrv.ClouditDriver)

	config := readConfigFile()
	connectionInfo := idrv.ConnectionInfo{
		CredentialInfo: idrv.CredentialInfo{
			IdentityEndpoint: config.Cloudit.IdentityEndpoint,
			Username:         config.Cloudit.Username,
			Password:         config.Cloudit.Password,
			TenantId:        config.Cloudit.TenantID,
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
	testVMHandler()
}

/*type Config struct {
	Cloudit struct {
		IdentityEndpoint string `yaml:"identity_endpoint"`
		Username         string `yaml:"user_id"`
		Password         string `yaml:"password"`
		TenantID         string `yaml:"tenant_id"`
		ServerId string `yaml:"server_id"`
	} `yaml:"cloudit"`
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
*/