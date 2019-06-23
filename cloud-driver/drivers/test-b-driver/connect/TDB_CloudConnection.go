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

type TBDCloudConnection struct{}

func (TBDCloudConnection) CreateVNetworkHandler() (irs.VNetworkHandler, error) {
        fmt.Println("TEST B Cloud Driver: called CreateVNetworkHandler()!")
        return nil, nil
}


func (TBDCloudConnection) CreateImageHandler() (irs.ImageHandler, error) {
	return nil, nil
}

func (TBDCloudConnection) CreateSecurityHandler() (irs.SecurityHandler, error) {
	return nil, nil
}
func (TBDCloudConnection) CreateKeyPairHandler() (irs.KeyPairHandler, error) {
	return nil, nil
}
func (TBDCloudConnection) CreateVNicHandler() (irs.VNicHandler, error) {
	return nil, nil
}
func (TBDCloudConnection) CreatePublicIPHandler() (irs.PublicIPHandler, error) {
	return nil, nil
}

func (TBDCloudConnection) CreateVMHandler() (irs.VMHandler, error) {
	return nil, nil
}

func (TBDCloudConnection) IsConnected() (bool, error) {
	return true, nil
}
func (TBDCloudConnection) Close() error {
	return nil
}

