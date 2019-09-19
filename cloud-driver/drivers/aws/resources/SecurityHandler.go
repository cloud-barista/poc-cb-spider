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
	input := &ec2.DescribeSecurityGroupsInput{
		GroupIds: []*string{
			nil,
		},
	}

	result, err := securityHandler.Client.DescribeSecurityGroups(input)
	//cblogger.Info("result : ", result)
	if err != nil {
		cblogger.Info("err : ", err)
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
		return nil, err
	}

	var results []*irs.SecurityInfo
	for _, securityGroup := range result.SecurityGroups {
		securityInfo := ExtractSecurityInfo(securityGroup)
		results = append(results, &securityInfo)
	}

	return results, nil
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

	securityInfo := ExtractSecurityInfo(result.SecurityGroups[0])
	/*
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
	*/
	return securityInfo, nil
}

func ExtractSecurityInfo(securityGroupResult *ec2.SecurityGroup) irs.SecurityInfo {
	var ipPermissions []*irs.SecurityRuleInfo
	var ipPermissionsEgress []*irs.SecurityRuleInfo

	cblogger.Info("===[그룹아이디:%s]===", *securityGroupResult.GroupId)
	ipPermissions = ExtractIpPermissions(securityGroupResult.IpPermissions)
	cblogger.Info("InBouds : ", ipPermissions)
	ipPermissionsEgress = ExtractIpPermissions(securityGroupResult.IpPermissionsEgress)
	cblogger.Info("OutBounds : ", ipPermissionsEgress)
	//spew.Dump(ipPermissionsEgress)

	securityInfo := irs.SecurityInfo{
		GroupName: *securityGroupResult.GroupName,
		GroupID:   *securityGroupResult.GroupId,

		IPPermissions:       ipPermissions,       //AWS:InBounds
		IPPermissionsEgress: ipPermissionsEgress, //AWS:OutBounds

		Description: *securityGroupResult.Description,
		VpcID:       *securityGroupResult.VpcId,
		OwnerID:     *securityGroupResult.OwnerId,
	}

	//Name은 Tag의 "Name" 속성에만 저장됨
	cblogger.Debug("Name Tag 찾기")
	for _, t := range securityGroupResult.Tags {
		if *t.Key == "Name" {
			securityInfo.Name = *t.Value
			cblogger.Debug("Name : ", securityInfo.Name)
			break
		}
	}

	return securityInfo
}

// IpPermission에서 공통정보 추출
func ExtractIpPermissionCommon(ip *ec2.IpPermission, securityRuleInfo *irs.SecurityRuleInfo) {
	//공통 정보
	if !reflect.ValueOf(ip.FromPort).IsNil() {
		securityRuleInfo.FromPort = *ip.FromPort
	}

	if !reflect.ValueOf(ip.ToPort).IsNil() {
		securityRuleInfo.ToPort = *ip.ToPort
	}

	securityRuleInfo.IPProtocol = *ip.IpProtocol
}

//@TODO : CIDR이 없는 경우 구조처 처리해야 함.(예: 타겟이 ELB거나 다른 보안 그룹일 경우))
//@TODO : InBound / OutBound의 배열 처리및 테스트해야 함.
func ExtractIpPermissions(ipPermissions []*ec2.IpPermission) []*irs.SecurityRuleInfo {

	var results []*irs.SecurityRuleInfo

	for _, ip := range ipPermissions {
		//securityRuleInfo := new(irs.SecurityRuleInfo)

		//ipv4 처리
		for _, ipv4 := range ip.IpRanges {
			cblogger.Info("Inbound/Outbound 정보 조회 : ", *ip.IpProtocol)
			securityRuleInfo := new(irs.SecurityRuleInfo)
			securityRuleInfo.Cidr = *ipv4.CidrIp

			/*
				//공통 정보
				if !reflect.ValueOf(ip.FromPort).IsNil() {
					securityRuleInfo.FromPort = *ip.FromPort
				}

				if !reflect.ValueOf(ip.ToPort).IsNil() {
					securityRuleInfo.ToPort = *ip.ToPort
				}

				securityRuleInfo.IPProtocol = *ip.IpProtocol
			*/

			ExtractIpPermissionCommon(ip, securityRuleInfo)
			results = append(results, securityRuleInfo)
		}

		//ipv6 처리
		for _, ipv6 := range ip.Ipv6Ranges {
			securityRuleInfo := new(irs.SecurityRuleInfo)
			securityRuleInfo.Cidr = *ipv6.CidrIpv6

			ExtractIpPermissionCommon(ip, securityRuleInfo)
			results = append(results, securityRuleInfo)
		}

		//ELB나 보안그룹 참조 방식 처리
		for _, userIdGroup := range ip.UserIdGroupPairs {
			securityRuleInfo := new(irs.SecurityRuleInfo)
			securityRuleInfo.Cidr = *userIdGroup.GroupId
			// *userIdGroup.GroupName / *userIdGroup.UserId

			ExtractIpPermissionCommon(ip, securityRuleInfo)
			results = append(results, securityRuleInfo)
		}

		/*
			if !reflect.ValueOf(ip.IpRanges).IsNil() {
				securityRuleInfo.Cidr = *ip.IpRanges[0].CidrIp
			} else {
				//ELB나 다른 보안그룹 참조처럼 IpRanges가 없고 UserIdGroupPairs가 있는 경우 처리
				//https://docs.aws.amazon.com/ko_kr/elasticloadbalancing/latest/classic/elb-security-groups.html
				if !reflect.ValueOf(ip.UserIdGroupPairs).IsNil() {
					securityRuleInfo.Cidr = *ip.UserIdGroupPairs[0].GroupId
				} else {
					cblogger.Error("미지원 보안 그룹 형태 발견 - 구조 파악 필요 ", ip)
				}
			}
		*/
	}

	return results
}

//@TODO : CIDR이 없는 경우 구조처 처리해야 함.(예: 타겟이 ELB거나 다른 보안 그룹일 경우))
//@TODO : InBound / OutBound의 배열 처리및 테스트해야 함.
func _ExtractIpPermissions(ipPermissions []*ec2.IpPermission) []*irs.SecurityRuleInfo {

	var results []*irs.SecurityRuleInfo

	for _, ip := range ipPermissions {
		cblogger.Info("Inbound/Outbound 정보 조회 : ", *ip.IpProtocol)
		securityRuleInfo := new(irs.SecurityRuleInfo)

		if !reflect.ValueOf(ip.FromPort).IsNil() {
			securityRuleInfo.FromPort = *ip.FromPort
		}

		if !reflect.ValueOf(ip.ToPort).IsNil() {
			securityRuleInfo.ToPort = *ip.ToPort
		}

		//IpRanges가 없고 UserIdGroupPairs가 있는 경우가 있음(ELB / 보안 그룹 참조 등)
		securityRuleInfo.IPProtocol = *ip.IpProtocol

		if !reflect.ValueOf(ip.IpRanges).IsNil() {
			securityRuleInfo.Cidr = *ip.IpRanges[0].CidrIp
		} else {
			//ELB나 다른 보안그룹 참조처럼 IpRanges가 없고 UserIdGroupPairs가 있는 경우 처리
			//https://docs.aws.amazon.com/ko_kr/elasticloadbalancing/latest/classic/elb-security-groups.html
			if !reflect.ValueOf(ip.UserIdGroupPairs).IsNil() {
				securityRuleInfo.Cidr = *ip.UserIdGroupPairs[0].GroupId
			} else {
				cblogger.Error("미지원 보안 그룹 형태 발견 - 구조 파악 필요 ", ip)
			}
		}

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
