package resources

import (
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/davecgh/go-spew/spew"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/compute/v2/extensions/secgroups"
	"github.com/rackspace/gophercloud/pagination"
)

type OpenStackSecurityHandler struct {
	Client *gophercloud.ServiceClient
}

// @TODO: SecurityInfo 리소스 프로퍼티 정의 필요
type SecurityInfo struct {
	ID          string
	Name        string
	Description string
	Rules       []SecurityRuleInfo
	TenantID    string
}
type SecurityRuleInfo struct {
	ID            string
	FromPort      int
	ToPort        int
	IPProtocol    string
	CIDR          string
	ParentGroupID string
	GroupTenantId string
	GroupName     string
}

func (securityInfo *SecurityInfo) setter(results secgroups.SecurityGroup) *SecurityInfo {
	securityInfo.ID = results.ID
	securityInfo.Name = results.Name
	securityInfo.Description = results.Description
	//securityInfo.Rules = results.Rules
	securityInfo.TenantID = results.TenantID

	return securityInfo
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
	var securityList []*SecurityInfo
	
	pager := secgroups.List(securityHandler.Client)
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		// Get SecurityGroup
		list, err := secgroups.ExtractSecurityGroups(page)
		if err != nil {
			return false, nil
		}
		// Add to List
		for _, s:= range list {
			securityInfo := new(SecurityInfo).setter(s)
			securityList = append(securityList, securityInfo)
		}
		return true, nil
	})
	if err != nil {
		panic(err)
	}
	
	spew.Dump(securityList)
	return nil, nil
}

func (securityHandler *OpenStackSecurityHandler) GetSecurity(securityID string) (irs.SecurityInfo, error) {
	securityGroup, err := secgroups.Get(securityHandler.Client, securityID).Extract()
	if err != nil {
		return irs.SecurityInfo{}, err
	}
	
	securityInfo := new(SecurityInfo).setter(*securityGroup)
	
	spew.Dump(securityInfo)
	return irs.SecurityInfo{}, nil
}

func (securityHandler *OpenStackSecurityHandler) DeleteSecurity(securityID string) (bool, error) {
	result := secgroups.Delete(securityHandler.Client, securityID)
	if result.Err != nil {
		return false, result.Err
	}
	return true, nil
}
