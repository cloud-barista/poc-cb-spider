package resources

type ConnectionKeyHandler interface {
	//CreateConnectionKey() (ConnectionKey, error)
	CreateConnectionKey()
	RegistConnectionKeyInfo()
	GetConnectionKeyInfoList()
	GetConnectionKeyInfo()
	RemoveConnectionKeyInfo()
	DeleteConnectionKey() 
}

