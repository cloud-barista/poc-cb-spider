package resources

import (
	"fmt"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/config"
	_ "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/config"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/davecgh/go-spew/spew"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/compute/v2/extensions/floatingip"
	_ "github.com/rackspace/gophercloud/openstack/compute/v2/extensions/keypairs"
	_ "github.com/rackspace/gophercloud/openstack/compute/v2/extensions/startstop"
	_ "github.com/rackspace/gophercloud/openstack/compute/v2/servers"
	"github.com/rackspace/gophercloud/pagination"
	_ "github.com/rackspace/gophercloud/pagination"
)

type OpenStackPublicIPHandler struct {
	Client *gophercloud.ServiceClient
}

type PublicIPInfo struct {
	ID string `mapstructure:"id"`
	// FixedIP is the IP of the instance related to the Floating IP
	FixedIP string `mapstructure:"fixed_ip,omitempty"`
	// InstanceID is the ID of the instance that is using the Floating IP
	InstanceID string `mapstructure:"instance_id"`
	// IP is the actual Floating IP
	IP string `mapstructure:"ip"`
	// Pool is the pool of floating IPs that this floating IP belongs to
	Pool string `mapstructure:"pool"`
}

func (publicIPInfo *PublicIPInfo) setter(results floatingip.FloatingIP) PublicIPInfo {
	publicIPInfo.ID = results.ID
	publicIPInfo.FixedIP = results.FixedIP       //유동IP와 관련된 인스턴스 IP
	publicIPInfo.InstanceID = results.InstanceID //유동IP를 사용하는 인스턴스ID
	publicIPInfo.IP = results.IP                 //유동IP
	publicIPInfo.Pool = results.Pool

	return *publicIPInfo
}

func (publicIPInfo *PublicIPInfo) printInfo() {
	fmt.Println("ID : ", publicIPInfo.ID)
	fmt.Println("FixedIP : ", publicIPInfo.FixedIP)
	fmt.Println("InstanceID : ", publicIPInfo.InstanceID)
	fmt.Println("IP : ", publicIPInfo.IP)
	fmt.Println("Pool : ", publicIPInfo.Pool)
}

func (publicIPHandler *OpenStackPublicIPHandler) CreatePublicIP(publicIPReqInfo irs.PublicIPReqInfo) (irs.PublicIPInfo, error) {

	createOpts := floatingip.CreateOpts{
		"public1", //Pool 확인용
	}
	publicIP, err := floatingip.Create(publicIPHandler.Client, createOpts).Extract()

	if err != nil {
		panic(err)
	}

	spew.Dump(publicIP)
	return irs.PublicIPInfo{}, nil
}

func (publicIPHandler *OpenStackPublicIPHandler) ListVNetwork() ([]*irs.PublicIPInfo, error) {

	var IPList = make([]PublicIPInfo, 20)

	pager := floatingip.List(publicIPHandler.Client)

	err := pager.EachPage(func(page pagination.Page) (bool, error) {

		//Get server
		publicIPList, err := floatingip.ExtractFloatingIPs(page)
		if err != nil {
			return false, err
		}

		var publicInfo PublicIPInfo

		//add to list
		for _, publicIPs := range publicIPList {
			publicIP := publicInfo.setter(publicIPs)
			IPList = append(IPList, publicIP)
			spew.Dump(publicIP)
			fmt.Println("-----------------------------------")
		}
		return true, nil
	})

	if err != nil {
		panic(err)
	}
	return nil, nil
}

func (publicIPHandler *OpenStackPublicIPHandler) GetVNetwork(publicIPID string) (irs.PublicIPInfo, error) {
	client, err := config.GetServiceClient()
	if err != nil {
		panic(err)
	}

	publicIP, err := floatingip.Get(client, publicIPID).Extract()

	var info PublicIPInfo
	info.setter(*publicIP)
	spew.Dump(info)

	return irs.PublicIPInfo{}, nil
}

func (publicIPHandler *OpenStackPublicIPHandler) DeleteVNetwork(publicIPID string) (bool, error) {
	client, err := config.GetServiceClient()
	if err != nil {
		panic(err)
	}
	result := floatingip.Delete(client, publicIPID)
	fmt.Println("Delete : ", result)

	return false, nil
}
