package resources

import (
	"fmt"
	//"github.com/Azure/azure-sdk-for-go/profiles/latest/recoveryservices/mgmt/backup"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client/dna/subnet"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/davecgh/go-spew/spew"
	//"github.com/rackspace/gophercloud/openstack/networking/v2/networks"
	//"github.com/rackspace/gophercloud/openstack/networking/v2/subnets"
	//"io"
	"strconv"
)

type ClouditVNetworkHandler struct {
	CredentialInfo idrv.CredentialInfo
	Client         *client.RestClient
}

func (vNetworkHandler *ClouditVNetworkHandler) CreateVNetwork(vNetReqInfo irs.VNetworkReqInfo) (irs.VNetworkInfo, error) {
	var authHeader map[string]string
	authHeader = make(map[string]string)
	authHeader["X-Auth-Token"] = vNetworkHandler.CredentialInfo.AuthToken

	// @TODO: Subnet 생성 요청 파라미터 정의 필요
	type VNetworkReqInfo struct {
		name       string
		protection int
	}
	reqInfo := VNetworkReqInfo{
		name:       "test-Dong222",
		protection: 0,
	}

	createOpts := client.RequestOpts{
		JSONBody: reqInfo,
		//RawBody: reqInfo,
		//JSONResponse: nil,
		//OkCodes:      nil,
		MoreHeaders: authHeader,
	}
	subnet, err := subnet.Create(vNetworkHandler.Client, &createOpts)
	if err != nil {
		return irs.VNetworkInfo{}, err
	}

	spew.Dump(subnet)
	return irs.VNetworkInfo{}, nil
}

func (vNetworkHandler *ClouditVNetworkHandler) ListVNetwork() ([]*irs.VNetworkInfo, error) {
	var authHeader map[string]string
	authHeader = make(map[string]string)
	authHeader["X-Auth-Token"] = vNetworkHandler.CredentialInfo.AuthToken
	requestOpts := client.RequestOpts{
		//JSONBody:     nil,
		//RawBody:      nil,
		//JSONResponse: nil,
		//OkCodes:      nil,
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
	var authHeader map[string]string
	authHeader = make(map[string]string)
	authHeader["X-Auth-Token"] = vNetworkHandler.CredentialInfo.AuthToken
	requestOpts := client.RequestOpts{
		//JSONBody:     nil,
		//RawBody:      nil,
		//JSONResponse: nil,
		//OkCodes:      nil,
		MoreHeaders: authHeader,
	}

	vNet, _ := subnet.Delete(vNetworkHandler.Client, vNetworkID, &requestOpts)
	spew.Dump(vNet)

	return false, nil
}
