// Proof of Concepts for the Cloud-Barista Multi-Cloud Project.
//      * Cloud-Barista: https://github.com/cloud-barista
//
// EC2 Hander (AWS SDK GO Version 1.16.26, Thanks AWS.)
//
// by powerkim@powerkim.co.kr, 2019.03.
package resources

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"

	cblog "github.com/cloud-barista/cb-log"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
)

type AwsVMHandler struct {
	Region idrv.RegionInfo
	//Client *ec2drv.EC2
	Client *ec2.EC2
}

var cblogger *logrus.Logger

func init() {
	// cblog is a global variable.
	cblogger = cblog.GetLogger("AWS VMHandler")
}

func Connect(region string) *ec2.EC2 {
	// setup Region
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	if err != nil {
		fmt.Println("Could not create instance", err)
		return nil
	}

	// Create EC2 service client
	svc := ec2.New(sess)

	return svc
}

// 1개의 VM만 생성되도록 수정 (MinCount / MaxCount 이용 안 함)
// @Todo : SecurityGroupId 배열 처리 방안
func (vmHandler *AwsVMHandler) StartVM(vmReqInfo irs.VMReqInfo) (irs.VMInfo, error) {
	//fmt.Println("Start VMHandler()::StartVM()")
	cblogger.Info("Start VMHandler()::StartVM()")
	spew.Dump(vmReqInfo)

	imageID := vmReqInfo.ImageInfo.Id
	instanceType := vmReqInfo.SpecID // "t2.micro"
	minCount := aws.Int64(1)
	maxCount := aws.Int64(1)
	keyName := vmReqInfo.KeyPairInfo.Name
	securityGroupID := vmReqInfo.SecurityInfo.Id // "sg-0df1c209ea1915e4b"
	subnetID := vmReqInfo.VNetworkInfo.Id        // "subnet-cf9ccf83"
	baseName := vmReqInfo.Name                   //"mcloud-barista-VMHandlerTest"

	cblogger.Info("Create EC2 Instance")
	//fmt.Println("Create EC2 Instance")

	//키페어 이름(예:mcloud-barista)은 아래 URL에 나오는 목록 중 "키페어 이름"의 값을 적으면 됨.
	//https://ap-northeast-2.console.aws.amazon.com/ec2/v2/home?region=ap-northeast-2#KeyPairs:sort=keyName

	// Specify the details of the instance that you want to create.
	runResult, err := vmHandler.Client.RunInstances(&ec2.RunInstancesInput{
		// An Amazon Linux AMI ID for t2.micro instances in the us-west-2 region
		ImageId:      aws.String(imageID),
		InstanceType: aws.String(instanceType),
		MinCount:     minCount,
		MaxCount:     maxCount,
		KeyName:      aws.String(keyName), // set a keypair Name, ex) aws.powerkim.keypair
		SecurityGroupIds: []*string{
			aws.String(securityGroupID), // set a security group.
		},
		SubnetId: aws.String(subnetID), // set a subnet.
	})

	if err != nil {
		//fmt.Println("Could not create instance", err)
		cblogger.Errorf("Could not create instance", err)
		return irs.VMInfo{}, err
	}

	//fmt.Println("Created instance", *runResult.Instances[0].InstanceId)
	cblogger.Info("Created instance", *runResult.Instances[0].InstanceId)

	// Tag에 VM Name 설정
	_, errtag := vmHandler.Client.CreateTags(&ec2.CreateTagsInput{
		Resources: []*string{runResult.Instances[0].InstanceId},
		Tags: []*ec2.Tag{
			{
				Key:   aws.String("Name"),
				Value: aws.String(baseName),
			},
		},
	})

	if errtag != nil {
		log.Println("Could not create tags for instance", runResult.Instances[0].InstanceId, errtag)
		return irs.VMInfo{}, err
	}

	//빠른 생성을 위해 Running 상태를 대기하지 않고 최소한의 정보만 리턴 함.
	//Running 상태를 대기 후 Public Ip 등의 정보를 추출하려면 GetVM()을 호출해서 최신 정보를 다시 받아와야 함.
	//vmInfo :=GetVM(runResult.Instances[0].InstanceId)

	//cblogger.Info("EC2 Running 상태 대기")
	//WaitForRun(vmHandler.Client, *runResult.Instances[0].InstanceId)
	//cblogger.Info("EC2 Running 상태 완료 : ", runResult.Instances[0].State.Name)

	vmInfo := ExtractDescribeInstances(runResult)
	if vmInfo.Name == "" {
		vmInfo.Name = baseName
	}

	return vmInfo, nil
	//return irs.VMInfo{}, nil
}

func WaitForRun(svc *ec2.EC2, instanceID string) {
	cblogger.Infof("EC2 ID : [%s]", instanceID)

	input := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceID),
		},
	}
	err := svc.WaitUntilInstanceRunning(input)
	if err != nil {
		//fmt.Println("failed to wait until instances exist: %v", err)
		cblogger.Errorf("failed to wait until instances exist: %v", err)
	}
	cblogger.Info("=========WaitForRun() 종료")
}

func (vmHandler *AwsVMHandler) ResumeVM(vmID string) {
	cblogger.Infof("vmID : [%s]", vmID)
	input := &ec2.StartInstancesInput{
		InstanceIds: []*string{
			aws.String(vmID),
		},
		DryRun: aws.Bool(true),
	}
	result, err := vmHandler.Client.StartInstances(input)
	awsErr, ok := err.(awserr.Error)

	if ok && awsErr.Code() == "DryRunOperation" {
		// Let's now set dry run to be false. This will allow us to start the instances
		input.DryRun = aws.Bool(false)
		result, err = vmHandler.Client.StartInstances(input)
		if err != nil {
			//fmt.Println("Error", err)
			cblogger.Error(err)
		} else {
			//fmt.Println("Success", result.StartingInstances)
			cblogger.Info("Success", result.StartingInstances)
		}
	} else { // This could be due to a lack of permissions
		//fmt.Println("Error", err)
		cblogger.Error(err)
	}
}

func (vmHandler *AwsVMHandler) SuspendVM(vmID string) {
	cblogger.Infof("vmID : [%s]", vmID)
	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{
			aws.String(vmID),
		},
		DryRun: aws.Bool(true),
	}
	result, err := vmHandler.Client.StopInstances(input)
	awsErr, ok := err.(awserr.Error)
	if ok && awsErr.Code() == "DryRunOperation" {
		input.DryRun = aws.Bool(false)
		result, err = vmHandler.Client.StopInstances(input)
		if err != nil {
			//fmt.Println("Error", err)
			cblogger.Error(err)
		} else {
			//fmt.Println("Success", result.StoppingInstances)
			cblogger.Info("Success", result.StoppingInstances)
		}
	} else {
		//fmt.Println("Error", err)
		cblogger.Error(err)
	}
}

func (vmHandler *AwsVMHandler) RebootVM(vmID string) {
	cblogger.Infof("vmID : [%s]", vmID)
	input := &ec2.RebootInstancesInput{
		InstanceIds: []*string{
			aws.String(vmID),
		},
		DryRun: aws.Bool(true),
	}
	result, err := vmHandler.Client.RebootInstances(input)
	awsErr, ok := err.(awserr.Error)
	if ok && awsErr.Code() == "DryRunOperation" {
		input.DryRun = aws.Bool(false)
		result, err = vmHandler.Client.RebootInstances(input)
		if err != nil {
			//fmt.Println("Error", err)
			cblogger.Error(err)
		} else {
			//fmt.Println("Success", result)
			cblogger.Info("Success", result)
		}
	} else { // This could be due to a lack of permissions
		//fmt.Println("Error", err)
		cblogger.Error(err)
	}
	return
}

func (vmHandler *AwsVMHandler) TerminateVM(vmID string) {
	cblogger.Infof("vmID : [%s]", vmID)
	input := &ec2.TerminateInstancesInput{
		//InstanceIds: instanceIds,
		InstanceIds: []*string{
			aws.String(vmID),
		},
	}

	_, err := vmHandler.Client.TerminateInstances(input)
	if err != nil {
		//fmt.Println("Could not termiate instances", err)
		cblogger.Error("Could not termiate instances", err)
	} else {
		cblogger.Info("Success")
	}
	return
}

func (vmHandler *AwsVMHandler) GetVMStatus(vmID string) irs.VMStatus {
	return ""
}

func (vmHandler *AwsVMHandler) GetVM(vmID string) irs.VMInfo {
	cblogger.Infof("vmID : [%s]", vmID)

	input := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{
			aws.String(vmID),
		},
	}

	result, err := vmHandler.Client.DescribeInstances(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				//fmt.Println(aerr.Error())
				cblogger.Error(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and Message from an error.
			//fmt.Println(err.Error())
			cblogger.Error(err.Error())
		}
		return irs.VMInfo{}
	}

	cblogger.Info("Success", result)

	/*
		- 보안그룹의 경우 멀티개 설정이 가능한데 현재는 1개만 입력 받음
		- SecurityID에 보안그룹 Name을 할당하는게 맞는지 확인 필요
	*/
	vmInfo := irs.VMInfo{}
	for _, i := range result.Reservations {
		//vmInfo := ExtractDescribeInstances(result.Reservations[0])
		vmInfo = ExtractDescribeInstances(i)
	}

	/*
		vmInfo := irs.VMInfo{
			Name: *result.Reservations[0].Instances[0].Tags[0].Value,
			Id:   *result.Reservations[0].Instances[0].InstanceId,
			Region: irs.RegionInfo{
				Region: *result.Reservations[0].Instances[0].Placement.AvailabilityZone,
			},
			ImageID:      *result.Reservations[0].Instances[0].ImageId,
			SpecID:       *result.Reservations[0].Instances[0].InstanceType,
			VNetworkID:   *result.Reservations[0].Instances[0].NetworkInterfaces[0].VpcId,
			SubNetworkID: *result.Reservations[0].Instances[0].NetworkInterfaces[0].SubnetId,
			SecurityID:   *result.Reservations[0].Instances[0].NetworkInterfaces[0].Groups[0].GroupId,
			//SecurityName: *result.Reservations[0].Instances[0].NetworkInterfaces[0].Groups[0].GroupName,
			VNIC:           "eth0 - 값 위치 확인 필요",
			PublicIP:       *result.Reservations[0].Instances[0].NetworkInterfaces[0].Association.PublicIp,
			PublicDNS:      *result.Reservations[0].Instances[0].NetworkInterfaces[0].Association.PublicDnsName,
			PrivateIP:      *result.Reservations[0].Instances[0].NetworkInterfaces[0].PrivateIpAddress,
			PrivateDNS:     *result.Reservations[0].Instances[0].NetworkInterfaces[0].PrivateDnsName,
			KeyPairID:      *result.Reservations[0].Instances[0].KeyName,
			GuestUserID:    "",
			GuestBootDisk:  *result.Reservations[0].Instances[0].RootDeviceName,
			GuestBlockDisk: *result.Reservations[0].Instances[0].BlockDeviceMappings[0].DeviceName,
			AdditionalInfo: "",
		}
	*/

	cblogger.Info("vmInfo", vmInfo)

	return vmInfo
	//return irs.VMInfo{}
}

// DescribeInstances결과에서 EC2 세부 정보 추출
// VM 생성 시에는 Running 이전 상태의 정보가 넘어오기 때문에
// 최종 정보 기반으로 리턴 받고 싶으면 GetVM에 통합해야 할 듯.
func ExtractDescribeInstances(reservation *ec2.Reservation) irs.VMInfo {
	//cblogger.Info("ExtractDescribeInstances", reservation)
	cblogger.Debug("Instances[0]", reservation.Instances[0])
	//cblogger.Info("ImageId : [%s]", *reservation.Instances[0].ImageId)

	//len()

	//Running 상태에서만 체크 가능한 값
	//"stopped"
	//"terminated"
	//"running"
	var state string
	state = *reservation.Instances[0].State.Name
	cblogger.Info("EC2 상태 : [%s]", state)

	vmInfo := irs.VMInfo{
		//Name = *reservation.Instances[0].Tags[0].Value
		Id:             *reservation.Instances[0].InstanceId,
		ImageID:        *reservation.Instances[0].ImageId,
		SpecID:         *reservation.Instances[0].InstanceType,
		KeyPairID:      *reservation.Instances[0].KeyName,
		GuestUserID:    "",
		AdditionalInfo: "State:" + *reservation.Instances[0].State.Name,
	}

	if state == "running" {
		vmInfo.PublicIP = *reservation.Instances[0].NetworkInterfaces[0].Association.PublicIp
		vmInfo.PublicDNS = *reservation.Instances[0].NetworkInterfaces[0].Association.PublicDnsName

		vmInfo.GuestBlockDisk = *reservation.Instances[0].BlockDeviceMappings[0].DeviceName
	}

	if state != "terminated" {
		vmInfo.Region = irs.RegionInfo{
			Region: *reservation.Instances[0].Placement.AvailabilityZone,
		}
		vmInfo.VNetworkID = *reservation.Instances[0].NetworkInterfaces[0].VpcId
		vmInfo.SubNetworkID = *reservation.Instances[0].NetworkInterfaces[0].SubnetId
		vmInfo.SecurityID = *reservation.Instances[0].NetworkInterfaces[0].Groups[0].GroupId
		//SecurityName: *reservation.Instances[0].NetworkInterfaces[0].Groups[0].GroupName,
		vmInfo.VNIC = "eth0 - 값 위치 확인 필요"
		vmInfo.PrivateIP = *reservation.Instances[0].NetworkInterfaces[0].PrivateIpAddress
		vmInfo.PrivateDNS = *reservation.Instances[0].NetworkInterfaces[0].PrivateDnsName
		vmInfo.GuestBootDisk = *reservation.Instances[0].RootDeviceName
	}

	//EC2 Name 찾기
	cblogger.Debug("Name Tag 찾기")
	for _, t := range reservation.Instances[0].Tags {
		if *t.Key == "Name" {
			//nt = *t.Value
			vmInfo.Name = *t.Value
			cblogger.Debug("EC2 명칭 : ", vmInfo.Name)
			break
		}
	}
	//fmt.Println(nt, *i.InstanceID, *i.State.Name)

	/*
		for _, i := range reservation.Instances {
			//var nt string
			for _, t := range i.Tags {
				if *t.Key == "Name" {
					//nt = *t.Value
					vmInfo.Name = *t.Value
					break
				}
			}
			//fmt.Println(nt, *i.InstanceID, *i.State.Name)
			cblogger.Info("EC2 명칭 : ", vmInfo.Name, *i.InstanceId, *i.State.Name)
		}
	*/
	return vmInfo
}

func (vmHandler *AwsVMHandler) ListVM() []*irs.VMInfo {
	var vmList []*irs.VMInfo
	return vmList
}

func (vmHandler *AwsVMHandler) ListVMStatus() []*irs.VMStatusInfo {
	var vmStatus []*irs.VMStatusInfo
	return vmStatus
}
