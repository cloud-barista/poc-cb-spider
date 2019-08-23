package resources

import (
	"fmt"
	"strconv"

	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client/ace/nic"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/davecgh/go-spew/spew"

)

type ClouditNicHandler struct {
	Client *client.RestClient
}

func (nicHandler *ClouditNicHandler) CreateVNic(vNicReqInfo irs.VNicReqInfo) (irs.VNicInfo, error) {
	return  irs.VNicInfo{}, nil
}

func (nicHandler *ClouditNicHandler) ListVNic() ([]*irs.VNicInfo, error) {
	vNicList, _ := nic.List(nicHandler.Client, "1")
	for i, nic := range *vNicList {
		fmt.Println("["+ strconv.Itoa(i) +"]")
		spew.Dump(nic)
	}
	return nil, nil
}

func (nicHandler *ClouditNicHandler) GetVNic(vNicID string) (irs.VNicInfo, error) {

	return  irs.VNicInfo{}, nil
}
func (nicHandler *ClouditNicHandler) DeleteVNic(vNicID string) (bool, error) {
	return false, nil
}
