// Cloud Driver Interface of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is Resouces interfaces of Cloud Driver.
//
// by powerkim@etri.re.kr, 2019.06.


package resources

import (
	irs "github.com/hyokyungk/poc-cb-spider/cloud-driver/interfaces/resources"
)

type OpenStackImageHandler struct{}

func (OpenStackImageHandler) CreateImage(imageReqInfo irs.ImageReqInfo) (*irs.ImageInfo, error) {
	return nil, nil
}

func (OpenStackImageHandler) ListImage() ([]*irs.ImageInfo, error) {
	//osConn := OpenStackCloudConnection {}
	return nil, nil
}

func (OpenStackImageHandler) GetImage(imageID string) (*irs.ImageInfo, error) {
	return nil, nil
}

func (OpenStackImageHandler) DeleteImage(imageID string) (bool, error) {
	return true, nil
}
