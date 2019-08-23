package subnet

import (
	"fmt"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client"
	"time"
)

type SubnetInfo struct {
	ID          string
	TenantId    string
	Addr        string
	Prefix      string
	Gateway     string
	Creator     string
	Protection  int
	Name        string
	State       string
	Vlan        int
	CreatedAt   time.Time
	NicCount    int
	Description string
}

func List(restClient *client.RestClient) (*[]SubnetInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.DNA, "")
	fmt.Println(requestURL)

	var result client.Result
	_, result.Err = restClient.Get(requestURL, &result.Body, nil)

	var subnet []SubnetInfo
	if err := result.ExtractInto(&subnet); err != nil {
		return nil, err
	}
	return &subnet, nil
}
