package resources

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
)

/*
var cblogger *logrus.Logger

func init() {
	// cblog is a global variable.
	cblogger = cblog.GetLogger("AWS VMHandler")
}
*/

type AwsKeyPairHandler struct {
	Region idrv.RegionInfo
	Client *ec2.EC2
}

// @TODO: KeyPairInfo 리소스 프로퍼티 정의 필요
type KeyPairInfo struct {
	Name        string
	Fingerprint string
}

func (keyPairHandler *AwsKeyPairHandler) ListKey() ([]*irs.KeyPairInfo, error) {
	cblogger.Debug("Start ListKey()")
	var keyPairList []*KeyPairInfo
	//spew.Dump(keyPairHandler)
	cblogger.Info(keyPairHandler)

	//  Returns a list of key pairs
	result, err := keyPairHandler.Client.DescribeKeyPairs(nil)
	if err != nil {
		//exitErrorf("Unable to get key pairs, %v", err)
		cblogger.Errorf("Unable to get key pairs, %v", err)
		return nil, err
	}

	//	fmt.Println("Key Pairs:")
	cblogger.Info("Key Pairs:")
	for _, pair := range result.KeyPairs {
		//fmt.Printf("%s: %s\n", *pair.KeyName, *pair.KeyFingerprint)
		cblogger.Debugf("%s: %s\n", *pair.KeyName, *pair.KeyFingerprint)
		//keyPairInfo := new(irs.KeyPairInfo)
		//keyPairInfo.Name = *pair.KeyName
		//keyPairInfo.Fingerprint = *pair.KeyFingerprint

		//keyPairList = append(keyPairList, keyPairInfo)
	}

	cblogger.Info(keyPairList)
	//spew.Dump(keyPairList)
	//return keyPairList, nil
	return nil, nil
}

func (keyPairHandler *AwsKeyPairHandler) CreateKey(keyPairReqInfo irs.KeyPairReqInfo) (irs.KeyPairInfo, error) {
	cblogger.Infof("Start CreateKey(%s)", keyPairReqInfo)

	// Creates a new  key pair with the given name
	result, err := keyPairHandler.Client.CreateKeyPair(&ec2.CreateKeyPairInput{
		KeyName: aws.String(keyPairReqInfo.Name),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == "InvalidKeyPair.Duplicate" {
			exitErrorf("Keypair %q already exists.", keyPairReqInfo.Name)
		}
		exitErrorf("Unable to create key pair: %s, %v.", keyPairReqInfo.Name, err)
	}

	fmt.Printf("Created key pair %q %s\n%s\n",
		*result.KeyName, *result.KeyFingerprint,
		*result.KeyMaterial)

	return irs.KeyPairInfo{}, nil
}

func (keyPairHandler *AwsKeyPairHandler) GetKey(keyPairID string) (irs.KeyPairInfo, error) {
	return irs.KeyPairInfo{}, nil
}

func (keyPairHandler *AwsKeyPairHandler) DeleteKey(keyPairID string) (bool, error) {
	return false, nil
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

/*
func (keyPairInfo *KeyPairInfo) setter(keypair keypairs.KeyPair) *KeyPairInfo {
	keyPairInfo.Name = keypair.Name
	keyPairInfo.Fingerprint = keypair.Fingerprint
	keyPairInfo.PublicKey = keypair.PublicKey
	keyPairInfo.PrivateKey = keypair.PrivateKey
	keyPairInfo.UserID = keypair.UserID

	return keyPairInfo
}

func (keyPairHandler *OpenStackKeyPairHandler) CreateKey(keyPairReqInfo irs.KeyPairReqInfo) (irs.KeyPairInfo, error) {

	create0pts := keypairs.CreateOpts{
		Name: keyPairReqInfo.Name,
	}
	keyPairInfo, err := keypairs.Create(keyPairHandler.Client, create0pts).Extract()
	if err != nil {
		return irs.KeyPairInfo{}, err
	}

	spew.Dump(keyPairInfo)
	return irs.KeyPairInfo{}, nil
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
	err := keypairs.Delete(keyPairHandler.Client, keyPairID).ExtractErr()
	if err != nil {
		return false, err
	}
	return true, nil
}
*/
