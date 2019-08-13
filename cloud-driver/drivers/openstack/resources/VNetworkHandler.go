package resources

import (
	"fmt"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/davecgh/go-spew/spew"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/networking/v2/networks"
	"github.com/rackspace/gophercloud/pagination"
)

type OpenStackVNetworkHandler struct {
	Client *gophercloud.ServiceClient
}

type VNetworkInfo struct {
	ID           string
	Name         string
	AdminStateUp bool
	Status       string
	Subnets      []string
	TenantID     string
	Shared       bool
}

func (vNetworkInfo *VNetworkInfo) printInfo() {
	fmt.Println("ID : ", vNetworkInfo.ID)
	fmt.Println("Name : ", vNetworkInfo.Name)
	fmt.Println("AdminStateUp : ", vNetworkInfo.AdminStateUp)
	fmt.Println("Status : ", vNetworkInfo.Status)
	fmt.Println("Subnets : ", vNetworkInfo.Subnets)
	fmt.Println("TenantID : ", vNetworkInfo.TenantID)
	fmt.Println("Shared : ", vNetworkInfo.Shared)
}

func (vNetworkInfo *VNetworkInfo) setter(network networks.Network) VNetworkInfo {
	vNetworkInfo.ID = network.ID
	vNetworkInfo.Name = network.Name
	vNetworkInfo.AdminStateUp = network.AdminStateUp
	vNetworkInfo.Status = network.Status
	vNetworkInfo.Subnets = network.Subnets
	vNetworkInfo.TenantID = network.TenantID
	vNetworkInfo.Shared = network.Shared

	return *vNetworkInfo
}

func (vNetworkHandler *OpenStackVNetworkHandler) CreateVNetwork(vNetworkReqInfo irs.VNetworkReqInfo) (irs.VNetworkInfo, error) {

	opts := networks.CreateOpts{Name: "testing Network", AdminStateUp: networks.Up} //이름 변경시 추가 매개변수 필요...

	network, err := networks.Create(vNetworkHandler.Client, opts).Extract()
	if err != nil {
		panic(err)
	}
	fmt.Println("Created : ", network)

	return irs.VNetworkInfo{}, nil
}

func (vNetworkHandler *OpenStackVNetworkHandler) ListVNetwork() ([]*irs.VNetworkInfo, error) {
	var vNetworkInfoSlices = make([]VNetworkInfo, 20)

	pager := networks.List(vNetworkHandler.Client, nil)
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		networkList, err := networks.ExtractNetworks(page)
		if err != nil {
			return false, err
		}

		var vNetworkInfo VNetworkInfo
		for _, network := range networkList {
			vNetwork := vNetworkInfo.setter(network)
			vNetworkInfoSlices = append(vNetworkInfoSlices, vNetwork)
			spew.Dump(network)
			fmt.Println("###############################")
		}

		return true, nil
	})

	if err != nil {
		panic(err)
	}

	return nil, nil
}

func (vNetworkHandler *OpenStackVNetworkHandler) GetVNetwork(vNetworkID string) (irs.VNetworkInfo, error) {
	network, err := networks.Get(vNetworkHandler.Client, vNetworkID).Extract()
	if err != nil {
		panic(err)
	}

	var vNetworkInfo VNetworkInfo
	vNetworkInfo.setter(*network)
	vNetworkInfo.printInfo()

	return irs.VNetworkInfo{}, nil
}

func (vNetworkHandler *OpenStackVNetworkHandler) DeleteVNetwork(vNetworkID string) (bool, error) {
	result := networks.Delete(vNetworkHandler.Client, vNetworkID)
	fmt.Println("Deleted : ", result)

	return false, nil
}
