package UpData

import (
	"Yu/rConfig"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

func UpdateDomainRecord(pr string,ipv6 string,RocId string)  {
	id , skey,_,_,_ := rConfig.Read()
	config := sdk.NewConfig()

	credential := credentials.NewAccessKeyCredential(id, skey)
	/* use STS Token
	credential := credentials.NewStsTokenCredential("<your-access-key-id>", "<your-access-key-secret>", "<your-sts-token>")
	*/
	client, err := alidns.NewClientWithOptions("cn-hangzhou", config, credential)
	if err != nil {
		fmt.Println(err)
	}

	request := alidns.CreateUpdateDomainRecordRequest()

	request.Scheme = "https"

	request.RecordId = RocId
	request.RR = pr
	//"@"
	request.Type = "AAAA"
	request.Value = ipv6

	response, err := client.UpdateDomainRecord(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("UpData ,Status: %#v\n", response)
}
