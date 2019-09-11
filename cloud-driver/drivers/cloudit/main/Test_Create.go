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

func main() {
	testCreateVM()
}

func testCreateVM() {

	//리소스 핸들러 로드
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

	//imageHandler, _ := cloudConnection.CreateImageHandler()
	vNetworkHandler, _ := cloudConnection.CreateVNetworkHandler()
	securityHandler, _ := cloudConnection.CreateSecurityHandler()
	vmHandler, _ := cloudConnection.CreateVMHandler()
	publicIPHandler, _ := cloudConnection.CreatePublicIPHandler()
	//vNicHandler, _ := cloudConnection.CreateVNicHandler()

	// 1. Virtual Network 생성
	fmt.Println("start CreateVNetwork() ...")
	vNetReqInfo := irs.VNetworkReqInfo{Name: config.Cloudit.VirtualNetwork.Name}
	_, err := vNetworkHandler.CreateVNetwork(vNetReqInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("Finish CreateVNetwork()")

	// 2. Security Group 생성
	fmt.Println("Start CreateSecurity() ...")
	secReqInfo := irs.SecurityReqInfo{Name: config.Cloudit.securityGroup.Name}
	_, err = securityHandler.CreateSecurity(secReqInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("Finish CreateSecurity()")

	// 3. VM 생성
	fmt.Println("Start Create VM ...")
	vmReqInfo := irs.VMReqInfo{
		Name: config.Cloudit.VMInfo.Name,
		ImageInfo: irs.ImageInfo{
			Name: config.Cloudit.Resource.Image.Name,
		},
		SpecID: config.Cloudit.VMInfo.SpecId,
		VNetworkInfo: irs.VNetworkInfo{
			Id: config.Cloudit.VMInfo.SubnetAddr,
			// TODO: 생성된 Subnet Id 가져오가
			//Name: config.Cloudit.VirtualNetwork.Name,
		},
		SecurityInfo: irs.SecurityInfo{
			Id: config.Cloudit.VMInfo.SecGroups,
			// TODO: 생성된 SG Id 가져오가
			//Id: config.Cloudit.securityGroup.Name,
		},
		LoginInfo: irs.LoginInfo{
			AdminPassword: config.Cloudit.VMInfo.RootPassword,
		},
	}

	vm, err := vmHandler.StartVM(vmReqInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("Finish Create VM")

	// 4. Public IP 생성
	fmt.Println("Start CreatePublicIP() ...")
	publicIPReqInfo := irs.PublicIPReqInfo{Name: config.Cloudit.publicIp.Name}
	_, err = publicIPHandler.CreatePublicIP(publicIPReqInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("Finish CreatePublicIP()")

	spew.Dump(vm)
}

func cleanResource() {

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

		VirtualNetwork struct {
			Name string `yaml:"name"`
			Addr string `yaml:"addr"`
			ID   string `yaml:"id"`
		} `yaml:"vnet_info"`

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
		} `yaml:"resource"`

		VMInfo struct {
			TemplateId   string `yaml:"template_id"`
			SpecId       string `yaml:"spec_id"`
			Name         string `yaml:"name"`
			RootPassword string `yaml:"root_password"`
			SubnetAddr   string `yaml:"subnet_addr"`
			SecGroups    string `yaml:"sec_groups"`
			Description  string `yaml:"description"`
			Protection   int    `yaml:"protection"`
		} `yaml:"vm_info"`
	} `yaml:"cloudit"`
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
