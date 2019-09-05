package main

import (
	"fmt"
	osdrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/openstack"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/openstack/connect"
	osrs "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/openstack/resources"
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
				reqInfo := irs.ImageReqInfo{Name: config.Openstack.Image.Name}
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
				fmt.Println("Exit")
				break Loop
			}
		}
	}
}

func testKeyPairHandler(config Config) {
	resourceHandler, err := getResourceHandler("keypair")
	if err != nil {
		panic(err)
	}

	keyPairHandler := resourceHandler.(irs.KeyPairHandler)

	fmt.Println("Test KeyPairHandler")
	fmt.Println("1. ListKey()")
	fmt.Println("2. GetKey()")
	fmt.Println("3. CreateKey()")
	fmt.Println("4. DeleteKey()")
	fmt.Println("5. Exit")

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
				keyPairHandler.GetKey(config.Openstack.KeyPair.Name)
				fmt.Println("Finish GetKey()")
			case 3:
				fmt.Println("Start CreateKey() ...")
				reqInfo := irs.KeyPairReqInfo{Name: config.Openstack.KeyPair.Name}
				_, err := keyPairHandler.CreateKey(reqInfo)
				if err != nil {
					panic(err)
				}
				fmt.Println("Finish CreateKey()")
			case 4:
				fmt.Println("Start DeleteKey() ...")
				keyPairHandler.DeleteKey(config.Openstack.KeyPair.Name)
				fmt.Println("Finish DeleteKey()")
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
				reqInfo := irs.SecurityReqInfo{Name: config.Openstack.SecurityGroup.Name}
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
	fmt.Println("5. Exit")

	var vNicId string

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
				reqInfo := irs.VNicReqInfo{}
				vNic, err := vNicHandler.CreateVNic(reqInfo)
				if err != nil {
					panic(err)
				}
				vNicId = vNic.Id
				fmt.Println("Finish CreateVNic()")
			case 4:
				fmt.Println("Start DeleteVNic() ...")
				vNicHandler.DeleteVNic(vNicId)
				fmt.Println("Finish DeleteVNic()")
			case 5:
				fmt.Println("Exit")
				break Loop
			}
		}
	}
}

func testRouterHandler(config Config) {
	resourceHandler, err := getResourceHandler("router")
	if err != nil {
		panic(err)
	}

	routerHandler := resourceHandler.(osrs.OpenStackRouterHandler)

	fmt.Println("Test RouterHandler")
	fmt.Println("1. ListVNic()")
	fmt.Println("2. GetVNic()")
	fmt.Println("3. CreateVNic()")
	fmt.Println("4. DeleteVNic()")
	fmt.Println("5. Exit")

	var routerId string

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
				fmt.Println("Start ListRouter() ...")
				routerHandler.ListRouter()
				fmt.Println("Finish ListRouter()")
			case 2:
				fmt.Println("Start GetRouter() ...")
				routerHandler.GetRouter(routerId)
				fmt.Println("Finish GetRouter()")
			case 3:
				fmt.Println("Start CreateRouter() ...")
				reqInfo := irs.RouterReqInfo{}
				router, err := routerHandler.CreateRouter(reqInfo)
				if err != nil {
					panic(err)
				}
				routerId = router.Id
				fmt.Println("Finish CreateRouter()")
			case 4:
				fmt.Println("Start DeleteRouter() ...")
				routerHandler.DeleteRouter(routerId)
				fmt.Println("Finish DeleteRouter()")
			case 5:
				fmt.Println("Exit")
				break Loop
			}
		}
	}
}

func getResourceHandler(resourceType string) (interface{}, error) {
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

	var resourceHandler interface{}
	var err error

	switch resourceType {
	case "image":
		resourceHandler, err = cloudConnection.CreateImageHandler()
	case "keypair":
		resourceHandler, err = cloudConnection.CreateKeyPairHandler()
	case "publicip":
		resourceHandler, err = cloudConnection.CreatePublicIPHandler()
	case "security":
		resourceHandler, err = cloudConnection.CreateSecurityHandler()
	case "vnetwork":
		resourceHandler, err = cloudConnection.CreateVNetworkHandler()
	case "vnic":
		resourceHandler, err = cloudConnection.CreateVNicHandler()
	case "router":
		osDriver := osdrv.OpenStackDriver{}
		cloudConn, err := osDriver.ConnectCloud(connectionInfo)
		if err != nil {
			panic(err)
		}
		osCloudConn := cloudConn.(*connect.OpenStackCloudConnection)
		resourceHandler = osrs.OpenStackRouterHandler{Client: osCloudConn.NetworkClient}
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
	fmt.Println("2. KeyPairHandler")
	fmt.Println("3. PublicIPHandler")
	fmt.Println("4. SecurityHandler")
	fmt.Println("5. VNetworkHandler")
	fmt.Println("6. VNicHandler")
	fmt.Println("7. RouterHandler")
	fmt.Println("8. Exit")
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
				testKeyPairHandler(config)
				showTestHandlerInfo()
			case 3:
				testPublicIPHanlder(config)
				showTestHandlerInfo()
			case 4:
				testSecurityHandler(config)
				showTestHandlerInfo()
			case 5:
				testVNetworkHandler(config)
				showTestHandlerInfo()
			case 6:
				testVNicHandler(config)
				showTestHandlerInfo()
			case 7:
				testRouterHandler(config)
				showTestHandlerInfo()
			case 8:
				fmt.Println("Exit Test ResourceHandler Program")
				break Loop
			}
		}
	}
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
