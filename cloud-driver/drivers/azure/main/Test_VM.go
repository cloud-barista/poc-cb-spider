package main

import (
	"fmt"
	azdrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/azure"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

// Create Instance
func createVM(config Config, vmHandler irs.VMHandler) {
	
	vmName := config.Azure.GroupName + ":" + config.Azure.VMName
	imageId := config.Azure.Image.Publisher + ":" + config.Azure.Image.Offer + ":" + config.Azure.Image.Sku + ":" + config.Azure.Image.Version
	vmReqInfo := irs.VMReqInfo{
		Name: vmName,
		ImageInfo: irs.ImageInfo{
			Id: imageId,
		},
		SpecID: config.Azure.VMSize,
		VNetworkInfo: irs.VNetworkInfo{
			Id: config.Azure.Nic.ID,
		},
		LoginInfo: irs.LoginInfo{
			AdminUsername: config.Azure.Os.AdminUsername,
			//AdminPassword: config.Azure.Os.AdminPassword,
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
	
	fmt.Println("==========================================================")
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
	fmt.Println("==========================================================")
	
	vmId := config.Azure.GroupName + ":" + config.Azure.VMName
	
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
				vmInfo := vmHandler.GetVM(vmId)
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
				vmStatus := vmHandler.GetVMStatus(vmId)
				fmt.Println(vmStatus)
				fmt.Println("Finish Get VMStatus")
			case 5:
				fmt.Println("Start Create VM ...")
				createVM(config, vmHandler)
				fmt.Println("Finish Create VM")
			case 6:
				fmt.Println("Start Suspend VM ...")
				vmHandler.SuspendVM(vmId)
				fmt.Println("Finish Suspend VM")
			case 7:
				fmt.Println("Start Resume  VM ...")
				vmHandler.ResumeVM(vmId)
				fmt.Println("Finish Resume VM")
			case 8:
				fmt.Println("Start Reboot  VM ...")
				vmHandler.RebootVM(vmId)
				fmt.Println("Finish Reboot VM")
			case 9:
				fmt.Println("Start Terminate  VM ...")
				vmHandler.TerminateVM(vmId)
				fmt.Println("Finish Terminate VM")
			}
		}
	}
}

func getVMHandler() (irs.VMHandler, error) {
	var cloudDriver idrv.CloudDriver
	cloudDriver = new(azdrv.AzureDriver)
	
	config := readConfigFile()
	connectionInfo := idrv.ConnectionInfo{
		CredentialInfo: idrv.CredentialInfo{
			ClientId:       config.Azure.ClientId,
			ClientSecret:   config.Azure.ClientSecret,
			TenantId:       config.Azure.TenantId,
			SubscriptionId: config.Azure.SubscriptionID,
		},
		RegionInfo: idrv.RegionInfo{
			Region: config.Azure.Location,
			ResourceGroup: config.Azure.GroupName,
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

type Config struct {
	Azure struct {
		ClientId       string `yaml:"client_id"`
		ClientSecret   string `yaml:"client_secret"`
		TenantId       string `yaml:"tenant_id"`
		SubscriptionID string `yaml:"subscription_id"`
		
		GroupName string `yaml:"group_name"`
		VMName    string `yaml:"vm_name"`
		
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
		Nic struct {
			ID      string `yaml:"id"`
		} `yaml:"nic"`
		ServerId string `yaml:"server_id"`
		
		ImageInfo struct {
			GroupName string `yaml:"group_name"`
			Name      string `yaml:"name"`
		} `yaml:"image_info"`
		
		PublicIP struct {
			GroupName string `yaml:"group_name"`
			Name      string `yaml:"name"`
		} `yaml:"public_ip"`
		
		Security struct {
			GroupName string `yaml:"group_name"`
			Name      string `yaml:"name"`
		} `yaml:"security_group"`
		
		VNetwork struct {
			GroupName string `yaml:"group_name"`
			Name      string `yaml:"name"`
		} `yaml:"virtual_network"`
		
		VNic struct {
			GroupName string `yaml:"group_name"`
			Name      string `yaml:"name"`
		} `yaml:"network_interface"`
		
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
