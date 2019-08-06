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

type KeyPairReqInfo struct {
	name string
	// @todo
}

type KeyPairInfo struct {
	Name string
	Id   string
	// @todo
}

type KeyPairHandler interface {
	CreateKey(keyPairReqInfo KeyPairReqInfo) (KeyPairInfo, error)
	ListKey() ([]*KeyPairInfo, error)
	GetKey(keyPairID string) (KeyPairInfo, error)
	DeleteKey(keyPairID string) (bool, error)
}
