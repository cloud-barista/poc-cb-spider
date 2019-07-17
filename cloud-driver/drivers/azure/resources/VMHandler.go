package resources

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-06-01/compute"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/config"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"strings"
	"time"
)

type AzureVMHandler struct{}

func (AzureVMHandler) StartVM(vmReqInfo irs.VMReqInfo) (irs.VMInfo, error) {
	config := config.ReadConfigFile()

	ctx, cancle, vmClient := getVMClient()
	defer cancle()

	vmOpts := compute.VirtualMachine{
		Location: &config.Azure.Location,
		VirtualMachineProperties: &compute.VirtualMachineProperties{
			HardwareProfile: &compute.HardwareProfile{
				VMSize: config.Azure.VMSize,
			},
			StorageProfile: &compute.StorageProfile{
				ImageReference: &compute.ImageReference{
					Publisher: &config.Azure.Image.Publisher,
					Offer: &config.Azure.Image.Offer,
					Sku: &config.Azure.Image.Sku,
					Version: &config.Azure.Image.Version,
				},
			},
			OsProfile: &compute.OSProfile{
				ComputerName: &config.Azure.Os.ComputeName,
				AdminUsername: &config.Azure.Os.AdminUsername,
				AdminPassword: &config.Azure.Os.AdminPassword,
			},
			NetworkProfile: &compute.NetworkProfile{
				NetworkInterfaces: &[]compute.NetworkInterfaceReference{
					{
						ID: &config.Azure.Network.ID,
						NetworkInterfaceReferenceProperties: &compute.NetworkInterfaceReferenceProperties{
							Primary: &config.Azure.Network.Primary,
						},
					},
				},
			},
		},
	}

	groupName := config.Azure.GroupName
	vmName := config.Azure.VMName

	future, err := vmClient.CreateOrUpdate(ctx, groupName, vmName, vmOpts)
	if err != nil {
		panic(err)
		return  irs.VMInfo{}, err
	}

	err = future.WaitForCompletionRef(ctx, vmClient.Client)
	if err != nil {
		panic(err)
		return  irs.VMInfo{}, err
	}

	vm, err := vmClient.Get(ctx, groupName, vmName, compute.InstanceView)
	if err != nil {
		panic(err)
	}
	vmInfo := MappingServerInfo(vm)

	return vmInfo, nil
}

func (AzureVMHandler) SuspendVM(vmID string) {
	config := config.ReadConfigFile()

	ctx, cancle, vmClient := getVMClient()
	defer cancle()
	//defer resources.Cleanup(ctx)

	future, err := vmClient.PowerOff(ctx, config.Azure.GroupName, vmID)
	if err != nil {
		panic(err)
	}
	err = future.WaitForCompletionRef(ctx, vmClient.Client)
	if err != nil {
		panic(err)
	}
}

func (AzureVMHandler) ResumeVM(vmID string) {
	config := config.ReadConfigFile()

	ctx, cancle, vmClient := getVMClient()
	defer cancle()
	//defer resources.Cleanup(ctx)

	future, err := vmClient.Start(ctx, config.Azure.GroupName, vmID)
	if err != nil {
		panic(err)
	}
	err = future.WaitForCompletionRef(ctx, vmClient.Client)
	if err != nil {
		panic(err)
	}
}

func (AzureVMHandler) RebootVM(vmID string) {
	config := config.ReadConfigFile()

	ctx, cancle, vmClient := getVMClient()
	defer cancle()
	//defer resources.Cleanup(ctx)

	future, err := vmClient.Restart(ctx, config.Azure.GroupName, vmID)
	if err != nil {
		panic(err)
	}
	err = future.WaitForCompletionRef(ctx, vmClient.Client)
	if err != nil {
		panic(err)
	}
}

func (AzureVMHandler) TerminateVM(vmID string) {
	config := config.ReadConfigFile()

	ctx, cancle, vmClient := getVMClient()
	defer cancle()
	//defer resources.Cleanup(ctx)

	//future, err := vmClient.Delete(ctx, config.Azure.GroupName, vmID)
	future, err := vmClient.Deallocate(ctx, config.Azure.GroupName, vmID)

	if err != nil {
		panic(err)
	}
	err = future.WaitForCompletionRef(ctx, vmClient.Client)
	if err != nil {
		panic(err)
	}
}

func (AzureVMHandler) ListVMStatus() []*irs.VMStatus {
	config := config.ReadConfigFile()

	ctx, cancle, vmClient := getVMClient()
	defer cancle()
	//defer resources.Cleanup(ctx)

	serverList, err := vmClient.List(ctx, config.Azure.GroupName)
	if err != nil {
		panic(err)
	}

	var vmStatusList []*irs.VMStatus
	for _, s := range serverList.Values() {
		//vmStatus := irs.VMStatus(*s.ProvisioningState)
		vmStatus := AzureVMHandler{}.GetVMStatus(*s.Name)
		vmStatusList = append(vmStatusList, &vmStatus)
	}

	return vmStatusList
}

func (AzureVMHandler) GetVMStatus(vmID string) irs.VMStatus {
	config := config.ReadConfigFile()

	ctx, cancle, vmClient := getVMClient()
	defer cancle()
	//defer resources.Cleanup(ctx)

	instanceView, err := vmClient.InstanceView(ctx, config.Azure.GroupName, vmID)
	if err != nil {
		panic(err)
	}

	// Get powerState, provisioningState
	var powerState, provisioningState string

	for _, stat := range *instanceView.Statuses {
		statArr := strings.Split(*stat.Code, "/")

		if statArr[0] == "PowerState" {
			powerState = statArr[1]
		} else if statArr[0] == "ProvisioningState" {
			provisioningState = statArr[1]
		}
	}

	// Set VM Status Info
	var vmState string
	if powerState != "" && provisioningState != "" {
		vmState = powerState + "(" + provisioningState + ")"
	} else if powerState != "" && provisioningState == "" {
		vmState = powerState
	} else if powerState == "" && provisioningState != "" {
		vmState = provisioningState
	} else {
		vmState = "-"
	}

	return irs.VMStatus(vmState)
}

func (AzureVMHandler) ListVM() []*irs.VMInfo {
	config := config.ReadConfigFile()

	ctx, cancle, vmClient := getVMClient()
	defer cancle()
	//defer resources.Cleanup(ctx)

	serverList, err := vmClient.List(ctx, config.Azure.GroupName)
	if err != nil {
		panic(err)
	}

	var vmList []*irs.VMInfo
	for _, server := range serverList.Values() {
		vmInfo := MappingServerInfo(server)
		vmList = append(vmList, &vmInfo)
	}

	return vmList
}

func (AzureVMHandler) GetVM(vmID string) irs.VMInfo {
	config := config.ReadConfigFile()

	ctx, cancle, vmClient := getVMClient()
	defer cancle()
	//defer resources.Cleanup(ctx)

	vm, err := vmClient.Get(ctx, config.Azure.GroupName, vmID, compute.InstanceView)
	if err != nil {
		panic(err)
	}

	vmInfo := MappingServerInfo(vm)
	return vmInfo
}

func MappingServerInfo(server compute.VirtualMachine) irs.VMInfo {

	// Get Default VM Info
	vmInfo := irs.VMInfo{
		Name: *server.Name,
		Id: *server.ID,
		Region: irs.RegionInfo{
			Region: *server.Location,
		},
		SpecID: string(server.VirtualMachineProperties.HardwareProfile.VMSize),
	}

	// Set VNetwork Info
	niList := *server.NetworkProfile.NetworkInterfaces
	for _, ni := range niList {
		if *ni.Primary {
			vmInfo.VNetworkID = *ni.ID
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
		vmInfo.GuestBootDisk  = *server.VirtualMachineProperties.StorageProfile.OsDisk.Name
	}

	return vmInfo
}

func getVMClient() (context.Context, context.CancelFunc, compute.VirtualMachinesClient) {
	config := config.ReadConfigFile()

	authorizer, err := auth.NewAuthorizerFromFile(azure.PublicCloud.ResourceManagerEndpoint)
	if err != nil {
		panic(err)
	}

	vmClient := compute.NewVirtualMachinesClient(config.Azure.SubscriptionID)
	vmClient.Authorizer = authorizer

	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Second)

	return ctx, cancel, vmClient
}
