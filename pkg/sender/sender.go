package sender

import (
	"log"
	"mail-server-test/pkg/config"
	"mail-server-test/pkg/responser"
	"net/http"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

type AliyunConfig struct {
	AccessKeyId      string
	AccessKeySecret  string
	RegionId         string
	EmailAccountName string
}

var SdkClient *sdk.Client
var aliyunConfig *AliyunConfig

func init() {
	aliyunConfig = new(AliyunConfig)
	aliyunConfig.AccessKeyId = config.GetString("ALIYUN_ACCESS_KEY_ID")
	aliyunConfig.AccessKeySecret = config.GetString("ALIYUN_ACCESS_KEY_SECRET")
	aliyunConfig.RegionId = config.GetString("ALIYUN_REGION_ID")
	aliyunConfig.EmailAccountName = config.GetString("ALIYUN_EMAIL_ACCOUNT_NAME")

	c, err := sdk.NewClientWithAccessKey(aliyunConfig.RegionId, aliyunConfig.AccessKeyId, aliyunConfig.AccessKeySecret)
	if err != nil {
		// Handle exceptions
		panic(err)
	}
	SetSdkClient(c)
}

func SetSdkClient(client *sdk.Client) {
	SdkClient = client
}

func SendSingleMail(email string, title string, body string) string {

	r := new(responser.Body)

	if len(email) <= 0 {
		r.Code = 100
		r.Message = "email错误"
		return responser.GenerateJSON(r)
	}
	if len(title) <= 0 {
		r.Code = 100
		r.Message = "title错误"
		return responser.GenerateJSON(r)
	}
	if len(body) <= 0 {
		r.Code = 100
		r.Message = "body错误"
		return responser.GenerateJSON(r)
	}

	request := requests.NewCommonRequest()
	request.Domain = "dm.aliyuncs.com"
	request.Version = "2015-11-23"
	request.Product = "DM"
	request.Method = "POST"
	request.ApiName = "SingleSendMail"
	request.QueryParams["AccountName"] = aliyunConfig.EmailAccountName
	request.QueryParams["AddressType"] = "1"
	request.QueryParams["ReplyToAddress"] = "false"
	request.QueryParams["Subject"] = title
	request.QueryParams["ToAddress"] = email
	request.QueryParams["Action"] = "SingleSendMail"
	request.QueryParams["TextBody"] = body
	request.TransToAcsRequest()

	respon, err := SdkClient.ProcessCommonRequest(request)
	if err != nil {
		log.Println(err)
		r.Code = -1
		r.Message = "消息发送失败，请查看系统日志"
		return responser.GenerateJSON(r)
	}
	if respon.GetHttpStatus() != http.StatusOK {
		log.Println(respon.String())
		r.Code = -1
		r.Message = "消息发送失败，请查看系统日志"
		return responser.GenerateJSON(r)
	}

	r.Code = 0
	r.Message = respon.GetHttpContentString()
	return responser.GenerateJSON(r)
}
