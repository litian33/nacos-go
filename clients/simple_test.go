package clients

import (
	"errors"
	"fmt"
	"github.com/litian33/nacos-go/clients/service_client"
	"github.com/litian33/nacos-go/common/constant"
	"github.com/litian33/nacos-go/vo"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestHeartBeat(t *testing.T) {
	client, err := nacosClient("172.24.28.3:8848")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	//param := vo.RegisterServiceInstanceParam{Ip: "172.22.1.102", Port: 33333, ServiceName: "testxxx", Enable: true, Healthy: true, Ephemeral: true}
	//res, err := client.RegisterServiceInstance(param)
	//assert.Nil(t, err)
	//assert.NotNil(t, res)

	//param1 := vo.KeepAliveParam{Ip: "172.22.1.102", Port: 8848, ServiceName: "testxxx", Healthy:true}
	xxx := make(map[string]string)
	xxx["node_type"]="xxxxx"
	param1 := vo.BeatTaskParam{Ip: "172.22.1.102", Port: 3335, ServiceName: "xxxxxxxxx", Metadata:xxx, Ephemeral:true}

	client.StartBeatTask(param1)
	//client.KeepAlive(param1)

	time.Sleep(10 * 60 * time.Second)
}

func nacosClient(servers string) (iClient service_client.IServiceClient,
	err error) {
	urls := strings.Split(servers, ",")
	var serverConfigs []constant.ServerConfig
	for _, server := range urls {
		ip, port, err := parseAddr(server)
		if err != nil || len(ip) < 4 {
			return nil, fmt.Errorf("invalid server url: %s", server)
		}
		serverConfigs = append(serverConfigs, constant.ServerConfig{IpAddr: ip, Port: uint64(port), ContextPath: "/nacos"})
	}

	if len(serverConfigs) == 0 {
		return nil, errors.New("no valid nacos server exists")
	}

	// 如果参数设置不合法，将抛出error
	client, err := CreateServiceClient(map[string]interface{}{
		constant.KEY_SERVER_CONFIGS: serverConfigs,
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}

func parseAddr(addr string) (ip string, port int, err error) {
	parts := strings.Split(addr, ":")
	if parts == nil || len(parts) != 2 {
		return "", 0, fmt.Errorf("invalid cluster address %s", addr)
	}
	ip = parts[0]
	port, err = strconv.Atoi(parts[1])
	return
}
