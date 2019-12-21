package handler

import (
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
)

type NocosConf struct {
	Endpoint string
	NamespaceId string
	AccessKey string
	SecretKey string
}

type Nocos struct {
	Client config_client.ConfigClient
}

func NewNocos(conf NocosConf) (*Nocos){
}