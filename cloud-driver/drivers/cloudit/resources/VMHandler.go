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
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/davecgh/go-spew/spew"
	"strconv"
)

type ClouditVMHandler struct {
	Client *client.RestClient
}

func (vmHandler *ClouditVMHandler) StartVM(vmReqInfo irs.VMReqInfo) (irs.VMInfo, error) {
	return irs.VMInfo{}, nil
}

func (vmHandler *ClouditVMHandler) SuspendVM(vmID string) {

}

func (vmHandler *ClouditVMHandler) ResumeVM(vmID string) {

}

func (vmHandler *ClouditVMHandler) RebootVM(vmID string) {

}

func (vmHandler *ClouditVMHandler) TerminateVM(vmID string) {

}

func (vmHandler *ClouditVMHandler) ListVMStatus() []*irs.VMStatusInfo {
	return nil
}

func (vmHandler *ClouditVMHandler) GetVMStatus(vmID string) irs.VMStatus {
	return irs.VMStatus("")
}

func (vmHandler *ClouditVMHandler) ListVM() []*irs.VMInfo {
	vmList, _ := server.List(vmHandler.Client)
	for i, vm := range *vmList {
		fmt.Println("["+ strconv.Itoa(i) +"]")
		spew.Dump(vm)
	}
	return nil
}

func (vmHandler *ClouditVMHandler) GetVM(vmID string) irs.VMInfo {
	vm, _ := server.Get(vmHandler.Client, vmID)
	spew.Dump(vm)
	return irs.VMInfo{}
}

func mappingServerInfo(server interface{}) irs.VMInfo {
	return irs.VMInfo{}
}
