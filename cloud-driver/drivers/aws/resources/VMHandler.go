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
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"

	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"

	cblog "github.com/cloud-barista/cb-log"
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

func (vmHandler *AwsVMHandler) StartVM(vmReqInfo irs.VMReqInfo) (irs.VMInfo, error) {
	// Set VM Create Information
	//imageID := vmReqInfo.ImageInfo.Id
	/*
		imageID := "ami-047f7b46bd6dd5d84"
		instanceType := "t2.micro"
		minCount := 1
		maxCount := 1
		keyName := "mcloud-barista"
		securityGroupID := "sg-0df1c209ea1915e4b"
		subnetID := "subnet-cf9ccf83"
		baseName := "mcloud-barista-VMHandlerTest"
	*/

	//fmt.Println("Start VMHandler()::StartVM()")
	cblogger.Info("Start VMHandler()::StartVM()")
	spew.Dump(vmReqInfo)

	imageID := vmReqInfo.ImageInfo.Id
	instanceType := vmReqInfo.SpecID // "t2.micro"
	minCount := 1
	maxCount := 1
	keyName := vmReqInfo.KeyPairInfo.Name
	securityGroupID := vmReqInfo.SecurityInfo.Id // "sg-0df1c209ea1915e4b"
	subnetID := vmReqInfo.VNetworkInfo.Id        // "subnet-cf9ccf83"
	baseName := vmReqInfo.Name                   //"mcloud-barista-VMHandlerTest"

	cblogger.Info("Create EC2 Instance")
	//fmt.Println("Create EC2 Instance")

	//키페어 이름(예:mcloud-barista)은 아래 URL에 나오는 목록 중 "키페어 이름"의 값을 적으면 됨.
	//https://ap-northeast-2.console.aws.amazon.com/ec2/v2/home?region=ap-northeast-2#KeyPairs:sort=keyName
	//instanceIds := ec2handler.CreateInstances(vmHandler.Client, imageID, instanceType, minCount, maxCount,
	instanceIds := CreateInstances(vmHandler.Client, imageID, instanceType, minCount, maxCount,
		keyName, securityGroupID, subnetID, baseName)

	// waiting for completion of new instance running.
	for _, v := range instanceIds {
		// wait until running status
		//ec2handler.WaitForRun(vmHandler.Client, *v)
		WaitForRun(vmHandler.Client, *v)
		// get public IP
		//publicIP, err := ec2handler.GetPublicIP(vmHandler.Client, *v)
		publicIP, err := GetPublicIP(vmHandler.Client, *v)
		if err != nil {
			//fmt.Println("Error", err)
			cblogger.Error(err)
			return irs.VMInfo{}, err
		}
		//fmt.Println(publicIP)
		cblogger.Info(publicIP)
	}

	return irs.VMInfo{}, nil
}

/*
func mappingServerInfo(server compute.VirtualMachine) irs.VMInfo {

	// Get Default VM Info
	vmInfo := irs.VMInfo{
		Name: *server.Name,
		Id:   *server.ID,
		Region: irs.RegionInfo{
			Region: *server.Location,
		},
		SpecID: string(server.VirtualMachineProperties.HardwareProfile.VMSize),
	}

	// Set VM Zone
	if server.Zones != nil {
		vmInfo.Region.Zone = (*server.Zones)[0]
	}

	// Set VM Image Info
	imageRef := server.VirtualMachineProperties.StorageProfile.ImageReference
	imageId := *imageRef.Publisher + ":" + *imageRef.Offer + ":" + *imageRef.Sku + ":" + *imageRef.Version
	vmInfo.ImageID = imageId

	// Set VNic Info
	niList := *server.NetworkProfile.NetworkInterfaces
	for _, ni := range niList {
		if ni.NetworkInterfaceReferenceProperties != nil {
			vmInfo.VNIC = *ni.ID
		}
	}

	// Set GuestUser Id/Pwd
	if server.VirtualMachineProperties.OsProfile.AdminUsername != nil {
		vmInfo.GuestUserID = *server.VirtualMachineProperties.OsProfile.AdminUsername
	}
	if server.VirtualMachineProperties.OsProfile.AdminPassword != nil {
		vmInfo.GuestUserID = *server.VirtualMachineProperties.OsProfile.AdminPassword
	}

	// Set BootDisk
	if server.VirtualMachineProperties.StorageProfile.OsDisk.Name != nil {
		vmInfo.GuestBootDisk = *server.VirtualMachineProperties.StorageProfile.OsDisk.Name
	}

	return vmInfo
}
*/

func (vmHandler *AwsVMHandler) StartVM2(vmReqInfo irs.VMReqInfo) (irs.VMInfo, error) {
	// Set VM Create Information
	//imageID := vmReqInfo.ImageInfo.Id
	imageID := "ami-047f7b46bd6dd5d84"
	instanceType := "t2.micro"
	minCount := 1
	maxCount := 1
	keyName := "mcloud-barista"
	securityGroupID := "sg-0df1c209ea1915e4b"
	subnetID := "subnet-cf9ccf83"
	baseName := "mcloud-barista-VMHandlerTest"

	runResult, err := vmHandler.Client.RunInstances(&ec2.RunInstancesInput{
		ImageId:      aws.String(imageID),        // set imageID ex) ami-047f7b46bd6dd5d84
		InstanceType: aws.String(instanceType),   // instance Type, ex) t2.micro
		MinCount:     aws.Int64(int64(minCount)), //
		MaxCount:     aws.Int64(int64(maxCount)),
		KeyName:      aws.String(keyName), // set a keypair Name, ex) aws.powerkim.keypair
		SecurityGroupIds: []*string{
			aws.String(securityGroupID), // set a security group.
		},
		//SubnetId: aws.String("subnet-8c4a53e4"),     // set a subnet.
		SubnetId: aws.String(subnetID), // set a subnet.
	})

	if err != nil {
		fmt.Println("Could not create instance", err)
		return irs.VMInfo{}, err
	}

	// copy Instances's ID
	instanceIds := make([]*string, len(runResult.Instances))
	for k, v := range runResult.Instances {
		instanceIds[k] = v.InstanceId
	}
	/*
	       for i:=0; i<maxCount; i++ {
	   	    fmt.Println("Created instance", *runResult.Instances[i].InstanceId)
	   	    instanceID = *runResult.Instances[i].InstanceId
	       }
	*/

	for i := 0; i < maxCount; i++ {
		// Add tags to the created instance
		_, errtag := vmHandler.Client.CreateTags(&ec2.CreateTagsInput{
			Resources: []*string{runResult.Instances[i].InstanceId},
			Tags: []*ec2.Tag{
				{
					Key:   aws.String("Name"),
					Value: aws.String(baseName + strconv.Itoa(i)),
				},
			},
		})
		if errtag != nil {
			log.Println("Could not create tags for instance", runResult.Instances[i].InstanceId, errtag)
			return irs.VMInfo{}, errtag
		}
		fmt.Println("Successfully tagged instance:" + baseName + strconv.Itoa(i))
	} // end of for

	return irs.VMInfo{}, nil
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

/*
func Close() {
}
*/

func CreateInstances(svc *ec2.EC2, imageID string, instanceType string,
	minCount int, maxCount int, keyName string, securityGroupID string,
	subnetID string, baseName string) []*string {

	runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
		ImageId:      aws.String(imageID),        // set imageID ex) ami-047f7b46bd6dd5d84
		InstanceType: aws.String(instanceType),   // instance Type, ex) t2.micro
		MinCount:     aws.Int64(int64(minCount)), //
		MaxCount:     aws.Int64(int64(maxCount)),
		KeyName:      aws.String(keyName), // set a keypair Name, ex) aws.powerkim.keypair
		SecurityGroupIds: []*string{
			aws.String(securityGroupID), // set a security group.
		},
		//SubnetId: aws.String("subnet-8c4a53e4"),     // set a subnet.
		SubnetId: aws.String(subnetID), // set a subnet.
	})

	if err != nil {
		fmt.Println("Could not create instance", err)
		return nil
	}

	// copy Instances's ID
	instanceIds := make([]*string, len(runResult.Instances))
	for k, v := range runResult.Instances {
		instanceIds[k] = v.InstanceId
	}
	/*
	       for i:=0; i<maxCount; i++ {
	   	    fmt.Println("Created instance", *runResult.Instances[i].InstanceId)
	   	    instanceID = *runResult.Instances[i].InstanceId
	       }
	*/

	for i := 0; i < maxCount; i++ {
		// Add tags to the created instance
		_, errtag := svc.CreateTags(&ec2.CreateTagsInput{
			Resources: []*string{runResult.Instances[i].InstanceId},
			Tags: []*ec2.Tag{
				{
					Key:   aws.String("Name"),
					Value: aws.String(baseName + strconv.Itoa(i)),
				},
			},
		})
		if errtag != nil {
			log.Println("Could not create tags for instance", runResult.Instances[i].InstanceId, errtag)
			return nil
		}
		fmt.Println("Successfully tagged instance:" + baseName + strconv.Itoa(i))
	} // end of for

	return instanceIds
}

func GetPublicIP(svc *ec2.EC2, instanceID string) (string, error) {
	var publicIP string

	input := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceID),
		},
	}
	// Call to get detailed information on each instance
	result, err := svc.DescribeInstances(input)
	if err != nil {
		fmt.Println("Error", err)
		return publicIP, err
	}

	//    fmt.Println(result)

	for i, _ := range result.Reservations {
		for _, inst := range result.Reservations[i].Instances {
			publicIP = *inst.PublicIpAddress
		}
	}
	return publicIP, err
}

func WaitForRun(svc *ec2.EC2, instanceID string) {
	input := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceID),
		},
	}
	err := svc.WaitUntilInstanceRunning(input)
	if err != nil {
		fmt.Println("failed to wait until instances exist: %v", err)
	}
}

func DestroyInstances(svc *ec2.EC2, instanceIds []*string) error {

	//input := &ec2.TerminateInstancesInput(instanceIds)
	input := &ec2.TerminateInstancesInput{
		InstanceIds: instanceIds,
	}

	_, err := svc.TerminateInstances(input)

	if err != nil {
		fmt.Println("Could not termiate instances", err)
	}

	return err
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

	instanceIds := []*string{&vmID}
	//spew.Dump(instanceIds)

	input := &ec2.TerminateInstancesInput{
		InstanceIds: instanceIds,
	}

	/*
		input := &ec2.DescribeInstancesInput{
			InstanceIds: []*string{
				aws.String(vmID),
			},
		}
	*/

	_, err := vmHandler.Client.TerminateInstances(input)
	if err != nil {
		//fmt.Println("Could not termiate instances", err)
		cblogger.Error(err)
	} else {
		//fmt.Println("Success", result)
		cblogger.Info("Success")
	}

	//fmt.Println("End VMHandler()::TerminateVM()")

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
		Name      string
		Id        string
		StartTime time.Time // Timezone: based on cloud-barista server location.

		Region       RegionInfo // ex) {us-east1, us-east1-c} or {ap-northeast-2}
		ImageID      string     // ex) ami-047f7b46bd6dd5d84 or projects/gce-uefi-images/global/images/centos-7-v20190326
		SpecID       string     // instance type or flavour, etc... ex) t2.micro or f1-micro
		VNetworkID   string     // ex) vpc-23ed0a4b
		SubNetworkID string     // ex) subnet-8c4a53e4
		SecurityID   string     // ex) sg-0b7452563e1121bb6

		VNIC       string // ex) eth0
		PublicIP   string // ex) 13.125.43.21
		PublicDNS  string // ex) ec2-13-125-43-0.ap-northeast-2.compute.amazonaws.com
		PrivateIP  string // ex) ip-172-31-4-60.ap-northeast-2.compute.internal
		PrivateDNS string // ex) 172.31.4.60

		KeyPairID    string // ex) powerkimKeyPair
		GuestUserID  string // ex) user1
		GuestUserPwd string

		GuestBootDisk  string // ex) /dev/sda1
		GuestBlockDisk string // ex)

		AdditionalInfo string // Any information to be good for users and developers.
	*/

	/*
		- 보안그룹의 경우 멀티개 설정이 가능한데 현재는 1개만 입력 받음
		- SecurityID에 보안그룹 Name을 할당하는게 맞는지 확인 필요
	*/

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

	return vmInfo
	//return irs.VMInfo{}
}

func (vmHandler *AwsVMHandler) ListVM() []*irs.VMInfo {
	var vmList []*irs.VMInfo
	return vmList
}

func (vmHandler *AwsVMHandler) ListVMStatus() []*irs.VMStatusInfo {
	var vmStatus []*irs.VMStatusInfo
	return vmStatus
}
