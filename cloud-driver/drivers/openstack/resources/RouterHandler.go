package resources

import (
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/davecgh/go-spew/spew"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/networking/v2/extensions/layer3/routers"
)

type OpenStackRouterHandler struct {
	Client *gophercloud.ServiceClient
}

func (routerHandler *OpenStackRouterHandler) CreateRouter(routerReqInfo irs.RouterReqInfo) (irs.RouterInfo, error) {
	
	// @TODO: Router 생성 요청 파라미터 정의 필요
	type RouterReqInfo struct {
		Name string
		GateWayId string
		AdminStateUp bool
	}
	
	reqInfo := RouterReqInfo{
		Name: routerReqInfo.Name,
		GateWayId: "b6610ceb-8089-48b0-9bfc-3c35e4e245cf",
		AdminStateUp: true,
	}
	
	createOpts := routers.CreateOpts{
		Name: reqInfo.Name,
		AdminStateUp: &reqInfo.AdminStateUp,
		GatewayInfo: &routers.GatewayInfo{
			NetworkID: reqInfo.GateWayId,
		},
	}
	
	// Create Router
	router, err := routers.Create(routerHandler.Client, createOpts).Extract()
	if err != nil {
		return irs.RouterInfo{}, err
	}
	
	spew.Dump(router)
	return irs.RouterInfo{Id: router.ID, Name: router.Name}, nil
}

func (routerHandler *OpenStackRouterHandler) ListRouter() ([]*irs.RouterInfo, error) {
	return nil, nil
}

func (routerHandler *OpenStackRouterHandler) GetRouter(routerID string) (irs.RouterInfo, error) {
	return irs.RouterInfo{}, nil
}

func (routerHandler *OpenStackRouterHandler) DeleteRouter(routerID string) (bool, error) {
	return false, nil
}

func (routerHandler *OpenStackRouterHandler) AddInterface(interfaceReqInfo irs.InterfaceReqInfo) (irs.InterfaceInfo, error) {
	
	createOpts := routers.InterfaceOpts{
		SubnetID: interfaceReqInfo.SubnetId,
	}
	
	// Add Interface
	ir, err := routers.AddInterface(routerHandler.Client, interfaceReqInfo.RouterId, createOpts).Extract()
	if err != nil {
		return irs.InterfaceInfo{}, err
	}
	
	spew.Dump(ir)
	return irs.InterfaceInfo{}, nil
}

func (routerHandler *OpenStackRouterHandler) DeleteInterface(interfaceReqInfo irs.InterfaceReqInfo) (irs.InterfaceInfo, error) {
	return irs.InterfaceInfo{}, nil
}
