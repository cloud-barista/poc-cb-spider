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
				fmt.Println("Finish Suspend VM")
			case 2:
				fmt.Println("Start Resume  VM ...")
				vmHandler.ResumeVM(vmId)
				fmt.Println("Finish Resume VM")
			case 3:
				fmt.Println("Start Reboot  VM ...")
				vmHandler.RebootVM(vmId)
				fmt.Println("Finish Reboot VM")
			case 4:
				fmt.Println("Start Terminate  VM ...")
				vmHandler.TerminateVM(vmId)
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
	config := readConfigFile()

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

	vm, err := vmHandler.StartVM(vmReqInfo)
	if err != nil {
		panic(err)
	}
	spew.Dump(vm)
	fmt.Println("Finish Create VM")
}

func testImageHandler() {
	imageHandler, err := setImageHandler()
	if err != nil {
		panic(err)
	}
	config := readConfigFile()

	fmt.Println("ImageHandler")
	fmt.Println("1. ListImage()")
	fmt.Println("2. GetImage()")
	fmt.Println("3. CreateImage()")
	fmt.Println("4. DeleteImage()")

	for {
		var commandNum int
		inputCnt, err := fmt.Scan(&commandNum)
		if err != nil {
			panic(err)
		}

		imageId := config.Azure.ImageInfo.GroupName + ":" + config.Azure.ImageInfo.Name

		if inputCnt == 1 {
			switch commandNum {
			case 1:
				fmt.Println("Start ListImage() ...")
				imageHandler.ListImage()
				fmt.Println("Finish ListImage()")
			case 2:
				fmt.Println("Start GetImage() ...")
				imageHandler.GetImage(imageId)
				fmt.Println("Finish GetImage()")
			case 3:
				/*fmt.Println("Start CreateImage() ...")
				fmt.Println("Finish CreateImage()")*/
			case 4:
				fmt.Println("Start DeleteImage() ...")
				imageHandler.DeleteImage(imageId)
				fmt.Println("Finish DeleteImage()")
			}
		}
	}
}

func testVNetworkHandler() {
	vNetHandler, err := setVNetHandler()
	if err != nil {
		panic(err)
	}
	config := readConfigFile()

	fmt.Println("ImageHandler")
	fmt.Println("1. ListVNetwork()")
	fmt.Println("2. GetVNetwork()")
	fmt.Println("3. CreateVNetwork()")
	fmt.Println("4. DeleteVNetwork()")
	fmt.Println("5. Exit Program")

Loop:
	for {
		var commandNum int
		inputCnt, err := fmt.Scan(&commandNum)
		if err != nil {
			panic(err)
		}

		networkId := config.Azure.VNetwork.GroupName + ":" + config.Azure.VNetwork.Name

		if inputCnt == 1 {
			switch commandNum {
			case 1:
				fmt.Println("Start ListVNetwork() ...")
				vNetHandler.ListVNetwork()
				fmt.Println("Finish ListVNetwork()")
			case 2:
				fmt.Println("Start GetVNetwork() ...")
				vNetHandler.GetVNetwork(networkId)
				fmt.Println("Finish GetVNetwork()")
			case 3:
				fmt.Println("Start CreateVNetwork() ...")
				reqInfo := irs.VNetworkReqInfo{Id: networkId}
				_, err := vNetHandler.CreateVNetwork(reqInfo)
				if err != nil {
					panic(err)
				}
				fmt.Println("Finish CreateVNetwork()")
			case 4:
				fmt.Println("Start DeleteVNetwork() ...")
				vNetHandler.DeleteVNetwork(networkId)
				fmt.Println("Finish DeleteVNetwork()")
			case 5:
				break Loop
			}
		}
	}
}

func testPublicIPHandler() {
	publicIPHandler, err := setPublicIPHandler()
	if err != nil {
		panic(err)
	}
	config := readConfigFile()

	fmt.Println("ImageHandler")
	fmt.Println("1. ListPublicIP()")
	fmt.Println("2. GetPublicIP()")
	fmt.Println("3. CreatePublicIP()")
	fmt.Println("4. DeletePublicIP()")
	fmt.Println("5. Exit Program")

Loop:
	for {
		var commandNum int
		inputCnt, err := fmt.Scan(&commandNum)
		if err != nil {
			panic(err)
		}

		publicIPId := config.Azure.PublicIP.GroupName + ":" + config.Azure.PublicIP.Name

		if inputCnt == 1 {
			switch commandNum {
			case 1:
				fmt.Println("Start ListPublicIP() ...")
				publicIPHandler.ListPublicIP()
				fmt.Println("Finish ListPublicIP()")
			case 2:
				fmt.Println("Start GetPublicIP() ...")
				publicIPHandler.GetPublicIP(publicIPId)
				fmt.Println("Finish GetPublicIP()")
			case 3:
				fmt.Println("Start CreatePublicIP() ...")
				reqInfo := irs.PublicIPReqInfo{Id: publicIPId}
				_, err := publicIPHandler.CreatePublicIP(reqInfo)
				if err != nil {
					panic(err)
				}
				fmt.Println("Finish CreatePublicIP()")
			case 4:
				fmt.Println("Start DeletePublicIP() ...")
				publicIPHandler.DeletePublicIP(publicIPId)
				fmt.Println("Finish DeletePublicIP()")
			case 5:
				break Loop
			}
		}
	}
}

func testSecurityHandler() {
	securityHandler, err := setSecurityHandler()
	if err != nil {
		panic(err)
	}
	config := readConfigFile()

	fmt.Println("Test SecurityHandler")
	fmt.Println("1. ListSecurity()")
	fmt.Println("2. GetSecurity()")
	fmt.Println("3. CreateSecurity()")
	fmt.Println("4. DeleteSecurity()")
	fmt.Println("5. Exit Program")

Loop:
	for {
		var commandNum int
		inputCnt, err := fmt.Scan(&commandNum)
		if err != nil {
			panic(err)
		}

		securityId := config.Azure.Security.GroupName + ":" + config.Azure.Security.Name

		if inputCnt == 1 {
			switch commandNum {
			case 1:
				fmt.Println("Start ListSecurity() ...")
				securityHandler.ListSecurity()
				fmt.Println("Finish ListSecurity()")
			case 2:
				fmt.Println("Start GetSecurity() ...")
				securityHandler.GetSecurity(securityId)
				fmt.Println("Finish GetSecurity()")
			case 3:
				fmt.Println("Start CreateSecurity() ...")
				reqInfo := irs.SecurityReqInfo{Id: securityId}
				_, err := securityHandler.CreateSecurity(reqInfo)
				if err != nil {
					panic(err)
				}
				fmt.Println("Finish CreateSecurity()")
			case 4:
				fmt.Println("Start DeleteSecurity() ...")
				securityHandler.DeleteSecurity(securityId)
				fmt.Println("Finish DeleteSecurity()")
			case 5:
				break Loop
			}
		}
	}
}

func getVNicInfo() {
	vNicHandler, err := setVNicHandler()
	if err != nil {
		panic(err)
	}
	config := readConfigFile()

	//vNicHandler.ListVNic()

	vNicId := config.Azure.VNic.GroupName + ":" + config.Azure.VNic.Name
	vNicHandler.GetVNic(vNicId)
}

func createVNicInfo() {
	vNicHandler, err := setVNicHandler()
	if err != nil {
		panic(err)
	}
	config := readConfigFile()

	vNicId := config.Azure.VNic.GroupName + ":" + config.Azure.VNic.Name
	reqInfo := irs.VNicReqInfo{
		Id: vNicId,
	}

	result, err := vNicHandler.CreateVNic(reqInfo)
	if err != nil {
		panic(err)
	}
	spew.Dump(result)
}

func setVMHandler() (irs.VMHandler, error) {
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
		},
	}

	cloudConnection, err := cloudDriver.ConnectCloud(connectionInfo)
	if err != nil {
		return nil, err
	}
	vmHandler, err := cloudConnection.CreateVMHandler()
	if err != nil {
		return nil, err
	}
	return vmHandler, nil
}

func setImageHandler() (irs.ImageHandler, error) {
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
		},
	}

	cloudConnection, err := cloudDriver.ConnectCloud(connectionInfo)
	if err != nil {
		return nil, err
	}
	imageHandler, err := cloudConnection.CreateImageHandler()
	if err != nil {
		return nil, err
	}
	return imageHandler, nil
}

func setPublicIPHandler() (irs.PublicIPHandler, error) {
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
		},
	}

	cloudConnection, err := cloudDriver.ConnectCloud(connectionInfo)
	if err != nil {
		return nil, err
	}
	publicIPHandler, err := cloudConnection.CreatePublicIPHandler()
	if err != nil {
		return nil, err
	}
	return publicIPHandler, nil
}

func setSecurityHandler() (irs.SecurityHandler, error) {
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
		},
	}

	cloudConnection, err := cloudDriver.ConnectCloud(connectionInfo)
	if err != nil {
		return nil, err
	}
	securityHandler, err := cloudConnection.CreateSecurityHandler()
	if err != nil {
		return nil, err
	}
	return securityHandler, nil
}

func setVNicHandler() (irs.VNicHandler, error) {
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
		},
	}

	cloudConnection, err := cloudDriver.ConnectCloud(connectionInfo)
	if err != nil {
		return nil, err
	}
	vNicHandler, err := cloudConnection.CreateVNicHandler()
	if err != nil {
		return nil, err
	}
	return vNicHandler, nil
}

func setVNetHandler() (irs.VNetworkHandler, error) {
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
		},
	}

	cloudConnection, err := cloudDriver.ConnectCloud(connectionInfo)
	if err != nil {
		return nil, err
	}
	vNetHandler, err := cloudConnection.CreateVNetworkHandler()
	if err != nil {
		return nil, err
	}
	return vNetHandler, nil
}

func main() {

	//getVMInfo()
	//handleVM()
	//createVM()

	//testImageHandler()
	//testVNetworkHandler()
	testPublicIPHandler()
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
		Network struct {
			ID      string `yaml:"id"`
			Primary bool   `yaml:"primary"`
		} `yaml:"network"`
		ServerId string `yaml:"server_id"`

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

		ImageInfo struct {
			GroupName string `yaml:"group_name"`
			Name      string `yaml:"name"`
		} `yaml:"image_info"`
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
