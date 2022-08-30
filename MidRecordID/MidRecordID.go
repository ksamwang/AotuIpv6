package MidRecordID

import (
	"Yu/rConfig"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

func SetAnalysis(Prname string,Ipv6Value string) {
	config := sdk.NewConfig()
	id , skey ,dname,_,_:= rConfig.Read()
	credential := credentials.NewAccessKeyCredential(id, skey)
	/* use STS Token
	credential := credentials.NewStsTokenCredential("<your-access-key-id>", "<your-access-key-secret>", "<your-sts-token>")
	*/
	client, err := alidns.NewClientWithOptions("cn-hangzhou", config, credential)
	if err != nil {
		panic(err)
	}

	request := alidns.CreateAddDomainRecordRequest()

	request.Scheme = "https"

	request.RR = Prname
	request.Type = "AAAA"
	request.Value = Ipv6Value
	request.DomainName = dname

	response, err := client.AddDomainRecord(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response.GetHttpContentString())
	//"{\"RequestId\":\"976F4B97-5BAF-56AA-87DC-D0B9CD5E84BC\",\"RecordId\":\"782295528690895872\"}"
}
