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

	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
)

type AwsCloudConnection struct{}

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

func (AwsCloudConnection) CreateVMHandler() (irs.VMHandler, error) {
	return nil, nil
}

func (AwsCloudConnection) IsConnected() (bool, error) {
	return true, nil
}
func (AwsCloudConnection) Close() error {
	return nil
}