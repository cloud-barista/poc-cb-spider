package resources

import (
	"fmt"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/config"
	"github.com/gophercloud/gophercloud"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/keypairs"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/startstop"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/pagination"
)

// modified by powerkim, 2019.07.29
type OpenStackVMHandler struct{
	Client  *gophercloud.ServiceClient
}


// modified by powerkim, 2019.07.29
func (vmHandler *OpenStackVMHandler) StartVM(vmReqInfo irs.VMReqInfo) (irs.VMInfo, error) {
	/* client, err := config.GetServiceClient()
	if err != nil {
		panic(err)
	}
	*/

	// Add Server Create Options
	serverCreateOpts := servers.CreateOpts{
		Name:      vmReqInfo.Name,
		ImageRef:  vmReqInfo.ImageInfo.Id,
		FlavorRef: vmReqInfo.SpecID,
		Networks: []servers.Network{
			{UUID: vmReqInfo.VNetworkInfo.Id},
		},
		SecurityGroups: []string{
			vmReqInfo.SecurityInfo.Name,
		},
		ServiceClient: vmHandler.Client,
	}

	// Add KeyPair
	createOpts := keypairs.CreateOptsExt{
		CreateOptsBuilder: serverCreateOpts,
		KeyName:           vmReqInfo.KeyPairInfo.Name,
	}

	server, err := servers.Create(vmHandler.Client, createOpts).Extract()
	if err != nil {
		return irs.VMInfo{}, err
	}

	vmInfo := MappingServerInfo(*server)
	return vmInfo, nil
}

func (OpenStackVMHandler) SuspendVM(vmID string) {
	client, err := config.GetServiceClient()
	if err != nil {
		panic(err)
	}

	err = startstop.Stop(client, vmID).Err
}

func (OpenStackVMHandler) ResumeVM(vmID string) {
	client, err := config.GetServiceClient()
	if err != nil {
		panic(err)
	}

	err = startstop.Start(client, vmID).Err
}

func (OpenStackVMHandler) RebootVM(vmID string) {
	client, err := config.GetServiceClient()
	if err != nil {
		panic(err)
	}

	rebootOpts := servers.RebootOpts{
		Type: servers.SoftReboot,
		//Type: servers.HardReboot,
	}
	err = servers.Reboot(client, vmID, rebootOpts).ExtractErr()
}

func (OpenStackVMHandler) TerminateVM(vmID string) {
	client, err := config.GetServiceClient()
	if err != nil {
		panic(err)
	}

	err = servers.Delete(client, vmID).ExtractErr()
}

func (OpenStackVMHandler) ListVMStatus() []*irs.VMStatus {
	client, err := config.GetServiceClient()
	if err != nil {
		return nil
	}

	var vmStatusList []*irs.VMStatus

	pager := servers.List(client, nil)
	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		// Get VM Status
		list, err := servers.ExtractServers(page)
		if err != nil {
			return false, err
		}
		// Add to List
		for _, s := range list {
			vmStatus := irs.VMStatus(s.Status)
			vmStatusList = append(vmStatusList, &vmStatus)
		}
		return true, nil
	})

	return vmStatusList
}

func (OpenStackVMHandler) GetVMStatus(vmID string) irs.VMStatus {
	client, err := config.GetServiceClient()
	if err != nil {
		return irs.VMStatus("")
	}

	serverResult, err := servers.Get(client, vmID).Extract()
	if err != nil {
		return irs.VMStatus("")
	}

	return irs.VMStatus(serverResult.Status)
}

func (OpenStackVMHandler) ListVM() []*irs.VMInfo {
	client, err := config.GetServiceClient()
	if err != nil {
		return nil
	}

	var vmList []*irs.VMInfo

	pager := servers.List(client, nil)
	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		// Get Servers
		list, err := servers.ExtractServers(page)
		if err != nil {
			return false, err
		}
		// Add to List
		for _, s := range list {
			vmInfo := MappingServerInfo(s)
			vmList = append(vmList, &vmInfo)
		}

		return true, nil
	})

	return vmList
}

func (OpenStackVMHandler) GetVM(vmID string) irs.VMInfo {
	client, err := config.GetServiceClient()
	if err != nil {
		return irs.VMInfo{}
	}

	serverResult, err := servers.Get(client, vmID).Extract()
	if err != nil {
		fmt.Println(err)
		return irs.VMInfo{}
	}

	vmInfo := MappingServerInfo(*serverResult)

	return vmInfo
}

func MappingServerInfo(server servers.Server) irs.VMInfo {

	// Get Default VM Info
	vmInfo := irs.VMInfo{
		Name:      server.Name,
		Id:        server.ID,
		StartTime: server.Updated,
		//ImageID:   server.Image["id"].(string),
		//SpecID:    server.Flavor["id"].(string),
	}

	if len(server.Image) != 0 {
		vmInfo.ImageID = server.Image["id"].(string)
	}
	if len(server.Flavor) != 0 {
		vmInfo.SpecID = server.Flavor["id"].(string)
	}

	// Get VM Subnet, Address Info
	for k, subnet := range server.Addresses {
		vmInfo.SubNetworkID = k
		for _, addr := range subnet.([]interface{}) {
			addrMap := addr.(map[string]interface{})
			if addrMap["OS-EXT-IPS:type"] == "floating" {
				vmInfo.PrivateIP = addrMap["addr"].(string)
			} else if addrMap["OS-EXT-IPS:type"] == "fixed" {
				vmInfo.PublicIP = addrMap["addr"].(string)
			}
		}
	}

	return vmInfo
}
