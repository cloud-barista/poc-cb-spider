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
				fmt.Println("Suspend Suspend VM ...")
				vmHandler.SuspendVM(config.Openstack.ServerId)
			case 2:
				fmt.Println("Resume  VM ...")
				vmHandler.ResumeVM(config.Openstack.ServerId)
			case 3:
				fmt.Println("Reboot  VM ...")
				vmHandler.RebootVM(config.Openstack.ServerId)
			case 4:
				fmt.Println("Terminate  VM ...")
				vmHandler.TerminateVM(config.Openstack.ServerId)
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

func getKeyPairInfo() {
	keypairHandler, err := setKeyPairHandler()
	if err != nil {
		panic(err)
	}

	req := irs.KeyPairReqInfo{}

	keypairHandler.CreateKey(req)
	//keypairHandler.ListKey()
	//keypairHandler.GetKey("test111")
	//keypairHandler.DeleteKey("ddddsa")
}

func getPublicIPInfo() {
	publicIPHandler, err := setPublicIPHandler()
	if err != nil {
		panic(err)
	}

	//publicIPHandler.ListVNetwork()
	publicIPHandler.GetVNetwork("dff3823e-29fb-40ef-af9b-a9f2250c4f79") //ID로 검색
	//publicIPHandler.DeleteVNetwork("")

	//pool 생성 (확인요함) 404에러
	//publicIPHandler.CreatePublicIP(irs.PublicIPReqInfo{})

}

func getSecurityInfo() {
	securityHandler, err := setSecurityHandler()
	if err != nil {
		panic(err)
	}

	//req := irs.SecurityReqInfo{
	//
	//}

	//securityHandler.CreateSecurity(req)
	//securityHandler.ListSecurity()
	//securityHandler.GetSecurity("e7d2752a-4d21-4e81-9c44-c274205f6d52")	//그룹 아이디로 검색
	securityHandler.DeleteSecurity("e7d2752a-4d21-4e81-9c44-c274205f6d52") //그룹 아이디로 삭제
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

func setPublicIPHandler() (irs.PublicIPHandler, error) {
	var cloudDriver idrv.CloudDriver
	cloudDriver = new(osdrv.OpenStackDriver)

	config := config.ReadConfigFile()
	connectionInfo := idrv.ConnectionInfo{
		RegionInfo: idrv.RegionInfo{
			Region: config.Openstack.Region,
		},
	}

	cloudConnection, _ := cloudDriver.ConnectCloud(connectionInfo)
	publicIPHandler, err := cloudConnection.CreatePublicIPHandler() //적용부분
	if err != nil {
		return nil, err
	}

	return publicIPHandler, nil
}

func setSecurityHandler() (irs.SecurityHandler, error) {
	var cloudDriver idrv.CloudDriver
	cloudDriver = new(osdrv.OpenStackDriver)

	config := config.ReadConfigFile()
	connectionInfo := idrv.ConnectionInfo{
		RegionInfo: idrv.RegionInfo{
			Region: config.Openstack.Region,
		},
	}

	cloudConnection, _ := cloudDriver.ConnectCloud(connectionInfo)
	securityHandler, err := cloudConnection.CreateSecurityHandler() //적용부분
	if err != nil {
		return nil, err
	}

	return securityHandler, nil
}

func main() {
	//getVMInfo()
	//handleVM()
	//createVM()
	//TestImageHandler()
	//getKeyPairInfo()
	getPublicIPInfo()
	//getSecurityInfo()
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
