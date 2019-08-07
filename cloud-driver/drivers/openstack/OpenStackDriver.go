// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is a Cloud Driver Example for PoC Test.
//
// by hyokyung.kim@innogrid.co.kr, 2019.07.

package openstack

import (
	"fmt"
	oscon "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/openstack/connect"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	icon "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/connect"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type OpenStackDriver struct{}

func (OpenStackDriver) GetDriverVersion() string {
	return "OPENSTACK DRIVER Version 1.0"
}

func (OpenStackDriver) GetDriverCapability() idrv.DriverCapabilityInfo {
	var drvCapabilityInfo idrv.DriverCapabilityInfo

	drvCapabilityInfo.ImageHandler = false
	drvCapabilityInfo.VNetworkHandler = false
	drvCapabilityInfo.SecurityHandler = false
	drvCapabilityInfo.KeyPairHandler = false
	drvCapabilityInfo.VNicHandler = false
	drvCapabilityInfo.PublicIPHandler = false
	drvCapabilityInfo.VMHandler = true

	return drvCapabilityInfo
}

/* org
func (OpenStackDriver) ConnectCloud(connectionInfo idrv.ConnectionInfo) (icon.CloudConnection, error) {
	// 1. get info of credential and region for Test A Cloud from connectionInfo.
	// 2. create a client object(or service  object) of Test A Cloud with credential info.
	// 3. create CloudConnection Instance of "connect/TDA_CloudConnection".
	// 4. return CloudConnection Interface of TDA_CloudConnection.

	// sample code, do not user like this^^
	var iConn icon.CloudConnection
	iConn = oscon.OpenStackCloudConnection{}

	return iConn, nil // return type: (icon.CloudConnection, error)
}
*/

// modifiled by powerkim, 2019.07.29.
func (driver *OpenStackDriver) ConnectCloud(connectionInfo idrv.ConnectionInfo) (icon.CloudConnection, error) {
	// 1. get info of credential and region for Test A Cloud from connectionInfo.
	// 2. create a client object(or service  object) of Test A Cloud with credential info.
	// 3. create CloudConnection Instance of "connect/TDA_CloudConnection".
	// 4. return CloudConnection Interface of TDA_CloudConnection.

	// sample code, do not user like this^^

	Client, err := getServiceClient()
	if err != nil {
		panic(err)
	}

	//var iConn icon.CloudConnection
	iConn := oscon.OpenStackCloudConnection{Client, nil}

	return &iConn, nil // return type: (icon.CloudConnection, error)
}

func (driver *OpenStackDriver) ConnectNetworkCloud(connectionInfo idrv.ConnectionInfo) (icon.CloudConnection, error) {

	NetworkClient, err := getNetworkClient()
	if err != nil {
		panic(err)
	}

	//var iConn icon.CloudConnection
	iConn := oscon.OpenStackCloudConnection{nil, NetworkClient}

	return &iConn, nil // return type: (icon.CloudConnection, error)
}

//--------- temporary  by powerkim, 2019.07.29.
type Config struct {
	Openstack struct {
		DomainName       string `yaml:"domain_name"`
		IdentityEndpoint string `yaml:"identity_endpoint"`
		Password         string `yaml:"password"`
		ProjectID        string `yaml:"project_id"`
		Username         string `yaml:"username"`
		Region           string `yaml:"region"`
		VMName           string `yaml:"vm_name"`
		ImageId          string `yaml:"image_id"`
		FlavorId         string `yaml:"flavor_id"`
		NetworkId        string `yaml:"network_id"`
		SecurityGroups   string `yaml:"security_groups"`
		KeypairName      string `yaml:"keypair_name"`

		ServerId string `yaml:"server_id"`
	} `yaml:"openstack"`
}

func readConfigFile() Config {
	// Set Environment Value of Project Root Path
	rootPath := os.Getenv("CBSPIDER_PATH")
	data, err := ioutil.ReadFile(rootPath + "/config/config.yaml")
	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	return config
}

//--------- temporary

// moved by powerkim, 2019.07.29.
func getServiceClient() (*gophercloud.ServiceClient, error) {

	// read configuration YAML file
	config := readConfigFile()

	fmt.Println(config)

	opts := gophercloud.AuthOptions{
		IdentityEndpoint: config.Openstack.IdentityEndpoint,
		Username:         config.Openstack.Username,
		Password:         config.Openstack.Password,
		DomainName:       config.Openstack.DomainName,
		TenantID:         config.Openstack.ProjectID,
		/*Scope: &gophercloud.AuthScope{
			ProjectID: config.Openstack.ProjectID,
		},*/
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return nil, err
	}

	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: config.Openstack.Region,
	})
	if err != nil {
		return nil, err
	}

	return client, err
}

var TestDriver OpenStackDriver

func getNetworkClient() (*gophercloud.ServiceClient, error) {

	config := readConfigFile()

	authOpts := gophercloud.AuthOptions{
		IdentityEndpoint: config.Openstack.IdentityEndpoint,
		Username:         config.Openstack.Username,
		Password:         config.Openstack.Password,
		DomainName:       config.Openstack.DomainName,
		TenantID:         config.Openstack.ProjectID,
	}

	provider, err := openstack.AuthenticatedClient(authOpts)
	if err != nil {
		return nil, err
	}

	client, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{
		Name:   "neutron",
		Region: "RegionOne",
	})
	if err != nil {
		return nil, err
	}

	return client, err
}
