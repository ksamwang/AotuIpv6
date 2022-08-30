package rConfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	AccessKeyID string `json:"accesskeyid"`
	AccessKeySecret string `json:"accesskeysecret"`
	DomainName string `json:"domainname"`
	PrName string `json:"prname"`
}
func checkError(err error){
	if err != nil{
		fmt.Println(err)
	}
}
func Read()(string,string,string,string,error){
	str, _ := os.Getwd()
	str = str+"\\config.json"

	// 打开json文件
	jsonFile, err := os.Open(str)

	// 最好要处理以下错误
	if err != nil {
		fmt.Println(err)
	}

	// 要记得关闭
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	checkError(err)
	config := &Config{}

	err = json.Unmarshal(byteValue, &config)
	checkError(err)

	return config.AccessKeyID ,config.AccessKeySecret,config.DomainName,config.PrName,nil
}
