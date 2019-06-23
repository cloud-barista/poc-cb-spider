// Cloud Driver Interface of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is Resouces interfaces of Cloud Driver.
//
// by powerkim@etri.re.kr, 2019.06.


package resources

import (
	"time"
)

type VMReqInfo struct {
	// region/zone: Do not specify, this driver already knew these in Connection.

	name string

	imageInfo ImageInfo 
	vNetworkInfo VNetworkInfo
	securityInfo SecurityInfo
	keyPairInfo KeyPairInfo
	specID string 	// instance type or flavour, etc...
	vNicInfo VNicInfo
	publicIPInfo PublicIPInfo
}

// GO do not support Enum. So, define like this. 
type VMStatus string
const (
        pending VMStatus = "PENDING" 		// from launch, suspended to running
        running VMStatus = "RUNNING"

        suspending VMStatus = "SUSPENDING"	// from running to suspended
	suspended VMStatus = "SUSPENDED"

        rebooting VMStatus = "REBOOTING"	// from running to running

        termiating VMStatus = "TERMINATING"	// from running, suspended to terminated
        termiated VMStatus = "TERMINATED"
)

type Region struct {
	region string
	zone string
}

type VMInfo struct {
	name string
	id string
	startTime time.Time 	// Timezone: based on cloud-barista server location.
	
	region Region		// ex) {us-east1, us-east1-c} or {ap-northeast-2}
	imageID string		// ex) ami-047f7b46bd6dd5d84 or projects/gce-uefi-images/global/images/centos-7-v20190326
	specID string 		// instance type or flavour, etc... ex) t2.micro or f1-micro
	vNetworkID string 	// ex) vpc-23ed0a4b
	subNetworkID string	// ex) subnet-8c4a53e4
	securityID string	// ex) sg-0b7452563e1121bb6
	
	vNIC string		// ex) eth0
	publicIP string		// ex) 13.125.43.21 
	publicDNS string	// ex) ec2-13-125-43-0.ap-northeast-2.compute.amazonaws.com
	privateIP string	// ex) ip-172-31-4-60.ap-northeast-2.compute.internal
	privateDNS string	// ex) 172.31.4.60
	
	keyPairID string	// ex) powerkimKeyPair
	guestUserID string	// ex) user1
	guestUserPwd string	

	guestBootDisk string	// ex) /dev/sda1
	guestBlockDisk string	// ex) 

	additionalInfo string // Any information to be good for users and developers.
}

type VMHandler interface {

	StartVM(vmReqInfo VMReqInfo) (VMInfo, error)
	SuspendVM(vmID string)
	ResumeVM(vmID string)
	RebootVM(vmID string)
	TerminateVM(vmID string)

	ListVMStatus() []*VMStatus 
	GetVMStatus(vmID string) VMStatus 

	ListVM() []*VMInfo 
	GetVM(vmID string) VMInfo
}

