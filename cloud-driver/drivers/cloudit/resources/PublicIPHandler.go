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
		IP:         "182.252.135.59",
		Name:       "test-dong1",
		PrivateIP:  "10.0.0.2",
		Protection: 0,
	}

	createOpts := client.RequestOpts{
		JSONBody:    reqInfo,
		MoreHeaders: authHeader,
	}

	publicIP, err := adaptiveip.Create(publicIPHandler.Client, &createOpts)
	if err != nil {
		return irs.PublicIPInfo{}, err
	}

	spew.Dump(publicIP)
	return irs.PublicIPInfo{}, nil
}

func (publicIPHandler *ClouditPublicIPHandler) ListPublicIP() ([]*irs.PublicIPInfo, error) {
	publicIPHandler.Client.TokenID = publicIPHandler.CredentialInfo.AuthToken
	authHeader := publicIPHandler.Client.AuthenticatedHeaders()

	requestOpts := client.RequestOpts{
		//JSONBody:     nil,
		//RawBody:      nil,
		//JSONResponse: nil,
		//OkCodes:      nil,
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

//Todo : GET 문서에 없음
func (publicIPHandler *ClouditPublicIPHandler) GetPublicIP(publicIPID string) (irs.PublicIPInfo, error) {

	return irs.PublicIPInfo{}, nil
}

//Todo : 테스트 필요  문서 API 포스트맨도 사용안됨(맞는 ip입력시 변화없음)
func (publicIPHandler *ClouditPublicIPHandler) DeletePublicIP(publicIPID string) (bool, error) {
	publicIPHandler.Client.TokenID = publicIPHandler.CredentialInfo.AuthToken
	authHeader := publicIPHandler.Client.AuthenticatedHeaders()

	requestOpts := client.RequestOpts{
		//JSONBody:     nil,
		//RawBody:      nil,
		//JSONResponse: nil,
		//OkCodes:      nil,
		MoreHeaders: authHeader,
	}

	publicIP, _ := adaptiveip.Delete(publicIPHandler.Client, publicIPID, &requestOpts)
	spew.Dump(publicIP)

	return true, nil
}
