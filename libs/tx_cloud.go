package libs

import (
	"sync"

	"github.com/spf13/viper"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

func init() {
	viper.MustBindEnv("TX_SECRET_ID")
	viper.MustBindEnv("TX_SECRET_KEY")
}

var txClient *dnspod.Client
var txClientOnce sync.Once

func GetTxClient() *dnspod.Client {
	txClientOnce.Do(func() {
		credential := common.NewCredential(
			viper.GetString("TC_SECRET_ID"),
			viper.GetString("TC_SECRET_KEY"),
		)
		cpf := profile.NewClientProfile()
		cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
		txClient, _ = dnspod.NewClient(credential, "", cpf)
	})
	return txClient
}
