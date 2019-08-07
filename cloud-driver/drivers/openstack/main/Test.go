package main

import (
	"fmt"
	config "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/config"
	osdrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/openstack"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

// Test OpenStack Connection
/*func TestConnection() {
	client, err := config.GetServiceClient()
	if err != nil {
		panic(err)
	}
	fmt.Println(client)
}*/

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

	// Get VM Info
	vmInfo := vmHandler.GetVM(config.Openstack.ServerId)
	spew.Dump(vmInfo)

	// Get VM Status List
	vmStatusList := vmHandler.ListVMStatus()
	for i, vmStatus := range vmStatusList {
		fmt.Println("[", i, "] ", *vmStatus)
	}

	// Get VM Status
	vmStatus := vmHandler.GetVMStatus(config.Openstack.ServerId)
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

		if inputCnt == 1 {
			switch commandNum {
			case 1:
				fmt.Println("Start Suspend VM ...")
				vmHandler.SuspendVM(config.Openstack.ServerId)
				fmt.Println("Finish Suspend VM")
			case 2:
				fmt.Println("Start Resume  VM ...")
				vmHandler.ResumeVM(config.Openstack.ServerId)
				fmt.Println("Finish Resume VM")
			case 3:
				fmt.Println("Start Reboot  VM ...")
				vmHandler.RebootVM(config.Openstack.ServerId)
				fmt.Println("Finish Reboot VM")
			case 4:
				fmt.Println("Start Terminate  VM ...")
				vmHandler.TerminateVM(config.Openstack.ServerId)
				fmt.Println("Finish Terminate VM")
			}
		}
	}
}

// Test VM Deployment
func createVM() {
	fmt.Println("Start Create VM ...")
	vmHandler, err := setVMHandler()
	if err != nil {
		panic(err)
	}

	config := config.ReadConfigFile()

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
	
	vm, err := vmHandler.StartVM(vmReqInfo)
	if err != nil {
		panic(err)
	}
	spew.Dump(vm)
	fmt.Println("Finish Create VM")
}

func getKeyPairInfo() {
	keypairHandler, err := setKeyPairHandler()
	if err != nil {
		panic(err)
	}

	req := irs.KeyPairReqInfo{
		//	Name: "keypair-name2",
	}
	keypairHandler.CreateKey(req)
	//keypairHandler.ListKey()
	//keypairHandler.GetKey("mcb-key")
	//keypairHandler.DeleteKey("ddddsa")
}

func setVMHandler() (irs.VMHandler, error) {
	var cloudDriver idrv.CloudDriver
	cloudDriver = new(osdrv.OpenStackDriver)

	config := config.ReadConfigFile()
	connectionInfo := idrv.ConnectionInfo{
		RegionInfo: idrv.RegionInfo{
			Region: config.Openstack.Region,
		},
	}

	cloudConnection, _ := cloudDriver.ConnectCloud(connectionInfo)
	vmHandler, err := cloudConnection.CreateVMHandler()
	if err != nil {
		return nil, err
	}
	return vmHandler, nil
}

func setKeyPairHandler() (irs.KeyPairHandler, error) {
	var cloudDriver idrv.CloudDriver
	cloudDriver = new(osdrv.OpenStackDriver)

	config := config.ReadConfigFile()
	connectionInfo := idrv.ConnectionInfo{
		RegionInfo: idrv.RegionInfo{
			Region: config.Openstack.Region,
		},
	}

	cloudConnection, _ := cloudDriver.ConnectCloud(connectionInfo)
	keyPairHandler, err := cloudConnection.CreateKeyPairHandler() //적용부분
	if err != nil {
		return nil, err
	}
	return keyPairHandler, nil
}

func setVNetworkHandler() (irs.VNetworkHandler, error) {
	var cloudDriver idrv.CloudDriver
	cloudDriver = new(osdrv.OpenStackDriver)

	config := config.ReadConfigFile()
	connectionInfo := idrv.ConnectionInfo{
		RegionInfo: idrv.RegionInfo{

			Region: config.Openstack.Region,
		},
	}

	cloudConnection, _ := cloudDriver.ConnectNetworkCloud(connectionInfo)
	vNetworkHandler, err := cloudConnection.CreateVNetworkHandler()
	if err != nil {
		return nil, err
	}
	return vNetworkHandler, nil
}

func VNetwork() {
	vNetworkHandler, err := setVNetworkHandler()
	if err != nil {
		panic(err)
	}

	//vNetworkHandler.CreateVNetwork(irs.VNetworkReqInfo{})
	//vNetworkHandler.GetVNetwork("b6610ceb-8089-48b0-9bfc-3c35e4e245cf")
	vNetworkHandler.ListVNetwork()
	//vNetworkHandler.DeleteVNetwork("b947ff7b-a586-4f98-828c-cdea04afc114")
}

func main() {
	//getVMInfo()
	//handleVM()
	//createVM()
	//TestImageHandler()
	getKeyPairInfo()
}

type Config struct {
	Openstack struct {
		DomainName       string `yaml:"domain_name"`
		IdentityEndpoint string `yaml:"identity_endpoint"`
		Password         string `yaml:"password"`
		ProjectID        string `yaml:"project_id"`
		Username         string `yaml:"username"`
		Region           string `yaml:"region"`
		VMName           string `yaml:"vm_name"`
		ImageId          string `yaml:"image_id"`
		FlavorId         string `yaml:"flavor_id"`
		NetworkId        string `yaml:"network_id"`
		SecurityGroups   string `yaml:"security_groups"`
		KeypairName      string `yaml:"keypair_name"`

		ServerId string `yaml:"server_id"`
	} `yaml:"openstack"`
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

/*
func TestImageHandler() {
	// Config Driver Info
	var cloudDriver idrv.CloudDriver
	cloudDriver = new(osdrv.OpenStackDriver)
	// Config Connection
	connectionInfo := idrv.ConnectionInfo{}
	cloudConnection, _ := cloudDriver.ConnectCloud(connectionInfo)
	// Load Handler (VM, Image, KeyPair ..)
	imageHandler, err := cloudConnection.CreateImageHandler()
	if err != nil {
		panic(err)
	}
	config := config.ReadConfigFile()
	// Use Handler Func
	//fmt.Println("Call CreateImage()")
	//reqParams := irs.ImageReqInfo{}
	//result, err := imageHandler.CreateImage(reqParams)
	//fmt.Println(result)
	fmt.Println("Call ListImage()")
	imageHandler.ListImage()
	fmt.Println("Call GetImage()")
	imageHandler.GetImage(config.Openstack.ImageId)
	//fmt.Println("Call DeleteImage()")
	//imageHandler.DeleteImage(config.Openstack.ImageId)
}*/
