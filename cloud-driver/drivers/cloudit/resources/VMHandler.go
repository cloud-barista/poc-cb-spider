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
	var authHeader map[string]string
	authHeader = make(map[string]string)
	authHeader["X-Auth-Token"] = vmHandler.CredentialInfo.AuthToken
	requestOpts := client.RequestOpts{
		//JSONBody:     nil,
		//RawBody:      nil,
		//JSONResponse: nil,
		//OkCodes:      nil,
		MoreHeaders: authHeader,
	}

	vm, _ := server.Start(vmHandler.Client, &requestOpts)
	spew.Dump(vm)

	return irs.VMInfo{}, nil
}

func (vmHandler *ClouditVMHandler) SuspendVM(vmID string) {
	var authHeader map[string]string
	authHeader = make(map[string]string)
	authHeader["X-Auth-Token"] = vmHandler.CredentialInfo.AuthToken
	requestOpts := client.RequestOpts{
		//JSONBody:     nil,
		//RawBody:      nil,
		//JSONResponse: nil,
		//OkCodes:      nil,
		MoreHeaders: authHeader,
	}

	vm, _ := server.Suspend(vmHandler.Client, vmID, &requestOpts)
	spew.Dump(vm)
}

func (vmHandler *ClouditVMHandler) ResumeVM(vmID string) {
	var authHeader map[string]string
	authHeader = make(map[string]string)
	authHeader["X-Auth-Token"] = vmHandler.CredentialInfo.AuthToken
	requestOpts := client.RequestOpts{
		//JSONBody:     nil,
		//RawBody:      nil,
		//JSONResponse: nil,
		//OkCodes:      nil,
		MoreHeaders: authHeader,
	}

	vm, _ := server.Resume(vmHandler.Client, vmID, &requestOpts)
	spew.Dump(vm)
}

func (vmHandler *ClouditVMHandler) RebootVM(vmID string) {
	var authHeader map[string]string
	authHeader = make(map[string]string)
	authHeader["X-Auth-Token"] = vmHandler.CredentialInfo.AuthToken
	requestOpts := client.RequestOpts{
		//JSONBody:     nil,
		//RawBody:      nil,
		//JSONResponse: nil,
		//OkCodes:      nil,
		MoreHeaders: authHeader,
	}

	vm, _ := server.Reboot(vmHandler.Client, vmID, &requestOpts)
	spew.Dump(vm)
}

func (vmHandler *ClouditVMHandler) TerminateVM(vmID string) {
	var authHeader map[string]string
	authHeader = make(map[string]string)
	authHeader["X-Auth-Token"] = vmHandler.CredentialInfo.AuthToken
	requestOpts := client.RequestOpts{
		//JSONBody:     nil,
		//RawBody:      nil,
		//JSONResponse: nil,
		//OkCodes:      nil,
		MoreHeaders: authHeader,
	}

	vm, _ := server.Terminate(vmHandler.Client, vmID, &requestOpts)
	spew.Dump(vm)
}

func (vmHandler *ClouditVMHandler) ListVMStatus() []*irs.VMStatusInfo {
	var authHeader map[string]string
	authHeader = make(map[string]string)
	authHeader["X-Auth-Token"] = vmHandler.CredentialInfo.AuthToken
	requestOpts := client.RequestOpts{
		//JSONBody:     nil,
		//RawBody:      nil,
		//JSONResponse: nil,
		//OkCodes:      nil,
		MoreHeaders: authHeader,
	}

	vm, _ := server.ListStatus(vmHandler.Client, &requestOpts)
	spew.Dump(vm)

	return nil
}

/*
# OpenStack Credential Info
  domain_name: {domain_name}
  identity_endpoint: http://{ip}:5000/v3
  password: {password}
  project_id: {project_id}
  username: {username}
  region: {region}

  # OpenStack VM Deployment Info
  vm_name: {vm_name}
  image_id: {image_id}
  flavor_id: {flavor_id}
  network_id: {network_id}
  security_groups: {security_groups}
  keypair_name: {keypair_name}

  # OpenStack Test VM Info
  server_id: {server_id}

  image_info:
    name: mcb_custom_image.iso

  keypair_info:
    name: mcb-test-key

  security_group_info:
    name: mcb-test-security

  vnet_info:
    name: mcb-test-vnet
*/

func (vmHandler *ClouditVMHandler) GetVMStatus(vmID string) irs.VMStatus {
	var authHeader map[string]string
	authHeader = make(map[string]string)
	authHeader["X-Auth-Token"] = vmHandler.CredentialInfo.AuthToken
	requestOpts := client.RequestOpts{
		//JSONBody:     nil,
		//RawBody:      nil,
		//JSONResponse: nil,
		//OkCodes:      nil,
		MoreHeaders: authHeader,
	}

	vm, _ := server.GetStatus(vmHandler.Client, vmID, &requestOpts)
	spew.Dump(vm)

	return irs.VMStatus("")
}

func (vmHandler *ClouditVMHandler) ListVM() []*irs.VMInfo {
	var authHeader map[string]string
	authHeader = make(map[string]string)
	authHeader["X-Auth-Token"] = vmHandler.CredentialInfo.AuthToken
	requestOpts := client.RequestOpts{
		//JSONBody:     nil,
		//RawBody:      nil,
		//JSONResponse: nil,
		//OkCodes:      nil,
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
	var authHeader map[string]string
	authHeader = make(map[string]string)
	authHeader["X-Auth-Token"] = vmHandler.CredentialInfo.AuthToken
	requestOpts := client.RequestOpts{
		//JSONBody:     nil,
		//RawBody:      nil,
		//JSONResponse: nil,
		//OkCodes:      nil,
		MoreHeaders: authHeader,
	}

	vm, _ := server.Get(vmHandler.Client, vmID, &requestOpts)
	spew.Dump(vm)
	return irs.VMInfo{}
}

func mappingServerInfo(server interface{}) irs.VMInfo {
	return irs.VMInfo{}
}
