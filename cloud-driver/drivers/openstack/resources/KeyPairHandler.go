package resources

import (
	"fmt"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/config"
	_ "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/config"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/davecgh/go-spew/spew"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/compute/v2/extensions/keypairs"
	_ "github.com/rackspace/gophercloud/openstack/compute/v2/extensions/keypairs"
	_ "github.com/rackspace/gophercloud/openstack/compute/v2/extensions/startstop"
	_ "github.com/rackspace/gophercloud/openstack/compute/v2/servers"
	"github.com/rackspace/gophercloud/pagination"
	_ "github.com/rackspace/gophercloud/pagination"
)

// modified by powerkim, 2019.07.29
type OpenStackKeyPairHandler struct {
	Client *gophercloud.ServiceClient
}

type KeyPairInfo struct {
	Name        string `mapstructure:"name"`
	Fingerprint string `mapstructure:"fingerprint"`
	PublicKey   string `mapstructure:"public_key"`
	PrivateKey  string `mapstructure:"private_key"`
	UserID      string `mapstructure:"user_id"`
}

func (keyPairInfo *KeyPairInfo) setter(results keypairs.KeyPair) KeyPairInfo {

	keyPairInfo.Name = results.Name
	keyPairInfo.Fingerprint = results.Fingerprint
	keyPairInfo.PublicKey = results.PublicKey
	keyPairInfo.PrivateKey = results.PrivateKey
	keyPairInfo.UserID = results.UserID

	return *keyPairInfo
}

func (keyPairInfo *KeyPairInfo) printInfo() {
	fmt.Println("Name : ", keyPairInfo.Name)
	fmt.Println("Fingerprint : ", keyPairInfo.Fingerprint)
	fmt.Println("PublicKey : ", keyPairInfo.PublicKey)
	fmt.Println("PrivateKey : ", keyPairInfo.PrivateKey)
	fmt.Println("UserID : ", keyPairInfo.UserID)
}

//create 1번
func (keyPairHandler *OpenStackKeyPairHandler) CreateKey(keyPairReqInfo irs.KeyPairReqInfo) (irs.KeyPairInfo, error) {

	pts := keypairs.CreateOpts{
		Name: "keypair-name2",
		//Name: keyPairReqInfo.Name,
	}

	keypair, err := keypairs.Create(keyPairHandler.Client, pts).Extract()

	if err != nil {
		panic(err)
	}

	//fmt.Println("%+v",keypair)
	spew.Dump(keypair)
	return irs.KeyPairInfo{}, nil
}

//create 2번
/*func (keyPairHandler *OpenStackKeyPairHandler) CreateKey(keyPairReqInfo irs.KeyPairReqInfo) (irs.KeyPairInfo, error) {
	client, err:= config.GetServiceClient()
	if err != nil {
		panic(err)
	}

	create0pts := keypairs.CreateOpts{
		Name: "ddddsa",
	}

	result, err := keypairs.Create(client, create0pts).Extract()

	fmt.Println("Create : ",result)

	return irs.KeyPairInfo{} ,nil
}*/

func (keyPairHandler *OpenStackKeyPairHandler) ListKey() ([]*irs.KeyPairInfo, error) {
	var keyList = make([]KeyPairInfo, 20)
	pager := keypairs.List(keyPairHandler.Client)
	err := pager.EachPage(func(page pagination.Page) (bool, error) {

		//Get server
		keyPairList, err := keypairs.ExtractKeyPairs(page)
		if err != nil {
			return false, err
		}

		var keyPairInfo KeyPairInfo

		//add to list
		for _, keypairs := range keyPairList {
			keyPair := keyPairInfo.setter(keypairs)
			keyList = append(keyList, keyPair)
			spew.Dump(keyPair)
			fmt.Println("-----------------------------------")
		}
		return true, nil
	})

	if err != nil {
		panic(err)
	}

	return nil, nil
}

func (keyPairHandler *OpenStackKeyPairHandler) GetKey(keyPairID string) (irs.KeyPairInfo, error) {
	client, err := config.GetServiceClient()
	if err != nil {
		panic(err)
	}

	keypair, err := keypairs.Get(client, keyPairID).Extract()

	var info KeyPairInfo
	info.setter(*keypair)
	spew.Dump(info)

	return irs.KeyPairInfo{}, nil
}

func (keyPairHandler *OpenStackKeyPairHandler) DeleteKey(keyPairID string) (bool, error) {
	client, err := config.GetServiceClient()
	if err != nil {
		panic(err)
	}
	result := keypairs.Delete(client, keyPairID)
	fmt.Println("Delete : ", result)
	return false, nil
}
