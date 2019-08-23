package adaptiveip

import (
	"fmt"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client"
	"time"
)

type AdaptiveIPInfo struct {
	ID          string
	IP          string
	Name        string
	Rules       interface{}
	TenantId    string
	Creator     string
	State       string
	CreatedAt   time.Time
	PrivateIp   string
	Protection  int
	RuleCount   int
	VmName      string
	Description string
}

func List(restClient *client.RestClient) (*[]AdaptiveIPInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.DNA, "adaptive-ips")
	fmt.Println(requestURL)
	
	var result client.Result
	_, result.Err = restClient.Get(requestURL, &result.Body, nil)
	
	var subnet []AdaptiveIPInfo
	if err := result.ExtractInto(&subnet); err != nil {
		return nil, err
	}
	return &subnet, nil
}
