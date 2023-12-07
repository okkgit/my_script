package scripts

import (
	"fmt"
	"go_script/libs"

	"github.com/imroc/req/v3"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

const Domain = "okkgit.top"

func updateIPToDNS(recordId uint64, subDomain string) {
	resp := req.MustGet("http://metadata.tencentyun.com/meta-data/public-ipv4")
	ip := string(resp.Bytes())

	request := dnspod.NewModifyRecordRequest()
	request.Domain = common.StringPtr(Domain)
	request.SubDomain = common.StringPtr(subDomain)
	request.RecordType = common.StringPtr("A")
	request.RecordLine = common.StringPtr("默认")
	request.Value = common.StringPtr(ip)
	request.TTL = common.Uint64Ptr(600)
	request.RecordId = common.Uint64Ptr(recordId)
	client := libs.GetTxClient()

	response, err := client.ModifyRecord(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

func UpdateIPToDNS() {
	updateIPToDNS(1676141943, "www")
	updateIPToDNS(1677124006, "@")
}
