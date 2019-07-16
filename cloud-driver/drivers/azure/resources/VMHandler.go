 resources

import (
	"context"
	"fmt"
	_ "github.com/Azure-Samples/azure-sdk-for-go-samples/resources"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-06-01/compute"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/config"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"time"
)

type AzureVMHandler struct{}

//authInfo, err := readJSON(os.Getenv("AZURE_AUTH_LOCATION"))
// authPath := os.Getenv("AZURE_AUTH_LOCATION")
// fmt.Println(authPath)

func (AzureVMHandler) StartVM(vmReqInfo irs.VMReqInfo) (irs.VMInfo, error) {
	fmt.Println("StartVM")
	config := config.ReadConfigFile()

	ctx, cancle, vmClient := getVMClient()
	defer cancle()



	return irs.VMInfo{}, nil
}

func (AzureVMHandler) SuspendVM(vmID string) {
	fmt.Println("SuspendVM")
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
	fmt.Println("SuspendVM Finished")
}

func (AzureVMHandler) ResumeVM(vmID string) {
	fmt.Println("ResumeVM")
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
	fmt.Println("ResumeVM Finished")
}

func (AzureVMHandler) RebootVM(vmID string) {
	fmt.Println("RebootVM")
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
	fmt.Println("RebootVM Finished")
}

func (AzureVMHandler) TerminateVM(vmID string) {
	fmt.Println("TerminateVM")
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
	fmt.Println("TerminateVM Finished")
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

	/*vm, err := vmClient.Get(ctx, config.Azure.GroupName, config.Azure.VMName, compute.InstanceView)
	if err != nil {
		panic(err)
	}*/
	vmClient.G

	return nil
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
	specInfo := server.VirtualMachineProperties.HardwareProfile.VMSize
	fmt.Println(specInfo)

	imageInfo := server.VirtualMachineProperties.StorageProfile.ImageReference
	fmt.Println(imageInfo)
	//vmInfo.SpecID = imageInfo.Publisher + imageInfo.Offer + imageInfo.Sku + imageInfo.Version
	//vmInfo.ImageID = *imageInfo.ID

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
