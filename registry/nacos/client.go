package nacos

import (
	nacosKratos "github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/log"

	nacosClients "github.com/nacos-group/nacos-sdk-go/clients"
	nacosConstant "github.com/nacos-group/nacos-sdk-go/common/constant"
	nacosVo "github.com/nacos-group/nacos-sdk-go/vo"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
)

// NewRegistry 创建一个注册发现客户端 - Nacos
func NewRegistry(c *conf.Registry) *nacosKratos.Registry {
	if c == nil || c.Nacos == nil {
		return nil
	}

	srvConf := []nacosConstant.ServerConfig{
		*nacosConstant.NewServerConfig(c.Nacos.Address, c.Nacos.Port),
	}

	cliConf := nacosConstant.ClientConfig{
		NamespaceId:          c.Nacos.NamespaceId,
		TimeoutMs:            uint64(c.Nacos.Timeout.AsDuration().Milliseconds()), // http请求超时时间，单位毫秒
		BeatInterval:         c.Nacos.BeatInterval.AsDuration().Milliseconds(),    // 心跳间隔时间，单位毫秒
		UpdateThreadNum:      int(c.Nacos.UpdateThreadNum),                        // 更新服务的线程数
		LogLevel:             c.Nacos.LogLevel,
		CacheDir:             c.Nacos.CacheDir,             // 缓存目录
		LogDir:               c.Nacos.LogDir,               // 日志目录
		NotLoadCacheAtStart:  c.Nacos.NotLoadCacheAtStart,  // 在启动时不读取本地缓存数据，true--不读取，false--读取
		UpdateCacheWhenEmpty: c.Nacos.UpdateCacheWhenEmpty, // 当服务列表为空时是否更新本地缓存，true--更新,false--不更新
	}

	cli, err := nacosClients.NewNamingClient(
		nacosVo.NacosClientParam{
			ClientConfig:  &cliConf,
			ServerConfigs: srvConf,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	reg := nacosKratos.New(cli)

	return reg
}
