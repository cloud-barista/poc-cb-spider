package resources

import (
	"fmt"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/config"
	_ "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/config"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/davecgh/go-spew/spew"
	"github.com/rackspace/gophercloud"
	_ "github.com/rackspace/gophercloud/openstack/compute/v2/extensions/keypairs"
	"github.com/rackspace/gophercloud/openstack/compute/v2/extensions/secgroups"
	_ "github.com/rackspace/gophercloud/openstack/compute/v2/extensions/secgroups"
	_ "github.com/rackspace/gophercloud/openstack/compute/v2/extensions/startstop"
	_ "github.com/rackspace/gophercloud/openstack/compute/v2/servers"
	_ "github.com/rackspace/gophercloud/openstack/networking/v2/extensions/security/rules"
	"github.com/rackspace/gophercloud/pagination"
	_ "github.com/rackspace/gophercloud/pagination"
)

type OpenStackSecurityHandler struct {
	Client *gophercloud.ServiceClient
}

type SecurityInfo struct {
	ID          string
	Name        string
	Description string

	//Rules []rules.SecGroupRule `json:"security_group_rules" mapstructure:"security_group_rules"`
	Rules    []secgroups.Rule `json:"security_group_rules" mapstructure:"security_group_rules"`
	TenantID string           `json:"tenant_id" mapstructure:"tenant_id"`
}

func (securityInfo *SecurityInfo) setter(results secgroups.SecurityGroup) SecurityInfo {

	securityInfo.ID = results.ID
	securityInfo.Name = results.Name
	securityInfo.Description = results.Description
	securityInfo.Rules = results.Rules
	securityInfo.TenantID = results.TenantID

	return *securityInfo
}

func (securityInfo *SecurityInfo) printfInfo() {
	fmt.Println("ID : ", securityInfo.ID)
	fmt.Println("Name : ", securityInfo.Name)
	fmt.Println("Description : ", securityInfo.Description)
	fmt.Println("Rules : ", securityInfo.Rules)
	fmt.Println("TenantID : ", securityInfo.TenantID)
}

func (securityHandler *OpenStackSecurityHandler) CreateSecurity(securityReqInfo irs.SecurityReqInfo) (irs.SecurityInfo, error) {

	opts := secgroups.CreateOpts{
		Name:        "mcb-sg",
		Description: "mcb-sg",
	}
	group, err := secgroups.Create(securityHandler.Client, opts).Extract()
	if err != nil {
		return irs.SecurityInfo{}, err
	}
	spew.Dump(group)

	ruleOpts := secgroups.CreateRuleOpts{
		ParentGroupID: group.ID,
		FromPort:      22,
		ToPort:        22,
		IPProtocol:    "TCP",
		CIDR:          "0.0.0.0/0",
	}
	sgRule, err := secgroups.CreateRule(securityHandler.Client, ruleOpts).Extract()
	if err != nil {
		panic(err)
	}
	spew.Dump(sgRule)
	return irs.SecurityInfo{}, nil
}

func (securityHandler *OpenStackSecurityHandler) ListSecurity() ([]*irs.SecurityInfo, error) {
	var secList = make([]SecurityInfo, 20)

	pager := secgroups.List(securityHandler.Client)
	err := pager.EachPage(func(page pagination.Page) (bool, error) {

		securityList, err := secgroups.ExtractSecurityGroups(page)
		if err != nil {
			return false, nil
		}

		var secuiryInfo SecurityInfo

		for _, securitys := range securityList {
			security := secuiryInfo.setter(securitys)
			secList = append(secList, security)
			spew.Dump(security)
		}

		return true, nil
	})

	if err != nil {
		panic(err)
	}

	return nil, nil
}

func (securityHandler *OpenStackSecurityHandler) GetSecurity(securityID string) (irs.SecurityInfo, error) {
	client, err := config.GetServiceClient()
	if err != nil {
		panic(err)
	}

	security, err := secgroups.Get(client, securityID).Extract()

	var info SecurityInfo
	info.setter(*security)
	spew.Dump(info)

	return irs.SecurityInfo{}, nil
}

func (securityHandler *OpenStackSecurityHandler) DeleteSecurity(securityID string) (bool, error) {
	client, err := config.GetServiceClient()
	if err != nil {
		panic(err)
	}

	result := secgroups.Delete(client, securityID)
	fmt.Println("Delete : ", result)
	return false, nil
}
