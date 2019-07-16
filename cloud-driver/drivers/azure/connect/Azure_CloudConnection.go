package connect

import (
	"fmt"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	azrs "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/azure/resources"
)

type AzureCloudConnection struct{}

func (AzureCloudConnection) CreateVNetworkHandler() (irs.VNetworkHandler, error) {
	fmt.Println("Azure Cloud Driver: called CreateVNetworkHandler()!")
	return nil, nil
}

func (AzureCloudConnection) CreateImageHandler() (irs.ImageHandler, error) {
	return nil, nil
}

func (AzureCloudConnection) CreateSecurityHandler() (irs.SecurityHandler, error) {
	return nil, nil
}
func (AzureCloudConnection) CreateKeyPairHandler() (irs.KeyPairHandler, error) {
	return nil, nil
}
func (AzureCloudConnection) CreateVNicHandler() (irs.VNicHandler, error) {
	return nil, nil
}
func (AzureCloudConnection) CreatePublicIPHandler() (irs.PublicIPHandler, error) {
	return nil, nil
}

func (AzureCloudConnection) CreateVMHandler() (irs.VMHandler, error) {
	var vmHandler irs.VMHandler
	vmHandler = azrs.AzureVMHandler{}
	return vmHandler, nil
}

func (AzureCloudConnection) IsConnected() (bool, error) {
	return true, nil
}
func (AzureCloudConnection) Close() error {
	return nil
}
