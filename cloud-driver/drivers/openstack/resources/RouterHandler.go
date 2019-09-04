package resources

import (
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/davecgh/go-spew/spew"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/networking/v2/extensions/layer3/routers"
	"github.com/rackspace/gophercloud/pagination"
)

type OpenStackRouterHandler struct {
	Client *gophercloud.ServiceClient
}

// @TODO: Router 리소스 프로퍼티 정의 필요
type RouterInfo struct {
	Id string
	Name string
	//GateWayId string
	//AdminStateUp bool
}

func (routerInfo *RouterInfo) setter(router routers.Router) *RouterInfo {
	return routerInfo
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
	var routerInfoList []*RouterInfo
	
	pager := routers.List(routerHandler.Client, routers.ListOpts{})
	err := pager.EachPage(func(page pagination.Page) (b bool, e error) {
		// Get Router
		list, err := routers.ExtractRouters(page)
		if err != nil {
			return false, err
		}
		for _, r := range list {
			routerInfo := new(RouterInfo).setter(r)
			routerInfoList = append(routerInfoList, routerInfo)
		}
		return true, nil
	})
	if err != nil {
		return  nil, err
	}
	
	spew.Dump(routerInfoList)
	return nil, nil
}

func (routerHandler *OpenStackRouterHandler) GetRouter(routerID string) (irs.RouterInfo, error) {
	router, err := routers.Get(routerHandler.Client, routerID).Extract()
	if err != nil {
		return irs.RouterInfo{}, err
	}
	
	routerInfo := new(RouterInfo).setter(*router)
	
	spew.Dump(routerInfo)
	return irs.RouterInfo{}, nil
}

func (routerHandler *OpenStackRouterHandler) DeleteRouter(routerID string) (bool, error) {
	err := routers.Delete(routerHandler.Client, routerID).ExtractErr()
	if err != nil {
		return false, err
	}
	return true, nil
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

func (routerHandler *OpenStackRouterHandler) DeleteInterface(routerID string, subnetID string) (irs.InterfaceInfo, error) {
	deleteOpts := routers.InterfaceOpts{
		SubnetID: subnetID,
	}
	
	// Delete Interface
	ir, err := routers.RemoveInterface(routerHandler.Client, routerID, deleteOpts).Extract()
	if err != nil {
		return irs.InterfaceInfo{}, err
	}
	
	spew.Dump(ir)
	return irs.InterfaceInfo{}, nil
}
