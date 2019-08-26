package resources

import (
	"fmt"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client/iam/securitygroup"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/davecgh/go-spew/spew"
	"strconv"
)

type ClouditSecurityHandler struct {
	Client *client.RestClient
}

func (securityHandler *ClouditSecurityHandler) CreateSecurity(securityReqInfo irs.SecurityReqInfo) (irs.SecurityInfo, error) {
	return irs.SecurityInfo{}, nil
}

func (securityHandler *ClouditSecurityHandler) ListSecurity() ([]*irs.SecurityInfo, error) {
	securityList, _ := securitygroup.List(securityHandler.Client)
	for i, security := range *securityList {
		fmt.Println("[" + strconv.Itoa(i) + "]")
		spew.Dump(security)
	}
	return nil, nil
}

func (securityHandler *ClouditSecurityHandler) GetSecurity(securityID string) (irs.SecurityInfo, error) {
	//security, _ := securitygroup.Get(securityHandler.Client,"b77a1163-01ef-4e14-9ffc-9fb626b367be")
	//spew.Dump(security)

	return irs.SecurityInfo{}, nil
}

func (securityHandler *ClouditSecurityHandler) DeleteSecurity(securityID string) (bool, error) {
	return true, nil
}
