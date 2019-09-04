package subnet

import (
	"fmt"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client"
)

type SubnetInfo struct {
	ID         string
	TenantId   string
	Addr       string
	Prefix     string
	Gateway    string
	Creator    string
	Protection int
	Name       string
	State      string
	Vlan       int
	//CreatedAt   time.Time
	NicCount    int
	Description string
}

func Create(restClient *client.RestClient, requestOpts *client.RequestOpts) (SubnetInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.DNA, "subnets")
	fmt.Println(requestURL)

	var result client.Result
	_, result.Err = restClient.Post(requestURL, requestOpts.JSONBody, &result.Body, requestOpts)
	if result.Err != nil {
		return SubnetInfo{}, result.Err
	}

	var subnet SubnetInfo
	if err := result.ExtractInto(&subnet); err != nil {
		return SubnetInfo{}, err
	}
	return subnet, nil
}

func List(restClient *client.RestClient, requestOpts *client.RequestOpts) (*[]SubnetInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.DNA, "subnets")
	fmt.Println(requestURL)

	var result client.Result
	if _, result.Err = restClient.Get(requestURL, &result.Body, requestOpts); result.Err != nil {
		return nil, result.Err
	}

	var subnet []SubnetInfo
	if err := result.ExtractInto(&subnet); err != nil {
		return nil, err
	}
	return &subnet, nil
}

func Delete(restClient *client.RestClient, addr string, requestOpts *client.RequestOpts) (*[]SubnetInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.DNA, "subnets", addr)
	fmt.Println(requestURL)

	var result client.Result
	if _, result.Err = restClient.Delete(requestURL, requestOpts); result.Err != nil {
		return nil, result.Err
	}

	var subnet []SubnetInfo
	if err := result.ExtractInto(&subnet); err != nil {
		return nil, err
	}

	return &subnet, nil
}
