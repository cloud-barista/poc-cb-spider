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
	irs "github.com/hyokyungk/poc-cb-spider/cloud-driver/interfaces/resources"
)

type TADCloudConnection struct{}

func (TADCloudConnection) CreateVNetworkHandler() (irs.VNetworkHandler, error) {
        fmt.Println("TEST A Cloud Driver: called CreateVNetworkHandler()!")
        return nil, nil
}


func (TADCloudConnection) CreateImageHandler() (irs.ImageHandler, error) {
	return nil, nil
}

func (TADCloudConnection) CreateSecurityHandler() (irs.SecurityHandler, error) {
	return nil, nil
}
func (TADCloudConnection) CreateKeyPairHandler() (irs.KeyPairHandler, error) {
	return nil, nil
}
func (TADCloudConnection) CreateVNicHandler() (irs.VNicHandler, error) {
	return nil, nil
}
func (TADCloudConnection) CreatePublicIPHandler() (irs.PublicIPHandler, error) {
	return nil, nil
}

func (TADCloudConnection) CreateVMHandler() (irs.VMHandler, error) {
	return nil, nil
}

func (TADCloudConnection) IsConnected() (bool, error) {
	return true, nil
}
func (TADCloudConnection) Close() error {
	return nil
}

