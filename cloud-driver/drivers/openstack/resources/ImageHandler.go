package resources

import (
	"fmt"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/config"
	_ "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/config"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	_ "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/keypairs"
	_ "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/startstop"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/images"
	_ "github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/pagination"
	_ "github.com/gophercloud/gophercloud/pagination"
)

type OpenStackImageHandler struct{}

type ImageInfo struct {
	ID                 string
	Created            string
	MinDisk            int
	MinRAM             int
	Name               string
	Progress           int
	Status             string
	Updated            string
	base_image_ref     interface{}
	boot_roles         interface{}
	description        interface{}
	image_location     interface{}
	image_state        interface{}
	image_type         interface{}
	instance_uuid      interface{}
	owner_id           interface{}
	owner_project_name interface{}
	owner_user_name    interface{}
	user_id            interface{}
}

func (imageInfo *ImageInfo) setter(results images.Image) ImageInfo {
	imageInfo.ID = results.ID
	imageInfo.Created = results.Created
	imageInfo.MinDisk = results.MinDisk
	imageInfo.MinRAM = results.MinRAM
	imageInfo.Name = results.Name
	imageInfo.Progress = results.Progress
	imageInfo.Status = results.Status
	imageInfo.Updated = results.Updated
	imageInfo.base_image_ref = results.Metadata["base_image_ref"]
	imageInfo.boot_roles = results.Metadata["boot_roles"]
	imageInfo.description = results.Metadata["description"]
	imageInfo.image_location = results.Metadata["image_location"]
	imageInfo.image_state = results.Metadata["image_state"]
	imageInfo.image_type = results.Metadata["image_type"]
	imageInfo.instance_uuid = results.Metadata["instance_uuid"]
	imageInfo.owner_id = results.Metadata["owner_id"]
	imageInfo.owner_project_name = results.Metadata["owner_project_name"]
	imageInfo.user_id = results.Metadata["user_id"]
	//Metadata에 어떤 종류 있는지 더 확인해서 추가

	return *imageInfo
}

func (imageInfo ImageInfo) printInfo() {
	fmt.Println("Name : ", imageInfo.Name)
	fmt.Println("Created : ", imageInfo.Created)
	fmt.Println("ID : ", imageInfo.ID)
	fmt.Println("MinDisk : ", imageInfo.MinDisk)
	fmt.Println("MinRAM : ", imageInfo.MinRAM)
	fmt.Println("Progress : ", imageInfo.Progress)
	fmt.Println("Status : ", imageInfo.Status)
	fmt.Println("Updated : ", imageInfo.Updated)
	fmt.Println("base_image_ref : ", imageInfo.base_image_ref)
	fmt.Println("boot_roles : ", imageInfo.boot_roles)
	fmt.Println("description : ", imageInfo.description)
	fmt.Println("image_location : ", imageInfo.image_location)
	fmt.Println("image_state : ", imageInfo.image_state)
	fmt.Println("image_type : ", imageInfo.image_type)
	fmt.Println("instance_uuid : ", imageInfo.instance_uuid)
	fmt.Println("owner_id : ", imageInfo.owner_id)
	fmt.Println("owner_project_name : ", imageInfo.owner_project_name)
	fmt.Println("owner_user_name : ", imageInfo.owner_user_name)
	fmt.Println("user_id : ", imageInfo.user_id)
}

func (openStackImageHandler OpenStackImageHandler) CreateImage(imageReqInfo irs.ImageReqInfo) (irs.ImageInfo, error) {
	fmt.Println("Call CreateImage()")
	return irs.ImageInfo{}, nil
}

func (openStackImageHandler OpenStackImageHandler) ListImage() ([]*irs.ImageInfo, error) {

	client, err := config.GetServiceClient()
	if err != nil {
		panic(err)
	}
	var infoSlices = make([]ImageInfo, 20)
	pager := images.ListDetail(client, images.ListOpts{})
	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		// Get Servers
		imageList, err := images.ExtractImages(page)
		if err != nil {
			return false, err
		}
		var imageInfo ImageInfo
		// Add to List
		for _, images := range imageList {
			image := imageInfo.setter(images)
			infoSlices = append(infoSlices, image)
			imageInfo.printInfo()
			fmt.Println("##############################")
		}
		return true, nil
	})
	/*opts := serviceImages.ListOpts{
	     Owner: "a7509e1ae65945fda83f3e52c6296017",
	  }
	  allPages, err := serviceImages.List(client, opts).AllPages()
	  if err != nil {
	     panic(err)
	  }
	  allImages, err := serviceImages.ExtractImages(allPages)
	  if err != nil {
	     panic(err)
	  }*/
	/*for _, image := range allImages {
	   //fmt.Printf("%+v\n", image)
	   var imageInfo ImageInfo
	   imageInfo.setter(image)
	   //imageInfo.printInfo()
	   infoSlices = append(infoSlices, imageInfo)
	   //fmt.Println(image.Name)
	}*/
	return nil, nil
}
func (openStackImageHandler OpenStackImageHandler) GetImage(imageID string) (irs.ImageInfo, error) {
	client, err := config.GetServiceClient()
	if err != nil {
		panic(err)
	}
	//image, err := images.IDFromName(client, imageID)
	image, err := images.Get(client, "39b30620-dd05-4168-b5a9-e18b2d5ba8d1").Extract()

	var info ImageInfo
	info.setter(*image)
	info.printInfo()
	/*imageId, err := images.IDFromName(client, "hhhh")
	  if err != nil {
	  }
	  fmt.Println(">>" , imageId)*/
	return irs.ImageInfo{}, nil
}

func (openStackImageHandler OpenStackImageHandler) DeleteImage(imageID string) (bool, error) {
	client, err := config.GetServiceClient()
	if err != nil {
		panic(err)
	}
	result := images.Delete(client, "8b0bea4d-23e2-4457-b549-7adbbd5e276b")
	fmt.Println("Delete : ", result)
	return false, nil
}

//ListImage() ([]*ImageInfo, error)
//GetImage(imageID string) (ImageInfo, error)
//DeleteImage(imageID string) (bool, error)
