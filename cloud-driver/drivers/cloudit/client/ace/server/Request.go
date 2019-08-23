package compute

import (
	"fmt"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client"
	"time"
)

type ServerInfo struct {
	VolumeInfoList struct{}
	VmNicInfoList  struct{}
	NicMapInfo     []struct {
		Name    string
		Count   int
		Address string
	}
	PoolMapInfo []struct {
		Name       string
		Count      int
		PoolID     string
		Filesystem string
	}
	ID                string
	TenantID          string
	CpuNum            float32
	MemSize           float32
	VncPort           int
	RepeaterPort      int
	State             string
	NodeIp            string
	NodeHostName      string
	Name              string
	Protection        int
	CreatedAt         time.Time
	IsoId             string
	IsoPath           string
	Iso               string
	Template          string
	TemplateID        string
	OsType            string
	RootPassword      string
	HostName          string
	Creator           string
	VolumeId          string
	VolumeSize        string
	VolumeMode        string
	MacAddr           string
	Spec              string
	SpecId            string
	Pool              string
	PoolId            string
	Cycle             string
	Metric            string
	MigrationPort     string
	MigrationIp       string
	Cloudinit         string
	DeleteVolume      string
	ServerCount       string
	PrivateIp         string
	AdaptiveIp        string
	InitCloud         string
	ClusterId         string
	ClusterName       string
	NicType           string
	Secgroups         string
	Ip                string
	SubnetAddr        string
	DeviceId          string
	Description       string
	DiskSize          string
	DiskCount         string
	IsoInsertedAt     string
	Puppet            string
	TemplateOwnership string
	VmStatInfo        string
}

func List(restClient *client.RestClient) (*[]ServerInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.ACE, "servers")
	fmt.Println(requestURL)

	var result client.Result
	if _, result.Err = restClient.Get(requestURL, &result.Body, nil); result.Err != nil {
		return nil, result.Err
	}

	var server []ServerInfo
	if err := result.ExtractInto(&server); err != nil {
		return nil, err
	}
	return &server, nil

	/*var serverList []ServerInfo
	resultBody := result.Body.(map[string]string)
	for _, body := range resultBody {
		server := extractServer(body)
	}

	return nil*/
}

func Get(restClient *client.RestClient, id string) (*ServerInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.ACE, "servers", id)
	fmt.Println(requestURL)

	var result client.Result
	_, result.Err = restClient.Get(requestURL, &result.Body, nil)

	var server ServerInfo
	if err := result.ExtractInto(&server); err != nil {
		return nil, err
	}
	return &server, nil
	//return extractServer(result)
}

func extractServer(result client.Result) (*ServerInfo, error) {
	server := new(ServerInfo)
	if err := result.ExtractInto(&server); err != nil {
		return nil, err
	}
	return server, nil
}
