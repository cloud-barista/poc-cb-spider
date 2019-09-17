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

import (
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
)

type AwsSecurityHandler struct {
	Region idrv.RegionInfo
	Client *ec2.EC2
}

func (securityHandler *AwsSecurityHandler) CreateSecurity(securityReqInfo irs.SecurityReqInfo) (irs.SecurityInfo, error) {
	return irs.SecurityInfo{}, nil
}

func (securityHandler *AwsSecurityHandler) ListSecurity() ([]*irs.SecurityInfo, error) {
	return nil, nil
}

func (securityHandler *AwsSecurityHandler) GetSecurity(securityID string) (irs.SecurityInfo, error) {
	cblogger.Infof("securityID : [%s]", securityID)
	input := &ec2.DescribeSecurityGroupsInput{
		GroupIds: []*string{
			aws.String(securityID),
		},
	}

	result, err := securityHandler.Client.DescribeSecurityGroups(input)
	cblogger.Info("result : ", result)
	cblogger.Info("err : ", err)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				cblogger.Error(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			cblogger.Error(err.Error())
		}
		return irs.SecurityInfo{}, err
	}

	var ipPermissions []*irs.SecurityRuleInfo
	var ipPermissionsEgress []*irs.SecurityRuleInfo

	ipPermissions = ExtractIpPermissions(result.SecurityGroups[0].IpPermissions)
	cblogger.Info("InBouds : ", ipPermissions)
	ipPermissionsEgress = ExtractIpPermissions(result.SecurityGroups[0].IpPermissionsEgress)
	cblogger.Info("OutBounds : ", ipPermissionsEgress)
	//spew.Dump(ipPermissionsEgress)

	securityInfo := irs.SecurityInfo{
		GroupName: *result.SecurityGroups[0].GroupName,
		GroupID:   *result.SecurityGroups[0].GroupId,

		IPPermissions:       ipPermissions,       //AWS:InBounds
		IPPermissionsEgress: ipPermissionsEgress, //AWS:OutBounds

		Description: *result.SecurityGroups[0].Description,
		VpcID:       *result.SecurityGroups[0].VpcId,
		OwnerID:     *result.SecurityGroups[0].OwnerId,
	}

	//Name은 Tag의 "Name" 속성에만 저장됨
	cblogger.Debug("Name Tag 찾기")
	for _, t := range result.SecurityGroups[0].Tags {
		if *t.Key == "Name" {
			securityInfo.Name = *t.Value
			cblogger.Debug("Name : ", securityInfo.Name)
			break
		}
	}

	return securityInfo, nil
}

func ExtractIpPermissions(ipPermissions []*ec2.IpPermission) []*irs.SecurityRuleInfo {

	var results []*irs.SecurityRuleInfo

	for _, ip := range ipPermissions {
		cblogger.Info("Inbound 정보 조회 : ", *ip.IpProtocol)
		securityRuleInfo := new(irs.SecurityRuleInfo)

		if !reflect.ValueOf(ip.FromPort).IsNil() {
			securityRuleInfo.FromPort = *ip.FromPort
		}

		if !reflect.ValueOf(ip.ToPort).IsNil() {
			securityRuleInfo.ToPort = *ip.FromPort
		}

		securityRuleInfo.IPProtocol = *ip.IpProtocol
		securityRuleInfo.Cidr = *ip.IpRanges[0].CidrIp

		results = append(results, securityRuleInfo)
	}

	return results
}

func (securityHandler *AwsSecurityHandler) DeleteSecurity(securityID string) (bool, error) {
	cblogger.Infof("securityID : [%s]", securityID)

	// Delete the security group.
	_, err := securityHandler.Client.DeleteSecurityGroup(&ec2.DeleteSecurityGroupInput{
		GroupId: aws.String(securityID),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case "InvalidGroupId.Malformed":
				fallthrough
			case "InvalidGroup.NotFound":
				cblogger.Errorf("%s.", aerr.Message())
				return false, err
			}
		}
		cblogger.Errorf("Unable to get descriptions for security groups, %v.", err)
		return false, err
	}

	cblogger.Infof("Successfully delete security group %q.", securityID)

	return true, nil
}
