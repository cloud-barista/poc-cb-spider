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
	"github.com/aws/aws-sdk-go/service/ec2"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
)

type AwsPublicIPHandler struct {
	Region idrv.RegionInfo
	Client *ec2.EC2
}

func (publicIpHandler *AwsPublicIPHandler) CreatePublicIP(publicIPReqInfo irs.PublicIPReqInfo) (irs.PublicIPInfo, error) {

	return irs.PublicIPInfo{}, nil
}

func (publicIpHandler *AwsPublicIPHandler) ListPublicIP() ([]*irs.PublicIPInfo, error) {
	return nil, nil
}

func (publicIpHandler *AwsPublicIPHandler) GetPublicIP(publicIPID string) (irs.PublicIPInfo, error) {
	return irs.PublicIPInfo{}, nil
}

func (publicIpHandler *AwsPublicIPHandler) DeletePublicIP(publicIPID string) (bool, error) {
	return false, nil
}
