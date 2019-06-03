package resources

type ImageHandler interface {
	//CreateImage() (ImageInfo, error)
	CreateImage()
	RegistImageInfo()
	GetImageInfoList()
	GetImageInfo() 
	RemoveImageInfo() 
	DeleteImage() 
}

