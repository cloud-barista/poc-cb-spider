// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is Resouces interfaces of Cloud Driver.
//
// by powerkim@etri.re.kr, 2019.06.


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

