package resources

import (
	"fmt"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client/ace/image"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/davecgh/go-spew/spew"
	"strconv"
)

type ClouditImageHandler struct {
	CredentialInfo idrv.CredentialInfo
	Client         *client.RestClient
}

func (imageHandler *ClouditImageHandler) CreateImage(imageReqInfo irs.ImageReqInfo) (irs.ImageInfo, error) {
	imageHandler.Client.TokenID = imageHandler.CredentialInfo.AuthToken
	authHeader := imageHandler.Client.AuthenticatedHeaders()

	type ImageReqInfo struct {
		Name         string `json:"name" required:"true"`
		VolumeId     string `json:"volumeId" required:"true"`
		Ownership    string `json:"ownership" required:"true"`
		Protection   int    `json:"protection" required:"true"`
		SnapshotId   string `json:"snapshotId" required:"true"`
		PoolId       string `json:"poolId" required:"true"`
		Size         int    `json:"size" required:"true"`
		TemplateType string `json:"templateType" required:"true"`
		Format       string `json:"format" required:"true"`
		SourceType   string `json:"sourceType" required:"true"`
	}

	reqInfo := ImageReqInfo{
		Name:         "Test-Dong1",
		VolumeId:     "7f7a87f0-3acb-4313-90f7-22eb65a6d33f",
		Ownership:    "TENANT",
		Protection:   0,
		SnapshotId:   "",
		PoolId:       "100.100.100.100-cloudit-vm4",
		Size:         100,
		TemplateType: "DEFAULT",
		Format:       "qcow2",
		SourceType:   "server",
	}

	createOpts := client.RequestOpts{
		JSONBody:    reqInfo,
		MoreHeaders: authHeader,
	}

	image, err := image.Create(imageHandler.Client, &createOpts)
	if err != nil {
		panic(err)
	}

	spew.Dump(image)

	return irs.ImageInfo{}, nil
}

func (imageHandler *ClouditImageHandler) ListImage() ([]*irs.ImageInfo, error) {
	imageHandler.Client.TokenID = imageHandler.CredentialInfo.AuthToken
	authHeader := imageHandler.Client.AuthenticatedHeaders()

	requestOpts := client.RequestOpts{
		MoreHeaders: authHeader,
	}

	imageList, err := image.List(imageHandler.Client, &requestOpts)
	if err != nil {
		panic(err)
	}

	for i, image := range *imageList {
		fmt.Println("[" + strconv.Itoa(i) + "]")
		spew.Dump(image)
	}

	return nil, nil
}

func (imageHandler *ClouditImageHandler) GetImage(imageID string) (irs.ImageInfo, error) {

	return irs.ImageInfo{}, nil
}

func (imageHandler *ClouditImageHandler) DeleteImage(imageID string) (bool, error) {
	imageHandler.Client.TokenID = imageHandler.CredentialInfo.AuthToken
	authHeader := imageHandler.Client.AuthenticatedHeaders()

	requestOpts := client.RequestOpts{
		MoreHeaders: authHeader,
	}

	err := image.Delete(imageHandler.Client, imageID, &requestOpts)
	if err != nil {
		panic(err)
	}

	return true, nil
}
