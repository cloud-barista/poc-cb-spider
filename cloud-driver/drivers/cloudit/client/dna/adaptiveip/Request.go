package adaptiveip

import (
	"fmt"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client"
)

type AdaptiveIPInfo struct {
	ID          string
	IP          string
	Name        string
	Rules       interface{}
	TenantId    string
	Creator     string
	State       string
	CreatedAt   string
	PrivateIp   string
	Protection  int
	RuleCount   int
	VmName      string
	Description string
}

func List(restClient *client.RestClient, requestOpts *client.RequestOpts) (*[]AdaptiveIPInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.DNA, "adaptive-ips")
	fmt.Println(requestURL)

	var result client.Result
	if _, result.Err = restClient.Get(requestURL, &result.Body, requestOpts); result.Err != nil {
		return nil, result.Err
	}

	var adaptiveIP []AdaptiveIPInfo
	if err := result.ExtractInto(&adaptiveIP); err != nil {
		return nil, err
	}
	return &adaptiveIP, nil
}

func Create(restClient *client.RestClient, requestOpts *client.RequestOpts) (*AdaptiveIPInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.DNA, "adaptive-ips")
	fmt.Println(requestURL)

	var result client.Result

	if _, result.Err = restClient.Post(requestURL, nil, &result.Body, requestOpts); result.Err != nil {
		return nil, result.Err
	}

	var adaptiveIP AdaptiveIPInfo
	if err := result.ExtractInto(&adaptiveIP); err != nil {
		return nil, err
	}

	return &adaptiveIP, nil
}

func Delete(restClient *client.RestClient, ip string, requestOpts *client.RequestOpts) error {
	requestURL := restClient.CreateRequestBaseURL(client.DNA, "adaptive-ips", ip)
	fmt.Println(requestURL)

	var result client.Result
	if _, result.Err = restClient.Delete(requestURL, requestOpts); result.Err != nil {
		return result.Err
	}

	return nil
}
