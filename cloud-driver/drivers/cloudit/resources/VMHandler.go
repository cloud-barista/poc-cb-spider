// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is a Cloud Driver Example for PoC Test.
//
// by hyokyung.kim@innogrid.co.kr, 2019.08.

package resources

import (
	"fmt"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client/ace/server"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/davecgh/go-spew/spew"
	"strconv"
)

type ClouditVMHandler struct {
	CredentialInfo idrv.CredentialInfo
	Client         *client.RestClient
}

func (vmHandler *ClouditVMHandler) StartVM(vmReqInfo irs.VMReqInfo) (irs.VMInfo, error) {
	vmHandler.Client.TokenID = vmHandler.CredentialInfo.AuthToken
	authHeader := vmHandler.Client.AuthenticatedHeaders()

	// @TODO: VM 생성 요청 파라미터 정의 필요
	type SecGroupInfo struct {
		Id string `json:"id" required:"true"`
	}
	type VMReqInfo struct {
		TemplateId   string         `json:"templateId" required:"true"`
		SpecId       string         `json:"specId" required:"true"`
		Name         string         `json:"name" required:"true"`
		HostName     string         `json:"hostName" required:"true"`
		RootPassword string         `json:"rootPassword" required:"true"`
		SubnetAddr   string         `json:"subnetAddr" required:"true"`
		Secgroups    []SecGroupInfo `json:"secgroups" required:"true"`
		Description  int            `json:"description" required:"false"`
		Protection   int            `json:"protection" required:"false"`
	}

	reqInfo := VMReqInfo{
		TemplateId:   vmReqInfo.ImageInfo.Id,
		SpecId:       vmReqInfo.SpecID,
		Name:         vmReqInfo.Name,
		HostName:     vmReqInfo.Name,
		RootPassword: vmReqInfo.LoginInfo.AdminPassword,
		SubnetAddr:   vmReqInfo.VNetworkInfo.Id,
		Secgroups: []SecGroupInfo{
			{Id: vmReqInfo.SecurityInfo.Id},
		},
	}

	requestOpts := client.RequestOpts{
		MoreHeaders: authHeader,
		JSONBody:    reqInfo,
	}

	spew.Dump(requestOpts)
	vm, err := server.Start(vmHandler.Client, &requestOpts)
	if err != nil {
		return irs.VMInfo{}, err
	}
	spew.Dump(vm)

	return irs.VMInfo{}, nil
}

func (vmHandler *ClouditVMHandler) SuspendVM(vmID string) {
	vmHandler.Client.TokenID = vmHandler.CredentialInfo.AuthToken
	authHeader := vmHandler.Client.AuthenticatedHeaders()

	requestOpts := client.RequestOpts{
		MoreHeaders: authHeader,
	}

	err := server.Suspend(vmHandler.Client, vmID, &requestOpts)
	if err != nil {
		panic(err)
	}
}

func (vmHandler *ClouditVMHandler) ResumeVM(vmID string) {
	vmHandler.Client.TokenID = vmHandler.CredentialInfo.AuthToken
	authHeader := vmHandler.Client.AuthenticatedHeaders()

	requestOpts := client.RequestOpts{
		MoreHeaders: authHeader,
	}

	err := server.Resume(vmHandler.Client, vmID, &requestOpts)
	if err != nil {
		panic(err)
	}
}

func (vmHandler *ClouditVMHandler) RebootVM(vmID string) {
	vmHandler.Client.TokenID = vmHandler.CredentialInfo.AuthToken
	authHeader := vmHandler.Client.AuthenticatedHeaders()

	requestOpts := client.RequestOpts{
		MoreHeaders: authHeader,
	}

	err := server.Reboot(vmHandler.Client, vmID, &requestOpts)
	if err != nil {
		panic(err)
	}
}

func (vmHandler *ClouditVMHandler) TerminateVM(vmID string) {
	vmHandler.Client.TokenID = vmHandler.CredentialInfo.AuthToken
	authHeader := vmHandler.Client.AuthenticatedHeaders()

	requestOpts := client.RequestOpts{
		MoreHeaders: authHeader,
	}

	err := server.Terminate(vmHandler.Client, vmID, &requestOpts)
	if err != nil {
		panic(err)
	}
}

func (vmHandler *ClouditVMHandler) ListVMStatus() []*irs.VMStatusInfo {
	vmHandler.Client.TokenID = vmHandler.CredentialInfo.AuthToken
	authHeader := vmHandler.Client.AuthenticatedHeaders()

	requestOpts := client.RequestOpts{
		MoreHeaders: authHeader,
	}

	vmList, err := server.List(vmHandler.Client, &requestOpts)
	if err != nil {
		panic(err)
	}

	var vmStatusList []*irs.VMStatusInfo

	for _, vm := range *vmList {
		vmStatusInfo := irs.VMStatusInfo{
			VmId:     vm.ID,
			VmStatus: irs.VMStatus(vm.State),
		}
		vmStatusList = append(vmStatusList, &vmStatusInfo)
	}

	return vmStatusList
}

func (vmHandler *ClouditVMHandler) GetVMStatus(vmID string) irs.VMStatus {
	vmHandler.Client.TokenID = vmHandler.CredentialInfo.AuthToken
	authHeader := vmHandler.Client.AuthenticatedHeaders()

	requestOpts := client.RequestOpts{
		MoreHeaders: authHeader,
	}

	vm, _ := server.Get(vmHandler.Client, vmID, &requestOpts)
	return irs.VMStatus(vm.State)
}

func (vmHandler *ClouditVMHandler) ListVM() []*irs.VMInfo {
	vmHandler.Client.TokenID = vmHandler.CredentialInfo.AuthToken
	authHeader := vmHandler.Client.AuthenticatedHeaders()

	requestOpts := client.RequestOpts{
		MoreHeaders: authHeader,
	}

	vmList, err := server.List(vmHandler.Client, &requestOpts)
	if err != nil {
		panic(err)
	}

	for i, vm := range *vmList {
		fmt.Println("[" + strconv.Itoa(i) + "]")
		spew.Dump(vm)
	}
	return nil
}

func (vmHandler *ClouditVMHandler) GetVM(vmID string) irs.VMInfo {
	vmHandler.Client.TokenID = vmHandler.CredentialInfo.AuthToken
	authHeader := vmHandler.Client.AuthenticatedHeaders()

	requestOpts := client.RequestOpts{
		MoreHeaders: authHeader,
	}

	vm, _ := server.Get(vmHandler.Client, vmID, &requestOpts)
	spew.Dump(vm)
	return irs.VMInfo{}
}
