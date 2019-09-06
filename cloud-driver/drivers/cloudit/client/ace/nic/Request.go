package nic

import (
	"fmt"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client"
)

type VmNicInfo struct {
	TenantId    string
	VmId        string
	Type        string
	Mac         string
	Dev         string
	Ip          string
	SubnetAddr  string
	Creator     string
	CreatedAt   string
	VmName      string
	NetworkName string
	AdaptiveIp  string
	State       string
	Template    string
	SpecName    string
	CpuNum      string
	MemSize     string
	VolumeSize  string
	Qos         int
	SecGroups   string
	//SecGroupMapInfo map[string]string
	SecGroupMapInfo []struct {
		TenantId   string `json:"tenant_id"`
		SecGroupId string `json:"secgroup_id"`
		Name       string
		//Created_At	string	`json:"created_at"`
		Protection int
		State      string
		Mac        string
	}
	//AdaptiveMapInfo map[string]string
	AdaptiveMapInfo interface{}
}

func List(restClient *client.RestClient, serverId string, requestOpts *client.RequestOpts) (*[]VmNicInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.ACE, "servers", serverId, "nics")
	fmt.Println(requestURL)

	var result client.Result
	if _, result.Err = restClient.Get(requestURL, &result.Body, requestOpts); result.Err != nil {
		return nil, result.Err
	}

	var nic []VmNicInfo
	if err := result.ExtractInto(&nic); err != nil {
		return nil, err
	}
	return &nic, nil
}
