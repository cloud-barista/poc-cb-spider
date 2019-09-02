package main

import (
	"fmt"
	osdrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/openstack"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/davecgh/go-spew/spew"
)

// Create Instance
func createVM(config Config, vmHandler irs.VMHandler) {

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
}

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
				vmList := vmHandler.ListVM()
				for i, vm := range vmList {
					fmt.Println("[", i, "] ")
					spew.Dump(vm)
				}
				fmt.Println("Finish List VM")
			case 2:
				fmt.Println("Start Get VM ...")
				vmInfo := vmHandler.GetVM(config.Openstack.ServerId)
				spew.Dump(vmInfo)
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
				vmStatus := vmHandler.GetVMStatus(config.Openstack.ServerId)
				fmt.Println(vmStatus)
				fmt.Println("Finish Get VMStatus")
			case 5:
				fmt.Println("Start Create VM ...")
				createVM(config, vmHandler)
				fmt.Println("Finish Create VM")
			case 6:
				fmt.Println("Start Suspend VM ...")
				vmHandler.SuspendVM(config.Openstack.ServerId)
				fmt.Println("Finish Suspend VM")
			case 7:
				fmt.Println("Start Resume  VM ...")
				vmHandler.ResumeVM(config.Openstack.ServerId)
				fmt.Println("Finish Resume VM")
			case 8:
				fmt.Println("Start Reboot  VM ...")
				vmHandler.RebootVM(config.Openstack.ServerId)
				fmt.Println("Finish Reboot VM")
			case 9:
				fmt.Println("Start Terminate  VM ...")
				vmHandler.TerminateVM(config.Openstack.ServerId)
				fmt.Println("Finish Terminate VM")
			}
		}
	}
}

func getVMHandler() (irs.VMHandler, error) {
	var cloudDriver idrv.CloudDriver
	cloudDriver = new(osdrv.OpenStackDriver)

	config := readConfigFile()
	connectionInfo := idrv.ConnectionInfo{
		CredentialInfo: idrv.CredentialInfo{
			IdentityEndpoint: config.Openstack.IdentityEndpoint,
			Username:         config.Openstack.Username,
			Password:         config.Openstack.Password,
			DomainName:       config.Openstack.DomainName,
			ProjectID:        config.Openstack.ProjectID,
		},
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

func main() {
	testVMHandler()
}

/*
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

		ServerId   string `yaml:"server_id"`
		PublicIPID string `yaml:"public_ip_id"`

		Image struct {
			Name string `yaml:"name"`
		} `yaml:"image_info"`

		KeyPair struct {
			Name string `yaml:"name"`
		} `yaml:"keypair_info"`

		SecurityGroup struct {
			Name string `yaml:"name"`
		} `yaml:"security_group_info"`

		VirtualNetwork struct {
			Name string `yaml:"name"`
		} `yaml:"vnet_info"`
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
*/
