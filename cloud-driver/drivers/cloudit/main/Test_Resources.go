package main

import (
	"fmt"
	cidrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

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

//SecurityGroup
func testSecurityHandler(config Config) {
	resourceHandler, err := getResourceHandler("security")
	if err != nil {
		panic(err)
	}

	securityHandler := resourceHandler.(irs.SecurityHandler)

	//fmt.Println("Test securityHandler")
	fmt.Println("1. ListSecurity()")
	fmt.Println("2. GetSecurity()")
	//fmt.Println("3. CreateSecurity()")
	//fmt.Println("4. DeleteSecurity()")
	fmt.Println("5. Exit")

	//var securityId string
	securityId := "b77a1163-01ef-4e14-9ffc-9fb626b367be"

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
				securityHandler.GetSecurity(securityId)
				fmt.Println("Finish GetSecurity()")
			case 3:
				fmt.Println("Start CreateSecurity() ...")
				reqInfo := irs.SecurityReqInfo{}
				security, err := securityHandler.CreateSecurity(reqInfo)
				if err != nil {
					panic(err)
				}
				securityId = security.Id
				fmt.Println("Finish CreateSecurity()")
			case 4:
				fmt.Println("Start DeleteSecurity() ...")
				securityHandler.DeleteSecurity(securityId)
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

	//fmt.Println("Test vNetworkHandler")
	fmt.Println("1. ListVNetwork()")
	fmt.Println("2. CreateVNetwork()")
	fmt.Println("3. -----VNetwork()")
	fmt.Println("4. DeleteVNetwork()")
	fmt.Println("5. Exit")

	var vNetworkId string
	//vNetworkId := ""
	var addr string
	addr = "10.0.12.0"

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
				fmt.Println("Start CreateVNetwork() ...")
				reqInfo := irs.VNetworkReqInfo{Name: config.Cloudit.VirtualNetwork.Name}
				vNetwork, err := vNetworkHandler.CreateVNetwork(reqInfo)
				if err != nil {
					panic(err)
				}
				vNetworkId = vNetwork.Id
				fmt.Println("Finish CreateVNetwork()")

			case 3:
				fmt.Println("Start UpdateVNetwork() ...")
				vNetworkHandler.GetVNetwork(vNetworkId)
				fmt.Println("Finish UpdateVNetwork()")
			case 4:
				fmt.Println("Start DeleteVNetwork() ...")
				//vNetworkHandler.DeleteVNetwork(config.Cloudit.VirtualNetwork.Addr)
				vNetworkHandler.DeleteVNetwork(addr)
				fmt.Println("Finish DeleteVNetwork()")
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
	fmt.Println("2. KeyPairHandler")
	fmt.Println("3. PublicIPHandler")
	fmt.Println("4. SecurityHandler")
	fmt.Println("5. VNetworkHandler")
	fmt.Println("6. VNicHandler")
	fmt.Println("7. Exit")
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
				//testImageHandler(config)
				showTestHandlerInfo()
			case 2:
				//testKeyPairHandler(config)
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
				//testVNicHandler(config)
				resourceHandler, err := getResourceHandler("vnic")
				if err != nil {
					panic(err)
				}
				vNicInfoList, err := resourceHandler.(irs.VNicHandler).ListVNic()
				if err != nil {
					panic(err)
				}
				for i, vNicInfo := range vNicInfoList {
					fmt.Println("[", i, "] ", *vNicInfo)
				}
				showTestHandlerInfo()

			case 7:
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

		VirtualNetwork struct {
			Name string `yaml:"name"`
			Addr string `yaml:"addr"`
		} `yaml:"vnet_info"`
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
	spew.Dump(config)
	return config
}
