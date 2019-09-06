package resources

import (
	"fmt"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client/dna/subnet"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/davecgh/go-spew/spew"
	"strconv"
)

type ClouditVNetworkHandler struct {
	CredentialInfo idrv.CredentialInfo
	Client         *client.RestClient
}

func (vNetworkHandler *ClouditVNetworkHandler) CreateVNetwork(vNetReqInfo irs.VNetworkReqInfo) (irs.VNetworkInfo, error) {
	vNetworkHandler.Client.TokenID = vNetworkHandler.CredentialInfo.AuthToken
	authHeader := vNetworkHandler.Client.AuthenticatedHeaders()

	// @TODO: Subnet 생성 요청 파라미터 정의 필요
	type VNetworkReqInfo struct {
		Name       string `json:"name" required:"true"`
		Protection int    `json:"protection" required:"true"`
		Prefix     string `json:"prefix" required:"true"`
		Gateway    string `json:"gateway" required:"true"`
		Addr       string `json:"addr" required:"true"`
	}
	reqInfo := VNetworkReqInfo{
		Name:       "test-Dong1",
		Protection: 0,
		Prefix:     "22",
		Gateway:    "10.0.12.1",
		Addr:       "10.0.12.0",
	}

	createOpts := client.RequestOpts{
		JSONBody:    reqInfo,
		MoreHeaders: authHeader,
	}

	subnet, err := subnet.Create(vNetworkHandler.Client, &createOpts)
	if err != nil {
		panic(err)
	}

	spew.Dump(subnet)

	return irs.VNetworkInfo{}, nil
}

func (vNetworkHandler *ClouditVNetworkHandler) ListVNetwork() ([]*irs.VNetworkInfo, error) {
	vNetworkHandler.Client.TokenID = vNetworkHandler.CredentialInfo.AuthToken
	authHeader := vNetworkHandler.Client.AuthenticatedHeaders()

	requestOpts := client.RequestOpts{
		MoreHeaders: authHeader,
	}

	vNetList, err := subnet.List(vNetworkHandler.Client, &requestOpts)
	if err != nil {
		panic(err)
	}

	for i, vNet := range *vNetList {
		fmt.Println("[" + strconv.Itoa(i) + "]")
		spew.Dump(vNet)
	}
	return nil, nil
}

func (vNetworkHandler *ClouditVNetworkHandler) GetVNetwork(vNetworkID string) (irs.VNetworkInfo, error) {
	return irs.VNetworkInfo{}, nil
}

func (vNetworkHandler *ClouditVNetworkHandler) DeleteVNetwork(vNetworkID string) (bool, error) {
	vNetworkHandler.Client.TokenID = vNetworkHandler.CredentialInfo.AuthToken
	authHeader := vNetworkHandler.Client.AuthenticatedHeaders()

	requestOpts := client.RequestOpts{
		MoreHeaders: authHeader,
	}

	err := subnet.Delete(vNetworkHandler.Client, vNetworkID, &requestOpts)
	if err != nil {
		panic(err)
	}

	return true, nil
}
