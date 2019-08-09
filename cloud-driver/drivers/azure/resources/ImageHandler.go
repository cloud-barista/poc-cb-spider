package resources

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-06-01/compute"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/davecgh/go-spew/spew"
	"strings"
)

type AzureImageHandler struct {
	Region idrv.RegionInfo
	Ctx    context.Context
	Client *compute.ImagesClient
}

// @TODO: ImageInfo 리소스 프로퍼티 정의 필요
type ImageInfo struct {
	ID            string
	Name          string
	Location      string
	OsType        string
	OsDiskSize    int32
	OsState       string
	ManagedDiskId string
}

func (imageInfo *ImageInfo) setter(image compute.Image) *ImageInfo {
	imageInfo.ID = *image.ID
	imageInfo.Name = *image.Name
	imageInfo.Location = *image.Location
	imageInfo.OsType = fmt.Sprint(image.ImageProperties.StorageProfile.OsDisk.OsType)
	imageInfo.OsDiskSize = *image.StorageProfile.OsDisk.DiskSizeGB
	imageInfo.OsState = fmt.Sprint(image.StorageProfile.OsDisk.OsState)

	if image.StorageProfile.OsDisk.ManagedDisk != nil {
		imageInfo.ManagedDiskId = *image.StorageProfile.OsDisk.ManagedDisk.ID
	}

	return imageInfo
}

func (imageHandler *AzureImageHandler) CreateImage(irs.ImageReqInfo) (irs.ImageInfo, error) {
	return irs.ImageInfo{}, nil
}

func (imageHandler *AzureImageHandler) ListImage() ([]*irs.ImageInfo, error) {
	resultList, err := imageHandler.Client.List(imageHandler.Ctx)
	if err != nil {
		panic(err)
	}

	var imageList []*ImageInfo
	for _, image := range resultList.Values() {

		imageInfo := new(ImageInfo).setter(image)
		imageList = append(imageList, imageInfo)
	}

	spew.Dump(imageList)
	return nil, nil
}

func (imageHandler *AzureImageHandler) GetImage(imageID string) (irs.ImageInfo, error) {
	imageIdArr := strings.Split(imageID, ":")

	image, err := imageHandler.Client.Get(imageHandler.Ctx, imageIdArr[0], imageIdArr[1], "")
	if err != nil {
		panic(err)
	}

	imageInfo := new(ImageInfo).setter(image)

	spew.Dump(imageInfo)
	return irs.ImageInfo{}, nil
}

func (imageHandler *AzureImageHandler) DeleteImage(imageID string) (bool, error) {
	imageIdArr := strings.Split(imageID, ":")

	future, err := imageHandler.Client.Delete(imageHandler.Ctx, imageIdArr[0], imageIdArr[1])
	if err != nil {
		return false, err
	}
	err = future.WaitForCompletionRef(imageHandler.Ctx, imageHandler.Client.Client)
	if err != nil {
		return false, err
	}
	return true, nil
}
