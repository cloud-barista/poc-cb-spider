package main

import (
	"fmt"
	azdrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/azure"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

func testImageHandler(config Config) {
	resourceHandler, err := getResourceHandler("image")
	if err != nil {
		panic(err)
	}

	imageHandler := resourceHandler.(irs.ImageHandler)
 
	fmt.Println("Test ImageHandler")
	fmt.Println("1. ListImage()")
	fmt.Println("2. GetImage()")
	fmt.Println("3. CreateImage()")
	fmt.Println("4. DeleteImage()")
	fmt.Println("5. Exit")
	
	imageId := config.Azure.ImageInfo.GroupName + ":" + config.Azure.ImageInfo.Name
	
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
				reqInfo := irs.ImageReqInfo{Id: imageId}
				_, err := imageHandler.CreateImage(reqInfo)
				if err != nil {
					panic(err)
				}
				fmt.Println("Finish CreateImage()")
			case 4:
				fmt.Println("Start DeleteImage() ...")
				imageHandler.DeleteImage(imageId)
				fmt.Println("Finish DeleteImage()")
			case 5:
				fmt.Println("Exit")
				break Loop
			}
		}
	}
}

func testPublicIPHanlder(config Config) {
	resourceHandler, err := getResourceHandler("publicip")
	if err != nil {
		panic(err)
	}
	
	publicIPHandler := resourceHandler.(irs.PublicIPHandler)
	
	fmt.Println("Test PublicIPHandler")
	fmt.Println("1. ListPublicIP()")
	fmt.Println("2. GetPublicIP()")
	fmt.Println("3. CreatePublicIP()")
	fmt.Println("4. DeletePublicIP()")
	fmt.Println("5. Exit")
	
	publicIPId := config.Azure.PublicIP.GroupName + ":" + config.Azure.PublicIP.Name

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
				fmt.Println("Exit")
				break Loop
			}
		}
	}
}

func testSecurityHandler(config Config) {
	resourceHandler, err := getResourceHandler("security")
	if err != nil {
		panic(err)
	}
	
	securityHandler := resourceHandler.(irs.SecurityHandler)
	
	fmt.Println("Test SecurityHandler")
	fmt.Println("1. ListSecurity()")
	fmt.Println("2. GetSecurity()")
	fmt.Println("3. CreateSecurity()")
	fmt.Println("4. DeleteSecurity()")
	fmt.Println("5. Exit")
	
	securityGroupId := config.Azure.Security.GroupName + ":" + config.Azure.Security.Name
	
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
				reqInfo := irs.SecurityReqInfo{Id: securityGroupId}
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
				fmt.Println("Exit")
				break Loop
			}
		}
	}
}

func testVNetworkHandler(config Config) {
	resourceHandler, err := getResourceHandler("vnetwork")
	if err != nil {
		panic(err)
	}
	
	vNetworkHandler := resourceHandler.(irs.VNetworkHandler)
	
	fmt.Println("Test VNetworkHandler")
	fmt.Println("1. ListVNetwork()")
	fmt.Println("2. GetVNetwork()")
	fmt.Println("3. CreateVNetwork()")
	fmt.Println("4. DeleteVNetwork()")
	fmt.Println("5. Exit")

	vNetworkId := config.Azure.VNetwork.GroupName + ":" + config.Azure.VNetwork.Name
	
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
				reqInfo := irs.VNetworkReqInfo{Id: vNetworkId}
				_, err := vNetworkHandler.CreateVNetwork(reqInfo)
				if err != nil {
					panic(err)
				}
				fmt.Println("Finish CreateVNetwork()")
			case 4:
				fmt.Println("Start DeleteVNetwork() ...")
				vNetworkHandler.DeleteVNetwork(vNetworkId)
				fmt.Println("Finish DeleteVNetwork()")
			case 5:
				fmt.Println("Exit")
				break Loop
			}
		}
	}
}

func testVNicHandler(config Config) {
	resourceHandler, err := getResourceHandler("vnic")
	if err != nil {
		panic(err)
	}
	
	vNicHandler := resourceHandler.(irs.VNicHandler)
	
	fmt.Println("Test VNicHandler")
	fmt.Println("1. ListVNic()")
	fmt.Println("2. GetVNic()")
	fmt.Println("3. CreateVNic()")
	fmt.Println("4. DeleteVNic()")
	fmt.Println("5. Exit Program")
	
	vNicId := config.Azure.VNic.GroupName + ":" + config.Azure.VNic.Name
	
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
				fmt.Println("Start ListVNic() ...")
				vNicHandler.ListVNic()
				fmt.Println("Finish ListVNic()")
			case 2:
				fmt.Println("Start GetVNic() ...")
				vNicHandler.GetVNic(vNicId)
				fmt.Println("Finish GetVNic()")
			case 3:
				fmt.Println("Start CreateVNic() ...")
				reqInfo := irs.VNicReqInfo{Id: vNicId}
				_, err := vNicHandler.CreateVNic(reqInfo)
				if err != nil {
					panic(err)
				}
				fmt.Println("Finish CreateVNic()")
			case 4:
				fmt.Println("Start DeleteVNic() ...")
				vNicHandler.DeleteVNic(vNicId)
				fmt.Println("Finish DeleteVNic()")
			case 5:
				fmt.Println("Exit Program")
				break Loop
			}
		}
	}
}

func getResourceHandler(resourceType string) (interface{}, error) {
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

	var resourceHandler interface{}
	var err error

	switch resourceType {
	case "image":
		resourceHandler, err = cloudConnection.CreateImageHandler()
	case "publicip":
		resourceHandler, err = cloudConnection.CreatePublicIPHandler()
	case "security":
		resourceHandler, err = cloudConnection.CreateSecurityHandler()
	case "vnetwork":
		resourceHandler, err = cloudConnection.CreateVNetworkHandler()
	case "vnic":
		resourceHandler, err = cloudConnection.CreateVNicHandler()
	}
	
	if err != nil {
		return nil, err
	}
	return resourceHandler, nil
}

func showTestHandlerInfo() {
	fmt.Println("==========================================================")
	fmt.Println("[Test ResourceHandler]")
	fmt.Println("1. ImageHandler")
	fmt.Println("2. PublicIPHandler")
	fmt.Println("3. SecurityHandler")
	fmt.Println("4. VNetworkHandler")
	fmt.Println("5. VNicHandler")
	fmt.Println("6. Exit")
	fmt.Println("==========================================================")
}

func main() {

	showTestHandlerInfo()      // ResourceHandler 테스트 정보 출력
	config := readConfigFile() // config.yaml 파일 로드

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
				testImageHandler(config)
				showTestHandlerInfo()
			case 2:
				testPublicIPHanlder(config)
				showTestHandlerInfo()
			case 3:
				testSecurityHandler(config)
				showTestHandlerInfo()
			case 4:
				testVNetworkHandler(config)
				showTestHandlerInfo()
			case 5:
				testVNicHandler(config)
				showTestHandlerInfo()
			case 6:
				fmt.Println("Exit Test ResourceHandler Program")
				break Loop
			}
		}
	}
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
