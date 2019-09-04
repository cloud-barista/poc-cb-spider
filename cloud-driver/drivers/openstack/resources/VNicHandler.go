package resources

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/networking/v2/ports"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/rackspace/gophercloud/pagination"
)

type OpenStackVNicworkHandler struct {
	Client *gophercloud.ServiceClient
}

// @TODO: KeyPairInfo 리소스 프로퍼티 정의 필요
type PortInfo struct {
	Id string
	Name string
}

func (portInfo *PortInfo) setter(port ports.Port) *PortInfo {
	portInfo.Id = port.ID
	port.Name = port.Name
	return portInfo
}

func (vNicHandler *OpenStackVNicworkHandler) CreateVNic(vNicReqInfo irs.VNicReqInfo) (irs.VNicInfo, error) {
	
	return irs.VNicInfo{}, nil
}

func (vNicHandler *OpenStackVNicworkHandler) ListVNic() ([]*irs.VNicInfo, error) {
	var portList []PortInfo
	
	pager := ports.List(vNicHandler.Client, nil)
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		// Get Port
		list, err := ports.ExtractPorts(page)
		if err != nil {
			return false, err
		}
		// Add to Port
		for _, p := range list {
			PortInfo := new(PortInfo).setter(p)
			portList = append(portList, *PortInfo)
		}
		return true, nil
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (vNicHandler *OpenStackVNicworkHandler) GetVNic(vNicID string) (irs.VNicInfo, error) {
	port, err := ports.Get(vNicHandler.Client, vNicID).Extract()
	if err != nil {
		return irs.VNicInfo{}, err
	}
	
	portInfo := new(PortInfo).setter(*port)
	
	spew.Dump(portInfo)
	return irs.VNicInfo{}, nil
}

func (vNicHandler *OpenStackVNicworkHandler) DeleteVNic(vNicID string) (bool, error) {
	err := ports.Delete(vNicHandler.Client, vNicID).ExtractErr()
	if err != nil {
		return false, err
	}
	return true, nil
}
