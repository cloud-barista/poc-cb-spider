package main

import (
	"fmt"
	osdrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/openstack"
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
/*func createVM() {
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
*/
func testImageHandler() {
	imageHandler, err := setImageHandler()
	if err != nil {
		panic(err)
	}
	config := readConfigFile()

	fmt.Println("Test ImageHandler")
	fmt.Println("1. ListImage()")
	fmt.Println("2. GetImage()")
	fmt.Println("3. CreateImage()")
	fmt.Println("4. DeleteImage()")
	fmt.Println("5. Exit Program")

	var imageId string

Loop:
	for {
		var commandNum int
		inputCnt, err := fmt.Scan(&commandNum)
		if err != nil {
			panic(err)
		}

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
				fmt.Println("Start CreateImage() ...")
				reqInfo := irs.ImageReqInfo{
					Name: config.Openstack.Image.Name,
				}
				image, err := imageHandler.CreateImage(reqInfo)
				if err != nil {
					panic(err)
				}
				imageId = image.Id
				fmt.Println("Finish CreateImage()")
			case 4:
				fmt.Println("Start DeleteImage() ...")
				imageHandler.DeleteImage(imageId)
				fmt.Println("Finish DeleteImage()")
			case 5:
				fmt.Println("Exit Program")
				break Loop
			}
		}
	}
}

func testKeyPairHandler() {
	keyPairHandler, err := setKeyPairHandler()
	if err != nil {
		panic(err)
	}
	config := readConfigFile()

	fmt.Println("Test KeyPairHandler")
	fmt.Println("1. ListKey()")
	fmt.Println("2. GetKey()")
	fmt.Println("3. CreateKey()")
	fmt.Println("4. DeleteKey()")
	fmt.Println("5. Exit Program")

Loop:
	for {
		var commandNum int
		inputCnt, err := fmt.Scan(&commandNum)
		if err != nil {
			panic(err)
		}

		if inputCnt == 1 {
			switch commandNum {
			case 1:
				fmt.Println("Start ListKey() ...")
				keyPairHandler.ListKey()
				fmt.Println("Finish ListKey()")
			case 2:
				fmt.Println("Start GetKey() ...")
				keyPairHandler.GetKey(config.Openstack.KeypairName)
				fmt.Println("Finish GetKey()")
			case 3:
				fmt.Println("Start CreateKey() ...")
				reqInfo := irs.KeyPairReqInfo{Name: config.Openstack.KeypairName}
				_, err := keyPairHandler.CreateKey(reqInfo)
				if err != nil {
					panic(err)
				}
				fmt.Println("Finish CreateKey()")
			case 4:
				fmt.Println("Start DeleteKey() ...")
				keyPairHandler.DeleteKey(config.Openstack.KeypairName)
				fmt.Println("Finish DeleteKey()")
			case 5:
				fmt.Println("Exit Program")
				break Loop
			}
		}
	}
}

func testPublicIPHanlder() {
	publicIPHandler, err := setPublicIPHandler()
	if err != nil {
		panic(err)
	}
	//config := readConfigFile()

	fmt.Println("Test PublicIPHandler")
	fmt.Println("1. ListPublicIP()")
	fmt.Println("2. GetPublicIP()")
	fmt.Println("3. CreatePublicIP()")
	fmt.Println("4. DeletePublicIP()")
	fmt.Println("5. Exit Program")
	
	var publicIPId string
	
Loop:
	for {
		var commandNum int
		inputCnt, err := fmt.Scan(&commandNum)
		if err != nil {
			panic(err)
		}

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
				reqInfo := irs.PublicIPReqInfo{}
				publicIP, err := publicIPHandler.CreatePublicIP(reqInfo)
				if err != nil {
					panic(err)
				}
				publicIPId = publicIP.Id
				fmt.Println("Finish CreatePublicIP()")
			case 4:
				fmt.Println("Start DeletePublicIP() ...")
				publicIPHandler.DeletePublicIP(publicIPId)
				fmt.Println("Finish DeletePublicIP()")
			case 5:
				fmt.Println("Exit Program")
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

	var securityGroupId string

Loop:

	for {
		var commandNum int
		inputCnt, err := fmt.Scan(&commandNum)
		if err != nil {
			panic(err)
		}

		if inputCnt == 1 {
			switch commandNum {
			case 1:
				fmt.Println("Start ListSecurity() ...")
				securityHandler.ListSecurity()
				fmt.Println("Finish ListSecurity()")
			case 2:
				fmt.Println("Start GetSecurity() ...")
				securityHandler.GetSecurity(securityGroupId)
				fmt.Println("Finish GetSecurity()")
			case 3:
				fmt.Println("Start CreateSecurity() ...")
				reqInfo := irs.SecurityReqInfo{Name: config.Openstack.SecurityGroups}
				securityGroup, err := securityHandler.CreateSecurity(reqInfo)
				if err != nil {
					panic(err)
				}
				securityGroupId = securityGroup.Id
				fmt.Println("Finish CreateSecurity()")
			case 4:
				fmt.Println("Start DeleteSecurity() ...")
				securityHandler.DeleteSecurity(securityGroupId)
				fmt.Println("Finish DeleteSecurity()")
			case 5:
				fmt.Println("Exit Program")
				break Loop
			}
		}
	}
}

func testVNetworkHandler() {
	vNetworkHandler, err := setVNetworkHandler()
	if err != nil {
		panic(err)
	}
	config := readConfigFile()
	
	fmt.Println("Test VNetworkHandler")
	fmt.Println("1. ListVNetwork()")
	fmt.Println("2. GetVNetwork()")
	fmt.Println("3. CreateVNetwork()")
	fmt.Println("4. DeleteVNetwork()")
	fmt.Println("5. Exit Program")
	
	var vNetworkId string

Loop:
	
	for {
		var commandNum int
		inputCnt, err := fmt.Scan(&commandNum)
		if err != nil {
			panic(err)
		}
		
		if inputCnt == 1 {
			switch commandNum {
			case 1:
				fmt.Println("Start ListVNetwork() ...")
				vNetworkHandler.ListVNetwork()
				fmt.Println("Finish ListVNetwork()")
			case 2:
				fmt.Println("Start GetVNetwork() ...")
				vNetworkHandler.GetVNetwork(vNetworkId)
				fmt.Println("Finish GetVNetwork()")
			case 3:
				fmt.Println("Start CreateVNetwork() ...")
				reqInfo := irs.VNetworkReqInfo{Name: config.Openstack.VirtualNetwork.Name}
				vNetwork, err := vNetworkHandler.CreateVNetwork(reqInfo)
				if err != nil {
					panic(err)
				}
				vNetworkId = vNetwork.Id
				fmt.Println("Finish CreateVNetwork()")
			case 4:
				fmt.Println("Start DeleteVNetwork() ...")
				vNetworkHandler.DeleteVNetwork(vNetworkId)
				fmt.Println("Finish DeleteVNetwork()")
			case 5:
				fmt.Println("Exit Program")
				break Loop
			}
		}
	}
}

func setVMHandler() (irs.VMHandler, error) {
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

func setImageHandler() (irs.ImageHandler, error) {
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
	imageHandler, err := cloudConnection.CreateImageHandler()
	if err != nil {
		return nil, err
	}
	return imageHandler, nil
}

func setKeyPairHandler() (irs.KeyPairHandler, error) {
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
	keyPairHandler, err := cloudConnection.CreateKeyPairHandler()
	if err != nil {
		return nil, err
	}
	return keyPairHandler, nil
}

func setPublicIPHandler() (irs.PublicIPHandler, error) {
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
	publicIPHandler, err := cloudConnection.CreatePublicIPHandler()
	if err != nil {
		return nil, err
	}

	return publicIPHandler, nil
}

func setSecurityHandler() (irs.SecurityHandler, error) {
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
	securityHandler, err := cloudConnection.CreateSecurityHandler()
	if err != nil {
		return nil, err
	}

	return securityHandler, nil
}

func setVNetworkHandler() (irs.VNetworkHandler, error) {
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
	vNetworkHandler, err := cloudConnection.CreateVNetworkHandler()
	if err != nil {
		return nil, err
	}
	return vNetworkHandler, nil
}

func main() {
	//getVMInfo()
	//handleVM()
	//createVM()
	//testImageHandler()
	//testKeyPairHandler()
	//testPublicIPHanlder()
	//testSecurityHandler()
	testVNetworkHandler()
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

		ServerId   string `yaml:"server_id"`
		PublicIPID string `yaml:"public_ip_id"`

		Image struct {
			Id   string `yaml:"id"`
			Name string `yaml:"name"`
		} `yaml:"image_info"`
		
		VirtualNetwork struct {
			Name string `yaml:"name"`
		} `yaml:"virtual_network"`
		
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
