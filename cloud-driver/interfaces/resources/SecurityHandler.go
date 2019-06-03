package resources

type SecurityHandler interface {
	//CreateSecurity() (SecurityInfo, error)
	CreateSecurity()
	RegistSecurityInfo()
	GetSecurityInfoList()
	GetSecurityInfo()
	RemoveSecurityInfo() 
	DeleteSecurity() 
}

