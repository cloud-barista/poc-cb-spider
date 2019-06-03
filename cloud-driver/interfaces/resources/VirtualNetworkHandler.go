package resources

type VirtualNetworkHandler interface {
	//CreateVirtualNetwork() (VirtualNetworkInfo, error)
	CreateVirtualNetwork()
	RegistVirtualNetworkInfo()
	GetVirtualNetworkInfoList() 
	GetVirtualNetworkInfo() 
	RemoveVirtualNetworkInfo() 
	DeleteVirtualNetwork() 
}

