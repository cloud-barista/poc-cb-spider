package securitygroup

import (
	"fmt"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client"
)

//Rules[]
type SecurityGroupRules struct {
	ID         string
	SecGroupID string
	Name       string
	Type       string
	Port       string
	Target     string
	Protocol   string
	Creator    string
	//CreatedAt  time.Time
}

type SecurityGroupInfo struct {
	ID       string
	Name     string
	TenantID string
	Creator  string
	State    string
	//CreatedAt  time.Time
	Protection int
	//Rules      []SecurityGroupRules
	Rules []struct {
		ID         string
		SecGroupID string
		Name       string
		Type       string
		Port       string
		Target     string
		Protocol   string
		Creator    string
		//CreatedAt  time.Time
	}
	RulesCount  int
	Description string
	AsID        string
}

func List(restClient *client.RestClient, requestOpts *client.RequestOpts) (*[]SecurityGroupInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.IAM, "securitygroups")
	fmt.Println(requestURL)

	var result client.Result
	if _, result.Err = restClient.Get(requestURL, &result.Body, requestOpts); result.Err != nil {
		return nil, result.Err
	}

	var securityGroup []SecurityGroupInfo
	if err := result.ExtractInto(&securityGroup); err != nil {
		return nil, err
	}
	return &securityGroup, nil
}

// Todo : 단일조회 제공이 안되서 SecurityGroup내에 Rule데이터 읽어오기(임시로) 넣어놈
func Get(restClient *client.RestClient, id string, requestOpts *client.RequestOpts) (*[]SecurityGroupInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.IAM, "securitygroups", id)
	fmt.Println(requestURL)

	var result client.Result
	if _, result.Err = restClient.Get(requestURL, &result.Body, requestOpts); result.Err != nil {
		return nil, result.Err
	}

	var securityGroup []SecurityGroupInfo
	if err := result.ExtractInto(&securityGroup); err != nil {
		return nil, err
	}
	return &securityGroup, nil

	//return extractServer(result)
}

func Create(restClient *client.RestClient, requestOpts *client.RequestOpts) (SecurityGroupInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.IAM, "securitygroups")
	fmt.Println(requestURL)

	var result client.Result
	_, result.Err = restClient.Post(requestURL, requestOpts.JSONBody, &result.Body, requestOpts)

	var securityGroup SecurityGroupInfo
	if err := result.ExtractInto(&securityGroup); err != nil {
		return SecurityGroupInfo{}, nil
	}

	return securityGroup, nil
}

func Delete(restClient *client.RestClient, securitygroupId string, requestOpts *client.RequestOpts) (*[]SecurityGroupInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.IAM, "securitygroups", securitygroupId)
	fmt.Println(requestURL)

	var result client.Result
	if _, result.Err = restClient.Delete(requestURL, requestOpts); result.Err != nil {
		return nil, result.Err
	}

	var securityGroup []SecurityGroupInfo
	if err := result.ExtractInto(&securityGroup); err != nil {
		return nil, err
	}

	return &securityGroup, nil
}

/*func extractSecurityGroup(result client.Result) (*SecurityGroupInfo, error) {
	securityGroup := new(SecurityGroupInfo)
	if err := result.ExtractInto(&securityGroup); err != nil {
		return nil, err
	}
	return securityGroup, nil
}*/
