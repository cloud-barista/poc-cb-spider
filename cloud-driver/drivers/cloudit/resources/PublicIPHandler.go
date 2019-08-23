package resources

import (
	"fmt"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client/dna/adaptiveip"
	"github.com/davecgh/go-spew/spew"
	"strconv"
)

type ClouditPublicIPHandler struct {
	Client *client.RestClient
}


func (publicIPHandler *ClouditPublicIPHandler) CreatePublicIP(publicIPReqInfo irs.PublicIPReqInfo) (irs.PublicIPInfo, error) {
	return irs.PublicIPInfo{}, nil
}

func (publicIPHandler *ClouditPublicIPHandler) ListPublicIP() ([]*irs.PublicIPInfo, error) {
	publicIPList, _ := adaptiveip.List(publicIPHandler.Client)
	for i, publicIP := range *publicIPList {
		fmt.Println("["+ strconv.Itoa(i) +"]")
		spew.Dump(publicIP)
	}
	return nil, nil
}

func (publicIPHandler *ClouditPublicIPHandler) GetPublicIP(publicIPID string) (irs.PublicIPInfo, error) {
	return irs.PublicIPInfo{}, nil
}

func (publicIPHandler *ClouditPublicIPHandler) DeletePublicIP(publicIPID string) (bool, error) {
	return true, nil
}
