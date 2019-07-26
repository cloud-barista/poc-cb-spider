package resources

import (
	_ "fmt"
	_ "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/config"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	_ "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/keypairs"
	_ "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/startstop"
	_ "github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	_ "github.com/gophercloud/gophercloud/pagination"
)

type OpenStackImageHandler struct{}

func (OpenStackImageHandler) CreateImage(imageReqInfo irs.ImageReqInfo) (irs.ImageInfo, error) {

	return irs.ImageInfo{}, nil
}

//ListImage() ([]*ImageInfo, error)
//GetImage(imageID string) (ImageInfo, error)
//DeleteImage(imageID string) (bool, error)
