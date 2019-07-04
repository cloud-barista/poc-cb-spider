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
	"github.com/gophercloud/gophercloud/openstack/compute/v2/images"
	"github.com/gophercloud/gophercloud/pagination"

	//irs "github.com/hyokyungk/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	irs "poc-cb-spider2/cloud-driver/interfaces/resources"
)

func GetServiceClient() (*gophercloud.ServiceClient, error) {

	scop := gophercloud.AuthScope{
		ProjectID:"{ProjectID}",
	}
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: "{EndPoint}",
		Username: "{Username}",
		Password: "{Password}",
		DomainName: "{DomainName}",
		Scope:&scop,
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return nil, err
	}

	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})
	if err != nil {
		return nil, err
	}

	return client, err
}

type OpenStackImageHandler struct{}

func (OpenStackImageHandler) CreateImage(imageReqInfo irs.ImageReqInfo) (*irs.ImageInfo, error) {
	return nil, nil
}

func (OpenStackImageHandler) ListImage() ([]*irs.ImageInfo, error) {
	client, err := GetServiceClient()
	if err != nil {
		return nil, err
	}

	opts := images.ListOpts{}
	pager := images.ListDetail(client, opts)

	var imageArr []*irs.ImageInfo

	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		pagerList, err := images.ExtractImages(page)
		if err != nil {
			return false, nil
		}
		// Add to imageArr
		for _, i := range pagerList {
			var img = irs.ImageInfo{
				Id: i.ID,
				Name: i.Name,
			}
			imageArr = append(imageArr, &img)
		}
		return true, nil
	})

	return imageArr, nil
}

func (OpenStackImageHandler) GetImage(imageID string) (*irs.ImageInfo, error) {
	client, err := GetServiceClient()
	if err != nil {
		return nil, err
	}

	image, err := images.Get(client, imageID).Extract()
	if err != nil {
		return nil, err
	}

	result := irs.ImageInfo{
		Id: image.ID,
		Name: image.Name,
	}

	return &result, nil
}

func (OpenStackImageHandler) DeleteImage(imageID string) (bool, error) {
	return true, nil
}
