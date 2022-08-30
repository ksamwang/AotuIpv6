package GetDescribeDomainRecords
//GetDescribeDomainRecords


import (
	"Yu/UpData"
	"Yu/rConfig"
	"encoding/json"
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	alidns "github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"

)

type Records struct {
	RR string `json:"RR"`
	Line string `json:"Line"`
	Status string `json:"Status"`
	Locked bool `json:"Locked"`
	Types string `json:"Types"`
	DomainName string `json:"DomainName"`
	Value string `json:"Value"`
	RecordId string `json:"RecordId"`
	TTL int `json:"TTL"`
	Weight int `json:"Weight"`
}

type DomainRecords struct {
	Record []interface{} `json:"record"`
}
type R struct {
	TotalCount int `json:"TotalCount"`
	RequestId string `json:"RequestId"`
	PageSize int `json:"PageSize"`
	DomainRecords  `json:"DomainRecords"`
	PageNumber int `json:"PageNumber"`
}
func GetAna(PrName string,IPC6Value string) (int,string,error){
	id , skey ,dname,_,_:= rConfig.Read()
	config := sdk.NewConfig()
	credential := credentials.NewAccessKeyCredential(id,skey)
	/* use STS Token
	credential := credentials.NewStsTokenCredential("<your-access-key-id>", "<your-access-key-secret>", "<your-sts-token>")
	*/
	client, err := alidns.NewClientWithOptions("cn-hangzhou", config, credential)
	if err != nil {
		fmt.Println(err)
	}

	request := alidns.CreateDescribeDomainRecordsRequest()

	request.Scheme = "https"

	request.DomainName = dname

	response, err := client.DescribeDomainRecords(request)
	if err != nil {
		fmt.Print(err.Error())
	}

	var r R
	json.Unmarshal([]byte(response.GetHttpContentString()),&r)
	i:=0
	for ;i<len(r.DomainRecords.Record); {
		var Info Records
		bstr,_ :=json.Marshal(r.DomainRecords.Record[i])
		json.Unmarshal([]byte(bstr),&Info)
		//fmt.Println(Info.RecordId)
		if Info.RR==PrName {
			if Info.Value==IPC6Value {
				return 0,Info.RecordId,nil
			}else{
				fmt.Println("没有找到主机记录为：",PrName,"的记录值，正在添加记录值:",IPC6Value)
				UpData.UpdateDomainRecord(PrName,IPC6Value,Info.RecordId)
				return 0,Info.RecordId,nil
			}
		}
		i++
	}
	return -1,"",nil
	
	
	
}
