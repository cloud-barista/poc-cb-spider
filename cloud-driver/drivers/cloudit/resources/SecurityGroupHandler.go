package resources

import (
	"fmt"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client/iam/securitygroup"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/davecgh/go-spew/spew"
	"strconv"
)

type ClouditSecurityHandler struct {
	CredentialInfo idrv.CredentialInfo
	Client         *client.RestClient
}

//Todo 수정 중
func (securityHandler *ClouditSecurityHandler) CreateSecurity(securityReqInfo irs.SecurityReqInfo) (irs.SecurityInfo, error) {
	securityHandler.Client.TokenID = securityHandler.CredentialInfo.AuthToken
	authHeader := securityHandler.Client.AuthenticatedHeaders()

	type SecurityReqInfo struct {
		Name       string `json:"name" required:"true"`
		Protection int    `json:"protection" required:"false"`
		Rules      []struct {
			Name     string `json:"name" required:"false"`
			Protocol string `json:"protocol" required:"false"`
			Port     int    `json:"port" required:"false"`
			Target   string `json:"target" required:"false"`
			Type     string `json:"type" required:"false"`
		} `json:"rules" required:"false"`
	}

	reqInfo := SecurityReqInfo{
		Name:       "test-Dong333",
		Protection: 0,
		//Rules : "",
	}

	createOpts := client.RequestOpts{
		JSONBody:    reqInfo,
		MoreHeaders: authHeader,
	}

	security, err := securitygroup.Create(securityHandler.Client, &createOpts)
	if err != nil {
		return irs.SecurityInfo{}, err
	}

	spew.Dump(security)
	return irs.SecurityInfo{}, nil
}

func (securityHandler *ClouditSecurityHandler) ListSecurity() ([]*irs.SecurityInfo, error) {
	securityHandler.Client.TokenID = securityHandler.CredentialInfo.AuthToken
	authHeader := securityHandler.Client.AuthenticatedHeaders()

	requestOpts := client.RequestOpts{
		//JSONBody:     nil,
		//RawBody:      nil,
		//JSONResponse: nil,
		//OkCodes:      nil,
		MoreHeaders: authHeader,
	}

	securityList, err := securitygroup.List(securityHandler.Client, &requestOpts)
	if err != nil {
		panic(err)
	}

	for i, security := range *securityList {
		fmt.Println("[" + strconv.Itoa(i) + "]")
		spew.Dump(security)
	}
	return nil, nil
}

func (securityHandler *ClouditSecurityHandler) GetSecurity(securityID string) (irs.SecurityInfo, error) {
	securityHandler.Client.TokenID = securityHandler.CredentialInfo.AuthToken
	authHeader := securityHandler.Client.AuthenticatedHeaders()

	requestOpts := client.RequestOpts{
		//JSONBody:     nil,
		//RawBody:      nil,
		//JSONResponse: nil,
		//OkCodes:      nil,
		MoreHeaders: authHeader,
	}

	securityRuleList, err := securitygroup.Get(securityHandler.Client, securityID, &requestOpts)
	if err != nil {
		panic(err)
	}

	for i, securityRule := range *securityRuleList {
		fmt.Println("[" + strconv.Itoa(i) + "]")
		spew.Dump(securityRule)
	}

	return irs.SecurityInfo{}, nil
}

func (securityHandler *ClouditSecurityHandler) DeleteSecurity(securityID string) (bool, error) {
	securityHandler.Client.TokenID = securityHandler.CredentialInfo.AuthToken
	authHeader := securityHandler.Client.AuthenticatedHeaders()

	requestOpts := client.RequestOpts{
		//JSONBody:     nil,
		//RawBody:      nil,
		//JSONResponse: nil,
		//OkCodes:      nil,
		MoreHeaders: authHeader,
	}

	security, _ := securitygroup.Delete(securityHandler.Client, securityID, &requestOpts)
	spew.Dump(security)

	return true, nil
}
