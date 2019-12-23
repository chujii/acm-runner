package handler

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
)

type Nocos struct {
	Client config_client.IConfigClient
}

func NewNocos(nsMap map[string]interface{}) (*Nocos, error){
	clientConfig  := constant.ClientConfig{
		Endpoint:       nsMap["end_point"].(string),
		NamespaceId:    nsMap["id"].(string),
		AccessKey:      nsMap["access_key"].(string),
		SecretKey:      nsMap["secret_key"].(string),
		TimeoutMs:      5 * 1000,
		ListenInterval: 30 * 1000,
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"clientConfig": clientConfig,
	})

	if err != nil {
		return nil, err
	}

	return &Nocos{
		Client: configClient,
	}, nil
}
func (nc *Nocos) ListenConfig(listMap []interface{},fun func(data string, filename string))  {
	for _, item := range listMap {
		itemMap := item.(map[interface{}]interface{})
		// 监听配置
		err := nc.Client.ListenConfig(vo.ConfigParam{
			DataId: itemMap["data_id"].(string),
			Group:  itemMap["group"].(string),
			OnChange: func(namespace, group, dataId, data string) {
				fun(data, itemMap["filename"].(string))
			},
		})

		if err != nil {
			log.Fatalln(err)
		}
	}
}