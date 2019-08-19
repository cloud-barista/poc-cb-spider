// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is a Cloud Driver Example for PoC Test.
//
// by devunet@mz.co.kr, 2019.08.

package main

import (
	"fmt"
	"io/ioutil"
	"os"

	awsdrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/aws"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v3"
)

// Test VM Deployment
func createVM() {
	fmt.Println("Start Create VM ...")
	vmHandler, err := setVMHandler()
	if err != nil {
		panic(err)
	}
	config := readConfigFile()

	//minCount := aws.Int64(int64(config.Aws.MinCount))
	//maxCount := aws.Int64(config.Aws.MaxCount)

	vmReqInfo := irs.VMReqInfo{
		Name: config.Aws.BaseName,
		ImageInfo: irs.ImageInfo{
			Id: config.Aws.ImageID,
		},
		SpecID: config.Aws.InstanceType,
		SecurityInfo: irs.SecurityInfo{
			Id: config.Aws.SecurityGroupID,
		},
		KeyPairInfo: irs.KeyPairInfo{
			Name: config.Aws.KeyName,
		},
		VNetworkInfo: irs.VNetworkInfo{
			Id: config.Aws.SubnetID,
		},
	}
	vm, err := vmHandler.StartVM(vmReqInfo)
	if err != nil {
		panic(err)
	}
	spew.Dump(vm)

	fmt.Println("Finish Create VM")
}

/*
func suspendVM(vmID string) {
	fmt.Println("Start Suspend VM Test.. [" + vmID + "]")
	vmHandler, err := setVMHandler()
	if err != nil {
		panic(err)
	}

	vmHandler.SuspendVM(vmID)
	fmt.Println("Finish Suspend VM")
}

func resumeVM(vmID string) {
	fmt.Println("Start ResumeVM VM Test.. [" + vmID + "]")
	vmHandler, err := setVMHandler()
	if err != nil {
		panic(err)
	}

	vmHandler.ResumeVM(vmID)
	fmt.Println("Finish ResumeVM VM")
}
*/

// Test VM Lifecycle Management (Create/Suspend/Resume/Reboot/Terminate)
func handleVM() {
	vmHandler, err := setVMHandler()
	if err != nil {
		panic(err)
	}
	config := readConfigFile()

	for {
		fmt.Println("VM Management")
		fmt.Println("0. Start(Create) VM")
		fmt.Println("1. Suspend VM")
		fmt.Println("2. Resume VM")
		fmt.Println("3. Reboot VM")
		fmt.Println("4. Terminate VM")

		var commandNum int
		inputCnt, err := fmt.Scan(&commandNum)
		if err != nil {
			panic(err)
		}

		VmID := config.Aws.VmID

		if inputCnt == 1 {
			switch commandNum {
			case 0:
				createVM()
			case 1:
				fmt.Println("Start Suspend VM ...")
				vmHandler.SuspendVM(VmID)
				fmt.Println("Finish Suspend VM")
			case 2:
				fmt.Println("Start Resume  VM ...")
				vmHandler.ResumeVM(VmID)
				fmt.Println("Finish Resume VM")
			case 3:
				fmt.Println("Start Reboot  VM ...")
				vmHandler.RebootVM(VmID)
				fmt.Println("Finish Reboot VM")
			case 4:
				fmt.Println("Start Terminate  VM ...")
				vmHandler.TerminateVM(VmID)
				fmt.Println("Finish Terminate VM")
			}
		}
	}
}

func main() {
	fmt.Println("AWS Driver Test")
	//createVM()
	//suspendVM(vmID)
	//RebootVM
	//resumeVM(vmID)
	handleVM()
}

func setVMHandler() (irs.VMHandler, error) {
	var cloudDriver idrv.CloudDriver
	cloudDriver = new(awsdrv.AwsDriver)

	config := readConfigFile()
	connectionInfo := idrv.ConnectionInfo{
		CredentialInfo: idrv.CredentialInfo{
			ClientId:     config.Aws.AawsAccessKeyID,
			ClientSecret: config.Aws.AwsSecretAccessKey,
		},
		RegionInfo: idrv.RegionInfo{
			Region: config.Aws.Region,
		},
	}

	cloudConnection, err := cloudDriver.ConnectCloud(connectionInfo)
	if err != nil {
		return nil, err
	}

	vmHandler, err := cloudConnection.CreateVMHandler()
	if err != nil {
		return nil, err
	}
	return vmHandler, nil
}

// Region : 사용할 리전명 (ex) ap-northeast-2
// ImageID : VM 생성에 사용할 AMI ID (ex) ami-047f7b46bd6dd5d84
// BaseName : 다중 VM 생성 시 사용할 Prefix이름 ("BaseName" + "_" + "숫자" 형식으로 VM을 생성 함.) (ex) mcloud-barista
// VmID : 라이프 사이트클을 테스트할 EC2 인스턴스ID
// InstanceType : VM 생성시 사용할 인스턴스 타입 (ex) t2.micro
// KeyName : VM 생성시 사용할 키페어 이름 (ex) mcloud-barista-keypair
// MinCount :
// MaxCount :
// SubnetId : VM이 생성될 VPC의 SubnetId (ex) subnet-cf9ccf83
// SecurityGroupID : 생성할 VM에 적용할 보안그룹 ID (ex) sg-0df1c209ea1915e4b
type Config struct {
	Aws struct {
		AawsAccessKeyID    string `yaml:"aws_access_key_id"`
		AwsSecretAccessKey string `yaml:"aws_secret_access_key"`
		Region             string `yaml:"region"`

		ImageID string `yaml:"image_id"`

		VmID         string `yaml:"ec2_instance_id"`
		BaseName     string `yaml:"base_name"`
		InstanceType string `yaml:"instance_type"`
		KeyName      string `yaml:"key_name"`
		MinCount     int64  `yaml:"min_count"`
		MaxCount     int64  `yaml:"max_count"`

		SubnetID        string `yaml:"subnet_id"`
		SecurityGroupID string `yaml:"security_group_id"`
	} `yaml:"aws"`
}

//환경 설정 파일 읽기
//환경변수 CBSPIDER_PATH 설정 후 해당 폴더 하위에 /config/config.yaml 파일 생성해야 함.
func readConfigFile() Config {
	// Set Environment Value of Project Root Path
	rootPath := os.Getenv("CBSPIDER_PATH")
	//rootpath := "D:/Workspace/mcloud-barista-config"
	// /mnt/d/Workspace/mcloud-barista-config/config/config.yaml
	data, err := ioutil.ReadFile(rootPath + "/config/config.yaml")
	//data, err := ioutil.ReadFile("D:/Workspace/mcloud-bar-config/config/config.yaml")
	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	fmt.Println("Loaded ConfigFile...")
	spew.Dump(config)
	return config
}
