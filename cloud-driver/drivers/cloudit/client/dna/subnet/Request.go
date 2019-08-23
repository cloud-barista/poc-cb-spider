package dna

import (
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
}