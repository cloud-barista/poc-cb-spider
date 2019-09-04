package resources

import (
	"fmt"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client/ace/nic"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	irs "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/resources"
	"github.com/davecgh/go-spew/spew"
	"strconv"
)

type ClouditNicHandler struct {
	CredentialInfo idrv.CredentialInfo
	Client         *client.RestClient
}

func (nicHandler *ClouditNicHandler) CreateVNic(vNicReqInfo irs.VNicReqInfo) (irs.VNicInfo, error) {
	return irs.VNicInfo{}, nil
}

//Todo : 401에러 수정 중
func (nicHandler *ClouditNicHandler) ListVNic() ([]*irs.VNicInfo, error) {
	nicHandler.Client.TokenID = nicHandler.CredentialInfo.AuthToken
	authHeader := nicHandler.Client.AuthenticatedHeaders()

	requestOpts := client.RequestOpts{
		//JSONBody:     nil,
		//RawBody:      nil,
		//JSONResponse: nil,
		//OkCodes:      nil,
		MoreHeaders: authHeader,
	}
	serverId := "7b35148e-b033-4749-a29e-12beddaebbfa"
	vNicList, err := nic.List(nicHandler.Client, serverId, &requestOpts)
	if err != nil {
		panic(err)
	}

	for i, nic := range *vNicList {
		fmt.Println("[" + strconv.Itoa(i) + "]")
		spew.Dump(nic)
	}
	return nil, nil
}

func (nicHandler *ClouditNicHandler) GetVNic(vNicID string) (irs.VNicInfo, error) {

	return irs.VNicInfo{}, nil
}
func (nicHandler *ClouditNicHandler) DeleteVNic(vNicID string) (bool, error) {
	return false, nil
}
