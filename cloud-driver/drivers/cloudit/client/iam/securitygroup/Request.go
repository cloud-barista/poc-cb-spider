package securitygroup

import (
	"fmt"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client"
	"time"
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
	CreatedAt  time.Time
}

type SecurityGroupInfo struct {
	ID         string
	Name       string
	TenantID   string
	Creator    string
	State      string
	CreatedAt  time.Time
	Protection int
	Rules      []SecurityGroupRules
	/*Rules 	[]struct{
		ID         string
		SecGroupID string
		Name       string
		Type       string
		Port       string
		Target     string
		Protocol   string
		Creator    string
		CreatedAt  time.Time
	}*/
	RulesCount  int
	Description string
	AsID        string
}

func List(restClient *client.RestClient) (*[]SecurityGroupInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.IAM, "securitygroups")
	fmt.Println(requestURL)

	var result client.Result
	if _, result.Err = restClient.Get(requestURL, &result.Body, nil); result.Err != nil {
		return nil, result.Err
	}

	var securityGroup []SecurityGroupInfo
	if err := result.ExtractInto(&securityGroup); err != nil {
		return nil, err
	}
	return &securityGroup, nil
}

// 단일조회
func Get(restClient *client.RestClient, Id string) (*SecurityGroupInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.IAM, "securitygroups", Id)
	fmt.Println(requestURL)

	var result client.Result
	_, result.Err = restClient.Get(requestURL, &result.Body, nil)

	var securityGroup SecurityGroupInfo
	if err := result.ExtractInto(&securityGroup); err != nil {
		return nil, err
	}
	return &securityGroup, nil
	//return extractServer(result)
}

/*func extractSecurityGroup(result client.Result) (*SecurityGroupInfo, error) {
	securityGroup := new(SecurityGroupInfo)
	if err := result.ExtractInto(&securityGroup); err != nil {
		return nil, err
	}
	return securityGroup, nil
}*/
