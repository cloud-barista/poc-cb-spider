package nic

import (
	"fmt"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client"
	"time"
)

type VmNicInfo struct {
	TenantId        string
	VmId            string
	Type            string
	Mac 			string
	Dev             string
	Ip              string
	SubnetAddr      string
	Creator         string
	CreatedAt       time.Time
	VmName          string
	NetworkName     string
	AdaptiveIp      string
	State           string
	Template        string
	SpecName        string
	CpuNum          int
	MemSize         int
	VolumeSize      int
	Qos             int
	SecGroups       string
	SecGroupMapInfo map[string]string
	AdaptiveMapInfo map[string]string
}

func List(restClient *client.RestClient, id string) (*[]VmNicInfo, error){
	requestURL := restClient.CreateRequestBaseURL(client.ACE, "servers", id, "nics")
	fmt.Println(requestURL)

	var result client.Result
	if _, result.Err = restClient.Get(requestURL, &result.Body, nil); result.Err != nil {
		return nil, result.Err
	}

	var nic []VmNicInfo
	if err := result.ExtractInto(&nic); err != nil {
		return nil, err
	}
	return &nic, nil
}