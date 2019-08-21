package resources

import (
	"bytes"
	"fmt"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/davecgh/go-spew/spew"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/compute/v2/images"
	imgsvc "github.com/rackspace/gophercloud/openstack/imageservice/v2/images"
	"github.com/rackspace/gophercloud/pagination"
	"io/ioutil"
	"os"
)

type OpenStackImageHandler struct {
	Client      *gophercloud.ServiceClient
	ImageClient *gophercloud.ServiceClient
}

type ImageInfo struct {
	ID       string
	Created  string
	MinDisk  int
	MinRAM   int
	Name     string
	Progress int
	Status   string
	Updated  string
	Metadata map[string]string
}

func (imageInfo *ImageInfo) setter(results images.Image) *ImageInfo {
	imageInfo.ID = results.ID
	imageInfo.Created = results.Created
	imageInfo.MinDisk = results.MinDisk
	imageInfo.MinRAM = results.MinRAM
	imageInfo.Name = results.Name
	imageInfo.Progress = results.Progress
	imageInfo.Status = results.Status
	imageInfo.Updated = results.Updated
	imageInfo.Metadata = results.Metadata
	
	return imageInfo
}

func (imageHandler *OpenStackImageHandler) CreateImage(imageReqInfo irs.ImageReqInfo) (irs.ImageInfo, error) {

	// @TODO: Image 생성 요청 파라미터 정의 필요
	type ImageReqInfo struct {
		Name            string
		ContainerFormat string
		DiskFormat      string
	}
	reqInfo := ImageReqInfo{
		Name:            imageReqInfo.Name,
		ContainerFormat: "bare",
		DiskFormat:      "iso",
	}

	createOpts := imgsvc.CreateOpts{
		Name:            reqInfo.Name,
		ContainerFormat: reqInfo.ContainerFormat,
		DiskFormat:      reqInfo.DiskFormat,
	}

	// Create Image
	image, err := imgsvc.Create(imageHandler.ImageClient, createOpts).Extract()
	if err != nil {
		return irs.ImageInfo{}, err
	}
	spew.Dump(image)

	// Update Image File
	rootPath := os.Getenv("CBSPIDER_PATH")
	imageBytes, err := ioutil.ReadFile(rootPath + "/image/coreos_production_iso_image.iso")
	if err != nil {
		return irs.ImageInfo{}, err
	}
	result := imgsvc.Upload(imageHandler.ImageClient, image.ID, bytes.NewReader(imageBytes))
	if result.Err != nil {
		return irs.ImageInfo{}, err
	}
	fmt.Println(result)
	
	imageInfo := irs.ImageInfo{
		Id: image.ID,
		Name: image.Name,
	}
	return imageInfo, nil
}

func (imageHandler *OpenStackImageHandler) ListImage() ([]*irs.ImageInfo, error) {
	var imageList []*ImageInfo

	pager := images.ListDetail(imageHandler.Client, images.ListOpts{})
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		// Get Image
		list, err := images.ExtractImages(page)
		if err != nil {
			return false, err
		}
		// Add to List
		for _, img := range list {
			imageInfo := new(ImageInfo).setter(img)
			imageList = append(imageList, imageInfo)
		}
		return true, nil
	})
	if err != nil {
		return nil, err
	}

	spew.Dump(imageList)
	return nil, nil
}

func (imageHandler *OpenStackImageHandler) GetImage(imageID string) (irs.ImageInfo, error) {
	image, err := images.Get(imageHandler.Client, imageID).Extract()
	if err != nil {
		return irs.ImageInfo{}, err
	}

	imageInfo := new(ImageInfo).setter(*image)

	spew.Dump(imageInfo)
	return irs.ImageInfo{}, nil
}

func (imageHandler *OpenStackImageHandler) DeleteImage(imageID string) (bool, error) {
	err := images.Delete(imageHandler.Client, imageID).ExtractErr()
	if err != nil {
		return false, err
	}
	return true, nil
}
