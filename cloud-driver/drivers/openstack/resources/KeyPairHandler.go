package resources

import (
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/davecgh/go-spew/spew"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/compute/v2/extensions/keypairs"
	"github.com/rackspace/gophercloud/pagination"
)

type OpenStackKeyPairHandler struct {
	Client *gophercloud.ServiceClient
}

type KeyPairInfo struct {
	Name        string
	Fingerprint string
	PublicKey   string
	PrivateKey  string
	UserID      string
}

func (keyPairInfo *KeyPairInfo) setter(keypair keypairs.KeyPair) *KeyPairInfo {
	keyPairInfo.Name = keypair.Name
	keyPairInfo.Fingerprint = keypair.Fingerprint
	keyPairInfo.PublicKey = keypair.PublicKey
	keyPairInfo.PrivateKey = keypair.PrivateKey
	keyPairInfo.UserID = keypair.UserID

	return keyPairInfo
}

//create 1번
func (keyPairHandler *OpenStackKeyPairHandler) CreateKey(keyPairReqInfo irs.KeyPairReqInfo) (irs.KeyPairInfo, error) {

	// @TODO: keyPair 생성 요청 파라미터 정의 필요
	type KeyPairReqInfo struct {
		Name string
	}
	reqInfo := KeyPairInfo{
		Name: "mcb-test-key",
	}

	// Check keyPair Exists
	/**/

	create0pts := keypairs.CreateOpts{
		Name: reqInfo.Name,
	}
	keyPairInfo, err := keypairs.Create(keyPairHandler.Client, create0pts).Extract()
	if err != nil {
		return irs.KeyPairInfo{}, err
	}

	// @TODO: 생성된 keyPair 정보 리턴
	spew.Dump(keyPairInfo)
	return irs.KeyPairInfo{}, nil
}

func (keyPairHandler *OpenStackKeyPairHandler) ListKey() ([]*irs.KeyPairInfo, error) {
	var keyPairList []*KeyPairInfo

	pager := keypairs.List(keyPairHandler.Client)
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		// Get KeyPair
		list, err := keypairs.ExtractKeyPairs(page)
		if err != nil {
			return false, err
		}
		// Add to List
		for _, k := range list {
			keyPairInfo := new(KeyPairInfo).setter(k)
			keyPairList = append(keyPairList, keyPairInfo)
		}
		return true, nil
	})
	if err != nil {
		return nil, err
	}

	spew.Dump(keyPairList)
	return nil, nil
}

func (keyPairHandler *OpenStackKeyPairHandler) GetKey(keyPairID string) (irs.KeyPairInfo, error) {
	keyPair, err := keypairs.Get(keyPairHandler.Client, keyPairID).Extract()
	if err != nil {
		return irs.KeyPairInfo{}, nil
	}

	keyPairInfo := new(KeyPairInfo).setter(*keyPair)

	spew.Dump(keyPairInfo)
	return irs.KeyPairInfo{}, nil
}

func (keyPairHandler *OpenStackKeyPairHandler) DeleteKey(keyPairID string) (bool, error) {
	result := keypairs.Delete(keyPairHandler.Client, keyPairID)
	if result.Err != nil {
		return false, result.Err
	}
	return false, nil
}
