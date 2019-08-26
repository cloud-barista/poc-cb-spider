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
	var keyPairList []*irs.KeyPairInfo
	//spew.Dump(keyPairHandler)
	cblogger.Info(keyPairHandler)

	input := &ec2.DescribeKeyPairsInput{
		KeyNames: []*string{
			nil,
		},
	}

	//  Returns a list of key pairs
	//result, err := keyPairHandler.Client.DescribeKeyPairs(nil)
	result, err := keyPairHandler.Client.DescribeKeyPairs(input)
	cblogger.Info(result)
	if err != nil {
		//exitErrorf("Unable to get key pairs, %v", err)
		cblogger.Errorf("Unable to get key pairs, %v", err)
		return keyPairList, err
	}

	//	fmt.Println("Key Pairs:")
	cblogger.Info("Key Pairs:")
	for _, pair := range result.KeyPairs {
		//fmt.Printf("%s: %s\n", *pair.KeyName, *pair.KeyFingerprint)
		cblogger.Debugf("%s: %s\n", *pair.KeyName, *pair.KeyFingerprint)
		keyPairInfo := new(irs.KeyPairInfo)
		keyPairInfo.Name = *pair.KeyName
		keyPairInfo.Id = *pair.KeyFingerprint

		keyPairList = append(keyPairList, keyPairInfo)
	}

	cblogger.Info(keyPairList)
	//spew.Dump(keyPairList)
	//return keyPairList, nil
	return keyPairList, nil
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
	cblogger.Infof("GetKey : [%s]", keyPairID)

	input := &ec2.DescribeKeyPairsInput{
		KeyNames: []*string{
			aws.String(keyPairID),
		},
	}

	result, err := keyPairHandler.Client.DescribeKeyPairs(input)
	cblogger.Info("result : ", result)
	cblogger.Info("err : ", err)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			cblogger.Info("aerr : ", aerr)
			cblogger.Info("aerr.Code()  : ", aerr.Code())
			cblogger.Info("ok : ", ok)
			switch aerr.Code() {
			default:
				//fmt.Println(aerr.Error())
				cblogger.Error(aerr.Error())
				return irs.KeyPairInfo{}, aerr
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			//fmt.Println(err.Error())
			cblogger.Error(err.Error())
			return irs.KeyPairInfo{}, err
		}
		return irs.KeyPairInfo{}, nil
	}

	cblogger.Info("KeyName : ", *result.KeyPairs[0].KeyName)
	cblogger.Info("Fingerprint : ", *result.KeyPairs[0].KeyFingerprint)

	keyPairInfo := irs.KeyPairInfo{
		Name: *result.KeyPairs[0].KeyName,
		Id:   *result.KeyPairs[0].KeyFingerprint,
	}

	/*
		keyPairInfo := KeyPairInfo{
			Name:        *result.KeyPairs[0].KeyName,
			Fingerprint: *result.KeyPairs[0].KeyFingerprint,
		}
	*/
	return keyPairInfo, nil
}

func (keyPairHandler *AwsKeyPairHandler) DeleteKey(keyPairID string) (bool, error) {
	cblogger.Infof("DeleteKeyPaid : [%s]", keyPairID)
	// Delete the key pair by name
	_, err := keyPairHandler.Client.DeleteKeyPair(&ec2.DeleteKeyPairInput{
		KeyName: aws.String(keyPairID),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == "InvalidKeyPair.Duplicate" {
			//exitErrorf("Key pair %q does not exist.", pairName)
			cblogger.Error("Key pair %q does not exist.", keyPairID)
			return false, err
		}
		//exitErrorf("Unable to delete key pair: %s, %v.", keyPairID, err)
		cblogger.Errorf("Unable to delete key pair: %s, %v.", keyPairID, err)
		return false, err
	}

	//fmt.Printf("Successfully deleted %q key pair\n", keyPairID)
	cblogger.Infof("Successfully deleted %q key pair\n", keyPairID)

	return true, nil
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
