// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is a Cloud Driver Example for PoC Test.
//
// by powerkim@etri.re.kr, 2019.06.

package connect

import (
	"fmt"

	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"

	ars "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/aws/resources"

	//ec2drv "github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2"
)

//type AwsCloudConnection struct{}
type AwsCloudConnection struct {
	Region idrv.RegionInfo
	//Client *ec2.EC2
	VMClient *ec2.EC2
}

func (AwsCloudConnection) CreateVNetworkHandler() (irs.VNetworkHandler, error) {
	fmt.Println("TEST AWS Cloud Driver: called CreateVNetworkHandler()!")
	return nil, nil
}

func (AwsCloudConnection) CreateImageHandler() (irs.ImageHandler, error) {
	return nil, nil
}

func (AwsCloudConnection) CreateSecurityHandler() (irs.SecurityHandler, error) {
	return nil, nil
}
func (AwsCloudConnection) CreateKeyPairHandler() (irs.KeyPairHandler, error) {
	return nil, nil
}
func (AwsCloudConnection) CreateVNicHandler() (irs.VNicHandler, error) {
	return nil, nil
}
func (AwsCloudConnection) CreatePublicIPHandler() (irs.PublicIPHandler, error) {
	return nil, nil
}

/*
func (AwsCloudConnection) CreateVMHandler() (irs.VMHandler, error) {
	fmt.Println("TEST AWS Cloud Driver: called CreateVMHandler()!")
	return nil, nil
}
*/

func (cloudConn *AwsCloudConnection) CreateVMHandler() (irs.VMHandler, error) {
	fmt.Println("AWS Cloud Driver: called CreateVMHandler()!")
	//vmHandler := ars.AzureVMHandler{cloudConn.Region, cloudConn.Ctx, cloudConn.VMClient}
	vmHandler := ars.AwsVMHandler{cloudConn.Region, cloudConn.VMClient}
	return &vmHandler, nil
}

func (AwsCloudConnection) IsConnected() (bool, error) {
	return true, nil
}
func (AwsCloudConnection) Close() error {
	return nil
}
