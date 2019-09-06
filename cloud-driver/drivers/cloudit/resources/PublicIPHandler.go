package resources

import (
	"fmt"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client/dna/adaptiveip"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/davecgh/go-spew/spew"
	"strconv"
)

type ClouditPublicIPHandler struct {
	CredentialInfo idrv.CredentialInfo
	Client         *client.RestClient
}

func (publicIPHandler *ClouditPublicIPHandler) CreatePublicIP(publicIPReqInfo irs.PublicIPReqInfo) (irs.PublicIPInfo, error) {
	publicIPHandler.Client.TokenID = publicIPHandler.CredentialInfo.AuthToken
	authHeader := publicIPHandler.Client.AuthenticatedHeaders()

	type PublicIPReqInfo struct {
		IP         string `json:"ip" required:"true"`
		Name       string `json:"name" required:"true"`
		PrivateIP  string `json:"privateIp" required:"true"`
		Protection int    `json:"protection" required:"true"`
	}
	reqInfo := PublicIPReqInfo{
		IP:         "182.252.135.54",
		Name:       "test-dong1",
		PrivateIP:  "10.0.8.2",
		Protection: 0,
	}

	createOpts := client.RequestOpts{
		JSONBody:    reqInfo,
		MoreHeaders: authHeader,
	}

	publicIP, err := adaptiveip.Create(publicIPHandler.Client, &createOpts)
	if err != nil {
		panic(err)
	}
	spew.Dump(publicIP)

	return irs.PublicIPInfo{}, nil
}

func (publicIPHandler *ClouditPublicIPHandler) ListPublicIP() ([]*irs.PublicIPInfo, error) {
	publicIPHandler.Client.TokenID = publicIPHandler.CredentialInfo.AuthToken
	authHeader := publicIPHandler.Client.AuthenticatedHeaders()

	requestOpts := client.RequestOpts{
		MoreHeaders: authHeader,
	}

	publicIPList, err := adaptiveip.List(publicIPHandler.Client, &requestOpts)
	if err != nil {
		panic(err)
	}

	for i, publicIP := range *publicIPList {
		fmt.Println("[" + strconv.Itoa(i) + "]")
		spew.Dump(publicIP)
	}
	return nil, nil
}

func (publicIPHandler *ClouditPublicIPHandler) GetPublicIP(publicIPID string) (irs.PublicIPInfo, error) {
	return irs.PublicIPInfo{}, nil
}

func (publicIPHandler *ClouditPublicIPHandler) DeletePublicIP(publicIPID string) (bool, error) {
	publicIPHandler.Client.TokenID = publicIPHandler.CredentialInfo.AuthToken
	authHeader := publicIPHandler.Client.AuthenticatedHeaders()

	requestOpts := client.RequestOpts{
		MoreHeaders: authHeader,
	}

	err := adaptiveip.Delete(publicIPHandler.Client, publicIPID, &requestOpts)
	if err != nil {
		panic(err)
	}

	return true, nil
}
