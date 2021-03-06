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

func testImageHandler(config Config) {

	var imageHandler irs.ImageHandler
	if resourceHandler, err := getResourceHandler("image"); err != nil {
		panic(err)
	} else {
		imageHandler = resourceHandler.(irs.ImageHandler)
	}

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
				if _, err := imageHandler.ListImage(); err != nil {
					panic(err)
				}
				fmt.Println("Finish ListImage()")
			case 2:
				fmt.Println("Start GetImage() ...")
				if _, err := imageHandler.GetImage(imageId); err != nil {
					panic(err)
				}
				fmt.Println("Finish GetImage()")
			case 3:
				fmt.Println("Start CreateImage() ...")
				reqInfo := irs.ImageReqInfo{Name: config.Cloudit.Resource.Image.Name}
				if image, err := imageHandler.CreateImage(reqInfo); err != nil {
					panic(err)
				} else {
					imageId = image.Id
				}
				fmt.Println("Finish CreateImage()")
			case 4:
				fmt.Println("Start DeleteImage() ...")
				if ok, err := imageHandler.DeleteImage(imageId); !ok {
					panic(err)
				}
				fmt.Println("Finish DeleteImage()")
			case 5:
				fmt.Println("Exit")
				break Loop
			}
		}
	}

}

//AdaptiveIP
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
				if _, err := publicIPHandler.ListPublicIP(); err != nil {
					panic(err)
				}
				fmt.Println("Finish ListPublicIP()")
			case 2:
				fmt.Println("Start GetPublicIP() ...")
				if _, err := publicIPHandler.GetPublicIP(publicIPId); err != nil {
					panic(err)
				}
				fmt.Println("Finish GetPublicIP()")
			case 3:
				fmt.Println("Start CreatePublicIP() ...")
				reqInfo := irs.PublicIPReqInfo{Name: config.Cloudit.Resource.PublicIP.Name}
				if publicIP, err := publicIPHandler.CreatePublicIP(reqInfo); err != nil {
					panic(err)
				} else {
					publicIPId = publicIP.Id
				}
				fmt.Println("Finish CreatePublicIP()")
			case 4:
				fmt.Println("Start DeletePublicIP() ...")
				if ok, err := publicIPHandler.DeletePublicIP(publicIPId); !ok {
					panic(err)
				}
				fmt.Println("Finish DeletePublicIP()")
			case 5:
				fmt.Println("Exit")
				break Loop
			}
		}
	}
}

//SecurityGroup
func testSecurityHandler(config Config) {
	resourceHandler, err := getResourceHandler("security")
	if err != nil {
		panic(err)
	}

	securityHandler := resourceHandler.(irs.SecurityHandler)

	fmt.Println("Test securityHandler")
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
				if _, err := securityHandler.ListSecurity(); err != nil {
					panic(err)
				}
				fmt.Println("Finish ListSecurity()")
			case 2:
				fmt.Println("Start GetSecurity() ...")
				if _, err := securityHandler.GetSecurity(securityGroupId); err != nil {
					panic(err)
				}
				fmt.Println("Finish GetSecurity()")
			case 3:
				fmt.Println("Start CreateSecurity() ...")
				reqInfo := irs.SecurityReqInfo{Name: config.Cloudit.Resource.Security.Name}
				if security, err := securityHandler.CreateSecurity(reqInfo); err != nil {
					panic(err)
				} else {
					securityGroupId = security.Id
				}
				fmt.Println("Finish CreateSecurity()")
			case 4:
				fmt.Println("Start DeleteSecurity() ...")
				if ok, err := securityHandler.DeleteSecurity(securityGroupId); !ok {
					panic(err)
				}
				fmt.Println("Finish DeleteSecurity()")
			case 5:
				fmt.Println("Exit")
				break Loop
			}
		}
	}
}

//Subnet
func testVNetworkHandler(config Config) {
	resourceHandler, err := getResourceHandler("vnetwork")
	if err != nil {
		panic(err)
	}

	vNetworkHandler := resourceHandler.(irs.VNetworkHandler)

	fmt.Println("Test vNetworkHandler")
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
				if _, err := vNetworkHandler.ListVNetwork(); err != nil {
					panic(err)
				}
				fmt.Println("Finish ListVNetwork()")
			case 2:
				fmt.Println("Start GetVNetwork() ...")
				if _, err := vNetworkHandler.GetVNetwork(vNetworkId); err != nil {
					panic(err)
				}
				fmt.Println("Finish GetVNetwork()")
			case 3:
				fmt.Println("Start CreateVNetwork() ...")
				reqInfo := irs.VNetworkReqInfo{Name: config.Cloudit.Resource.VirtualNetwork.Name}
				if vNetwork, err := vNetworkHandler.CreateVNetwork(reqInfo); err != nil {
					panic(err)
				} else {
					vNetworkId = vNetwork.Id
				}
				fmt.Println("Finish CreateVNetwork()")
			case 4:
				fmt.Println("Start DeleteVNetwork() ...")
				if ok, err := vNetworkHandler.DeleteVNetwork(vNetworkId); !ok {
					panic(err)
				}
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

	fmt.Println("Test vNetworkHandler")
	fmt.Println("1. ListVNic()")
	fmt.Println("2. GetVNic()")
	fmt.Println("3. CreateVNic()")
	fmt.Println("4. DeleteVNic()")
	fmt.Println("5. Exit")

	nicId := config.Cloudit.Resource.VNic.Mac
	
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
				if _, err := vNicHandler.ListVNic(); err != nil {
					panic(err)
				}
				fmt.Println("Finish ListVNic()")
			case 2:
				fmt.Println("Start GetVNic() ...")
				if _, err := vNicHandler.GetVNic(nicId); err != nil {
					panic(err)
				}
				fmt.Println("Finish GetVNic()")
			case 3:
				fmt.Println("Start CreateVNic() ...")
				reqInfo := irs.VNicReqInfo{}
				if _, err := vNicHandler.CreateVNic(reqInfo); err != nil {
					panic(err)
				}
				fmt.Println("Finish CreateVNic()")
			case 4:
				fmt.Println("Start DeleteVNic() ...")
				if ok, err := vNicHandler.DeleteVNic(nicId); !ok {
					panic(err)
				}
				fmt.Println("Finish DeleteVNic()")
			case 5:
				fmt.Println("Exit")
				break Loop
			}
		}
	}
}

func getResourceHandler(resourceType string) (interface{}, error) {
	var cloudDriver idrv.CloudDriver
	cloudDriver = new(cidrv.ClouditDriver)

	config := readConfigFile()
	connectionInfo := idrv.ConnectionInfo{
		CredentialInfo: idrv.CredentialInfo{
			IdentityEndpoint: config.Cloudit.IdentityEndpoint,
			Username:         config.Cloudit.Username,
			Password:         config.Cloudit.Password,
			TenantId:         config.Cloudit.TenantID,
			AuthToken:        config.Cloudit.AuthToken,
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
	Cloudit struct {
		IdentityEndpoint string `yaml:"identity_endpoint"`
		Username         string `yaml:"user_id"`
		Password         string `yaml:"password"`
		TenantID         string `yaml:"tenant_id"`
		ServerId         string `yaml:"server_id"`
		AuthToken        string `yaml:"auth_token"`

		Image struct {
			Name string `yaml:"name"`
			ID   string `yaml:"id"`
		} `yaml:"image_info"`

		/*VirtualNetwork struct {
			Name string `yaml:"name"`
			Addr string `yaml:"addr"`
			ID   string `yaml:"id"`
		} `yaml:"vnet_info"`*/

		publicIp struct {
			Name string `yaml:"name"`
			ID   string `yaml:"id"`
			IP   string `yaml:"ip"`
		} `yaml:"publicIp_info"`

		securityGroup struct {
			Name           string `yaml:"name"`
			ID             string `yaml:"id"`
			SecuiryGroupID string `yaml:"securitygroupid"`
		}

		Resource struct {
			Image struct {
				Name string `yaml:"name"`
			} `yaml:"image"`

			PublicIP struct {
				Name string `yaml:"name"`
			} `yaml:"public_ip"`

			Security struct {
				Name string `yaml:"name"`
			} `yaml:"security_group"`

			VirtualNetwork struct {
				Name string `yaml:"name"`
			} `yaml:"vnet_info"`
			
			VNic struct{
				Mac string `yaml:"mac"`
			} `yaml:"vnic_info"`
			
		} `yaml:"resource"`
	} `yaml:"cloudit"`
}

func readConfigFile() Config {
	// Set Environment Value of Project Root Path4
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
	//spew.Dump(config)
	return config
}
