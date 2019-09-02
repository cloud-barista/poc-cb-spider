package server

import (
	"fmt"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client"
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
	ID           string
	TenantID     string
	CpuNum       float32
	MemSize      float32
	VncPort      int
	RepeaterPort int
	State        string
	NodeIp       string
	NodeHostName string
	Name         string
	//Protection        int
	//CreatedAt         time.Time
	//CreatedAt         string
	//IsoId             string
	//IsoPath           string
	//Iso               string
	//Template          string
	//TemplateID        string
	//OsType            string
	//RootPassword      string
	//HostName          string
	//Creator           string
	//VolumeId          string
	//VolumeSize        int
	//VolumeMode        string
	//MacAddr           string
	//Spec              string
	//SpecId            string
	//Pool              string
	//PoolId            string
	//Cycle             int
	//Metric            int
	//MigrationPort     int
	//MigrationIp       string
	//Cloudinit         bool
	//DeleteVolume      bool
	//ServerCount       int
	//PrivateIp         string
	//AdaptiveIp        string
	//InitCloud         int
	//ClusterId         string
	//ClusterName       string
	//NicType           string
	//Secgroups         string
	//Ip                string
	//SubnetAddr        string
	//DeviceId          string
	//Description       string
	//DiskSize          int
	//DiskCount         int
	//IsoInsertedAt     time.Time
	//Puppet            int
	//TemplateOwnership string
	//VmStatInfo        string
}

// create
// TODO: 테스트 필요
func Start(restClient *client.RestClient, requestOpts *client.RequestOpts) (*[]ServerInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.ACE, "servers")
	fmt.Println(requestURL)

	var result client.Result
	if _, result.Err = restClient.Post(requestURL, nil, &result.Body, requestOpts); result.Err != nil {
		return nil, result.Err
	}

	var server []ServerInfo
	if err := result.ExtractInto(&server); err != nil {
		return nil, err
	}

	return &server, nil
}

//shutdown
// TODO: 테스트 필요
func Suspend(restClient *client.RestClient, id string, requestOpts *client.RequestOpts) (*[]ServerInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.ACE, "servers", id, "shutdown")
	fmt.Println(requestURL)

	var result client.Result
	if _, result.Err = restClient.Post(requestURL, nil, &result.Body, requestOpts); result.Err != nil {
		return nil, result.Err
	}

	var server []ServerInfo
	if err := result.ExtractInto(&server); err != nil {
		return nil, err
	}

	return &server, nil
}

//start
// TODO: 테스트 필요
func Resume(restClient *client.RestClient, id string, requestOpts *client.RequestOpts) (*[]ServerInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.ACE, "servers", id, "start")
	fmt.Println(requestURL)

	var result client.Result
	//post  2번재 body interface -> result.Body?
	if _, result.Err = restClient.Post(requestURL, nil, &result.Body, requestOpts); result.Err != nil {
		return nil, result.Err
	}

	var server []ServerInfo
	if err := result.ExtractInto(&server); err != nil {
		return nil, err
	}

	return &server, nil
}

//reboot
// TODO: 테스트 필요
func Reboot(restClient *client.RestClient, id string, requestOpts *client.RequestOpts) (*[]ServerInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.ACE, "servers", id, "reboot")
	fmt.Println(requestURL)

	var result client.Result
	//post  2번재 body interface -> result.Body?
	if _, result.Err = restClient.Post(requestURL, nil, &result.Body, requestOpts); result.Err != nil {
		return nil, result.Err
	}

	var server []ServerInfo
	if err := result.ExtractInto(&server); err != nil {
		return nil, err
	}

	return &server, nil
}

//delete
// TODO: 테스트 필요
func Terminate(restClient *client.RestClient, id string, requestOpts *client.RequestOpts) (*[]ServerInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.ACE, "servers", id)
	fmt.Println(requestURL)

	var result client.Result
	if _, result.Err = restClient.Delete(requestURL, requestOpts); result.Err != nil {
		return nil, result.Err
	}

	var server []ServerInfo
	if err := result.ExtractInto(&server); err != nil {
		return nil, err
	}

	return &server, nil
}

func ListStatus(restClient *client.RestClient, requestOpts *client.RequestOpts) (*[]ServerInfo, error) {
	return nil, nil
}

func GetStatus(restClient *client.RestClient, requestOpts *client.RequestOpts) (*ServerInfo, error) {

	return nil, nil
}

func List(restClient *client.RestClient, requestOpts *client.RequestOpts) (*[]ServerInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.ACE, "servers")
	fmt.Println(requestURL)

	var result client.Result
	if _, result.Err = restClient.Get(requestURL, &result.Body, requestOpts); result.Err != nil {
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

// TODO: 테스트 필요
func Get(restClient *client.RestClient, id string, requestOpts *client.RequestOpts) (*ServerInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.ACE, "servers", id)
	fmt.Println(requestURL)

	var result client.Result
	_, result.Err = restClient.Get(requestURL, &result.Body, requestOpts)

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
