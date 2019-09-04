package resources

type RouterReqInfo struct {
	Id string
	Name string
	// @todo
}

type RouterInfo struct {
	Id string
	Name string
	// @todo
}

type InterfaceReqInfo struct {
	RouterId string
	SubnetId string
	// @todo
}

type InterfaceInfo struct {
	Id string
	Name string
	// @todo
}

type RouterHandler interface {
	CreateRouter(routerReqInfo RouterReqInfo) (RouterInfo, error)
	ListRouter() ([]*RouterInfo, error)
	GetRouter(routerID string) (RouterInfo, error)
	DeleteRouter(routerID string) (bool, error)
	AddInterface(interfaceReqInfo InterfaceReqInfo) (InterfaceInfo, error)
}
