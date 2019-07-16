package resources

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-06-01/compute"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/config"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
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

	future, err := vmClient.CreateOrUpdate(ctx, config.Azure.GroupName, config.Azure.Os.ComputeName, vmOpts)
	if err != nil {
		panic(err)
		return  irs.VMInfo{}, err
	}

	err = future.WaitForCompletionRef(ctx, vmClient.Client)
	if err != nil {
		panic(err)
		return  irs.VMInfo{}, err
	}

	return irs.VMInfo{}, nil
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
	return nil
}

func (AzureVMHandler) GetVMStatus(vmID string) irs.VMStatus {
	return irs.VMStatus("")
}

func (AzureVMHandler) ListVM() []*irs.VMInfo {
	config := config.ReadConfigFile()

	ctx, cancle, vmClient := getVMClient()
	defer cancle()
	//defer resources.Cleanup(ctx)

	serverList, err := vmClient.List(ctx, config.Azure.GroupName)
	fmt.Println(serverList.Values())
	if err != nil {
		panic(err)
	}

	var vmList []*irs.VMInfo
	for _, server := range serverList.Values() {
		vmInfo := MappingServerInfo(server)
		fmt.Println(vmInfo)
		vmList = append(vmList, &vmInfo)
	}

	return vmList
}

func (AzureVMHandler) GetVM(vmID string) irs.VMInfo {
	config := config.ReadConfigFile()

	ctx, cancle, vmClient := getVMClient()
	defer cancle()
	//defer resources.Cleanup(ctx)

	vm, err := vmClient.Get(ctx, config.Azure.GroupName, config.Azure.VMName, compute.InstanceView)
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
		//SpecID:
	}

	// Get Spec Info
	//specInfo := server.VirtualMachineProperties.HardwareProfile.VMSize
	//fmt.Println(specInfo)
	//imageInfo := server.VirtualMachineProperties.StorageProfile.ImageReference
	//fmt.Println(imageInfo)
	//vmInfo.SpecID = imageInfo.Publisher + imageInfo.Offer + imageInfo.Sku + imageInfo.Version
	//vmInfo.ImageID = *imageInfo.ID

	return vmInfo
}

func getVMClient() (context.Context, context.CancelFunc, compute.VirtualMachinesClient) {
	config := config.ReadConfigFile()

	// Need to Initialize environment variable (AZURE_AUTH_LOCATION)
	//azureAuthPath := os.Getenv("AZURE_AUTH_LOCATION")
	//fmt.Println("AZURE_AUTH_LOCATION=", azureAuthPath)

	authorizer, err := auth.NewAuthorizerFromFile(azure.PublicCloud.ResourceManagerEndpoint)
	if err != nil {
		panic(err)
	}

	vmClient := compute.NewVirtualMachinesClient(config.Azure.SubscriptionID)
	vmClient.Authorizer = authorizer

	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Second)

	return ctx, cancel, vmClient
}
